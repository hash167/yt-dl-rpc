package server

import (
	"context"
	"log/slog"

	rpc "github.com/hash167/yt-dl-rpc/api/rpc"
	"github.com/hash167/yt-dl-rpc/internal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RPCServiceServerImpl struct {
	Logger                            *slog.Logger
	Mdb                               *internal.MemoryDB
	Mq                                *internal.MessageQueue
	rpc.UnimplementedRPCServiceServer // Embed the unimplemented server
}

func (s *RPCServiceServerImpl) Exec(ctx context.Context, req *rpc.DownloadRequest) (*rpc.ProcessResponse, error) {
	p := &internal.Process{
		Url:    req.URL,
		Params: req.Params,
		Logger: s.Logger,
		Output: &rpc.Output{
			Path:     req.Path,
			Filename: req.Rename,
		},
		Executor: internal.RealCommandExecutor{},
	}
	id := s.Mdb.Set(p)
	s.Mq.AddProcess(p)
	pr := &rpc.ProcessResponse{Id: id}
	return pr, nil
}

func (s *RPCServiceServerImpl) ExecPlaylist(ctx context.Context, req *rpc.DownloadRequest) (*rpc.ProcessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecPlaylist not implemented")
}

func (s *RPCServiceServerImpl) Progress(ctx context.Context, req *rpc.Args) (*rpc.DownloadProgress, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Progress not implemented")
}

func (s *RPCServiceServerImpl) Formats(ctx context.Context, req *rpc.Args) (*rpc.DownloadFormats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Formats not implemented")
}

func (s *RPCServiceServerImpl) Pending(ctx context.Context, req *rpc.Empty) (*rpc.PendingList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pending not implemented")
}

func (s *RPCServiceServerImpl) Running(ctx context.Context, req *rpc.Empty) (*rpc.RunningList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Running not implemented")
}

func (s *RPCServiceServerImpl) Kill(ctx context.Context, req *rpc.KillRequest) (*rpc.ProcessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Kill not implemented")
}

func (s *RPCServiceServerImpl) KillAll(ctx context.Context, req *rpc.Empty) (*rpc.ProcessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KillAll not implemented")
}

func (s *RPCServiceServerImpl) Clear(ctx context.Context, req *rpc.ClearRequest) (*rpc.ProcessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Clear not implemented")
}

func (s *RPCServiceServerImpl) FreeSpace(ctx context.Context, req *rpc.Empty) (*rpc.FreeSpaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FreeSpace not implemented")
}

func (s *RPCServiceServerImpl) DirectoryTree(ctx context.Context, req *rpc.Empty) (*rpc.DirectoryTreeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DirectoryTree not implemented")
}

func (s *RPCServiceServerImpl) UpdateExecutable(ctx context.Context, req *rpc.Empty) (*rpc.UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateExecutable not implemented")
}
