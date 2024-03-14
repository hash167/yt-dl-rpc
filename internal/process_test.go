package internal

import (
	"os/exec"
	"time"

	rpc "github.com/hash167/yt-dl-rpc/api/rpc"
)

type MockCommandExecutor struct{}

func (m MockCommandExecutor) Command(name string, arg ...string) *exec.Cmd {
	// return a command that does something harmless
	return exec.Command("echo", "hello")
}

type mockProcess struct {
	Url      string
	Progress *rpc.DownloadProgress
	Executor CommandExecutor
}

func (mp *mockProcess) Start() {
	mp.Progress = &rpc.DownloadProgress{
		Status:     StatusDownloading,
		Percentage: "0",
		Speed:      0,
		ETA:        0,
	}
	time.Sleep(time.Duration(3) * time.Second)
	mp.Complete()
}

func (mp *mockProcess) Complete() {
	mp.Progress = &rpc.DownloadProgress{
		Status:     StatusCompleted,
		Percentage: "100",
		Speed:      0,
		ETA:        0,
	}
}
