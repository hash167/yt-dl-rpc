package main

import (
	"flag"
	"io"
	"log/slog"
	"net"
	"os"
	"runtime"
	"strconv"

	rpc "github.com/hash167/yt-dl-rpc/api/rpc"

	"github.com/hash167/yt-dl-rpc/internal"
	server "github.com/hash167/yt-dl-rpc/server"
	"github.com/hash167/yt-dl-rpc/server/config"
	"google.golang.org/grpc"
)

var (
	host            string
	port            int
	queueSize       int
	configFile      string
	downloadPath    string
	downloaderPath  string
	sessionFilePath string
)

func init() {

	flag.StringVar(&host, "host", "0.0.0.0", "Host where server will listen at")
	flag.IntVar(&port, "port", 3033, "Port where server will listen at")
	flag.IntVar(&queueSize, "qs", runtime.NumCPU(), "Download queue size")

	flag.StringVar(&configFile, "conf", "./config.yml", "Config file path")
	flag.StringVar(&downloadPath, "out", ".", "Where files will be saved")
	flag.StringVar(&downloaderPath, "driver", "yt-dlp", "yt-dlp executable path")
	flag.StringVar(&sessionFilePath, "session", ".", "session file path")

	flag.Parse()
}

func main() {
	logger := slog.New(
		slog.NewTextHandler(
			io.MultiWriter(os.Stdout),
			nil,
		),
	)
	c := config.Instance()

	c.Host = host
	c.Port = port
	c.QueueSize = queueSize
	c.DownloadPath = downloadPath
	c.DownloaderPath = downloaderPath
	c.SessionFilePath = sessionFilePath
	logger.Info("config", slog.String("config", c.String()))

	// if config file is found it will be merged with the current config struct
	if err := c.LoadFile(configFile); err != nil {
		logger.Warn("failed to load config file", slog.String("err", err.Error()))
	}

	// Listen on a port
	lis, err := net.Listen("tcp", host+":"+strconv.Itoa(port))
	if err != nil {
		logger.Warn("failed to listen", slog.String("err", err.Error()))
	}

	// Create a gRPC server object
	s := grpc.NewServer()

	// initialize the memory db and message queue
	var mdb internal.MemoryDB
	mq := internal.NewMessageQueue()
	go mq.Subscriber()

	rpcServer := server.RPCServiceServerImpl{
		Logger: logger,
		Mdb:    &mdb,
		Mq:     mq,
	}

	// Attach the RPC service to the server
	rpc.RegisterRPCServiceServer(s, &rpcServer)

	// Add rpc handlers

	// Serve the gRPC server
	logger.Info("yt rpc service started", slog.Int("port", port))

	if err := s.Serve(lis); err != nil {
		logger.Warn("http server stopped", slog.String("err", err.Error()))
	}

}
