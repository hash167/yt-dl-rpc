package internal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log/slog"
	"regexp"
	"sync"
	"syscall"

	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	rpc "github.com/hash167/yt-dl-rpc/api/rpc"
	"github.com/hash167/yt-dl-rpc/internal/rx"
	"github.com/hash167/yt-dl-rpc/internal/slices"
	"github.com/hash167/yt-dl-rpc/server/cli"
	"github.com/hash167/yt-dl-rpc/server/config"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const template = `download:
{
	"eta":%(progress.eta)s, 
	"percentage":"%(progress._percent_str)s",
	"speed":%(progress.speed)s
}`

const (
	StatusPending = iota
	StatusDownloading
	StatusCompleted
	StatusErrored
)

type ProgressTemplate struct {
	Percentage string  `json:"percentage"`
	Speed      float32 `json:"speed"`
	Size       string  `json:"size"`
	Eta        float32 `json:"eta"`
}

type ProcessInterface interface {
	Start()
}

type CommandExecutor interface {
	Command(name string, arg ...string) *exec.Cmd
}

type RealCommandExecutor struct{}

func (r RealCommandExecutor) Command(name string, arg ...string) *exec.Cmd {
	return exec.Command(name, arg...)
}

// Process descriptor
type Process struct {
	Id       string
	Url      string
	Params   []string
	Info     *rpc.DownloadInfo
	Progress *rpc.DownloadProgress
	Output   *rpc.Output
	proc     *os.Process
	Logger   *slog.Logger
	Executor CommandExecutor
}

// Starts spawns/forks a new yt-dlp process and parse its stdout.
// The process is spawned to outputting a custom progress text that
// Resembles a JSON Object in order to Unmarshal it later.
// This approach is anyhow not perfect: quotes are not escaped properly.
// Each process is not identified by its PID but by a UUIDv4
func (p *Process) Start() {
	// escape bash variable escaping and command piping, you'll never know
	// what they might come with...
	p.Params = slices.Filter(p.Params, func(e string) bool {
		match, _ := regexp.MatchString(`(\$\{)|(\&\&)`, e)
		return !match
	})

	p.Params = slices.Filter(p.Params, func(e string) bool {
		return e != ""
	})

	out := rpc.Output{
		Path:     config.Instance().DownloadPath,
		Filename: "%(title)s.%(ext)s",
	}

	if p.Output.Path != "" {
		out.Path = p.Output.Path
	}

	if p.Output.Filename != "" {
		out.Filename = p.Output.Filename
	}

	buildFilename(p.Output, p.Logger)

	params := []string{
		strings.Split(p.Url, "?list")[0], //no playlist
		"--newline",
		"--no-colors",
		"--no-playlist",
		"--progress-template",
		strings.NewReplacer("\n", "", "\t", "", " ", "").Replace(template),
	}

	// if user asked to manually override the output path...
	if !(slices.Includes(params, "-P") || slices.Includes(params, "--paths")) {
		params = append(params, "-o")
		params = append(params, fmt.Sprintf("%s/%s", out.Path, out.Filename))
	}

	params = append(params, p.Params...)

	// ----------------- main block ----------------- //
	cmd := p.Executor.Command(config.Instance().DownloaderPath, params...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	r, err := cmd.StdoutPipe()
	if err != nil {
		p.Logger.Error(
			"failed to connect to stdout",
			slog.String("err", err.Error()),
		)
		panic(err)
	}
	scan := bufio.NewScanner(r)

	err = cmd.Start()
	if err != nil {
		p.Logger.Error(
			"failed to start yt-dlp process",
			slog.String("err", err.Error()),
		)
		panic(err)
	}

	p.proc = cmd.Process

	// --------------- progress block --------------- //
	var (
		sourceChan = make(chan []byte)
		doneChan   = make(chan struct{})
	)

	// spawn a goroutine that does the dirty job of parsing the stdout
	// filling the channel with as many stdout line as yt-dlp produces (producer)
	go func() {
		defer func() {
			r.Close()
			p.Complete()
			doneChan <- struct{}{}
			close(sourceChan)
			close(doneChan)
		}()

		for scan.Scan() {
			sourceChan <- scan.Bytes()
		}
	}()

	// Slows down the unmarshal operation to every 500ms
	go func() {
		rx.Sample(time.Millisecond*500, sourceChan, doneChan, func(event []byte) {
			stdout := ProgressTemplate{}
			err := json.Unmarshal(event, &stdout)
			if err == nil {
				p.Progress = &rpc.DownloadProgress{
					Status:     StatusDownloading,
					Percentage: stdout.Percentage,
					Speed:      stdout.Speed,
					ETA:        stdout.Eta,
				}
				p.Logger.Info("progress",
					slog.String("id", p.getShortId()),
					slog.String("url", p.Url),
					slog.String("percentege", stdout.Percentage),
				)
			}
		})
	}()

	// ------------- end progress block ------------- //
	cmd.Wait()
}

// Keep process in the memoryDB but marks it as complete
// Convention: All completed processes has progress -1
// and speed 0 bps.
func (p *Process) Complete() {
	p.Progress = &rpc.DownloadProgress{
		Status:     StatusCompleted,
		Percentage: "-1",
		Speed:      0,
		ETA:        0,
	}

	p.Logger.Info("finished",
		slog.String("id", p.getShortId()),
		slog.String("url", p.Url),
	)
}

// Kill a process and remove it from the memory
func (p *Process) Kill() error {
	// yt-dlp uses multiple child process the parent process
	// has been spawned with setPgid = true. To properly kill
	// all subprocesses a SIGTERM need to be sent to the correct
	// process group
	if p.proc != nil {
		pgid, err := syscall.Getpgid(p.proc.Pid)
		if err != nil {
			return err
		}
		err = syscall.Kill(-pgid, syscall.SIGTERM)

		p.Logger.Info("killed process", slog.String("id", p.Id))
		return err
	}

	return nil
}

// Returns the available format for this URL
func (p *Process) GetFormatsSync() (*rpc.DownloadFormats, error) {
	cmd := exec.Command(config.Instance().DownloaderPath, p.Url, "-J")
	stdout, err := cmd.Output()

	if err != nil {
		return &rpc.DownloadFormats{}, err
	}

	info := rpc.DownloadFormats{URL: p.Url}
	best := rpc.Format{}

	var (
		wg            sync.WaitGroup
		decodingError error
	)

	wg.Add(2)

	if err != nil {
		return &rpc.DownloadFormats{}, err
	}

	log.Println(
		cli.BgRed, "Metadata", cli.Reset,
		cli.BgBlue, "Formats", cli.Reset,
		p.Url,
	)

	p.Logger.Info(
		"retrieving metadata",
		slog.String("caller", "getFormats"),
		slog.String("url", p.Url),
	)

	go func() {
		decodingError = json.Unmarshal(stdout, &info)
		wg.Done()
	}()

	go func() {
		decodingError = json.Unmarshal(stdout, &best)
		wg.Done()
	}()

	wg.Wait()

	if decodingError != nil {
		return &rpc.DownloadFormats{}, err
	}

	info.Best = &best

	return &info, nil
}

func (p *Process) SetPending() {
	p.Progress.Status = StatusPending
}

func (p *Process) SetMetadata() error {
	cmd := exec.Command(config.Instance().DownloaderPath, p.Url, "-J")
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		p.Logger.Error("failed retrieving info",
			slog.String("id", p.getShortId()),
			slog.String("url", p.Url),
			slog.String("err", err.Error()),
		)
		return err
	}

	info := rpc.DownloadInfo{
		URL:       p.Url,
		CreatedAt: timestamppb.New(time.Now()),
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	p.Logger.Info("retrieving metadata",
		slog.String("id", p.getShortId()),
		slog.String("url", p.Url),
	)

	err = json.NewDecoder(stdout).Decode(&info)
	if err != nil {
		return err
	}

	p.Info = &info
	p.Progress.Status = StatusPending

	err = cmd.Wait()

	return err
}

func (p *Process) getShortId() string {
	return strings.Split(p.Id, "-")[0]
}

func buildFilename(o *rpc.Output, p *slog.Logger) {
	p.Info("building filename", slog.String("filename", o.Filename))
	if o.Filename != "" && strings.Contains(o.Filename, ".%(ext)s") {
		o.Filename += ".%(ext)s"
	}

	o.Filename = strings.Replace(
		o.Filename,
		".%(ext)s.%(ext)s",
		".%(ext)s",
		1,
	)
}
