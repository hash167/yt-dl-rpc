// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: rpc.proto

package rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RPCService_Exec_FullMethodName             = "/rpc.RPCService/Exec"
	RPCService_ExecPlaylist_FullMethodName     = "/rpc.RPCService/ExecPlaylist"
	RPCService_Progress_FullMethodName         = "/rpc.RPCService/Progress"
	RPCService_Formats_FullMethodName          = "/rpc.RPCService/Formats"
	RPCService_Pending_FullMethodName          = "/rpc.RPCService/Pending"
	RPCService_Running_FullMethodName          = "/rpc.RPCService/Running"
	RPCService_Kill_FullMethodName             = "/rpc.RPCService/Kill"
	RPCService_KillAll_FullMethodName          = "/rpc.RPCService/KillAll"
	RPCService_Clear_FullMethodName            = "/rpc.RPCService/Clear"
	RPCService_FreeSpace_FullMethodName        = "/rpc.RPCService/FreeSpace"
	RPCService_DirectoryTree_FullMethodName    = "/rpc.RPCService/DirectoryTree"
	RPCService_UpdateExecutable_FullMethodName = "/rpc.RPCService/UpdateExecutable"
)

// RPCServiceClient is the client API for RPCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCServiceClient interface {
	Exec(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*ProcessResponse, error)
	ExecPlaylist(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*ProcessResponse, error)
	Progress(ctx context.Context, in *Args, opts ...grpc.CallOption) (*DownloadProgress, error)
	Formats(ctx context.Context, in *Args, opts ...grpc.CallOption) (*DownloadFormats, error)
	Pending(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PendingList, error)
	Running(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RunningList, error)
	Kill(ctx context.Context, in *KillRequest, opts ...grpc.CallOption) (*ProcessResponse, error)
	KillAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ProcessResponse, error)
	Clear(ctx context.Context, in *ClearRequest, opts ...grpc.CallOption) (*ProcessResponse, error)
	FreeSpace(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*FreeSpaceResponse, error)
	DirectoryTree(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DirectoryTreeResponse, error)
	UpdateExecutable(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UpdateResponse, error)
}

type rPCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCServiceClient(cc grpc.ClientConnInterface) RPCServiceClient {
	return &rPCServiceClient{cc}
}

func (c *rPCServiceClient) Exec(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*ProcessResponse, error) {
	out := new(ProcessResponse)
	err := c.cc.Invoke(ctx, RPCService_Exec_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCServiceClient) ExecPlaylist(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*ProcessResponse, error) {
	out := new(ProcessResponse)
	err := c.cc.Invoke(ctx, RPCService_ExecPlaylist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCServiceClient) Progress(ctx context.Context, in *Args, opts ...grpc.CallOption) (*DownloadProgress, error) {
	out := new(DownloadProgress)
	err := c.cc.Invoke(ctx, RPCService_Progress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCServiceClient) Formats(ctx context.Context, in *Args, opts ...grpc.CallOption) (*DownloadFormats, error) {
	out := new(DownloadFormats)
	err := c.cc.Invoke(ctx, RPCService_Formats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCServiceClient) Pending(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PendingList, error) {
	out := new(PendingList)
	err := c.cc.Invoke(ctx, RPCService_Pending_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCServiceClient) Running(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RunningList, error) {
	out := new(RunningList)
	err := c.cc.Invoke(ctx, RPCService_Running_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCServiceClient) Kill(ctx context.Context, in *KillRequest, opts ...grpc.CallOption) (*ProcessResponse, error) {
	out := new(ProcessResponse)
	err := c.cc.Invoke(ctx, RPCService_Kill_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCServiceClient) KillAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ProcessResponse, error) {
	out := new(ProcessResponse)
	err := c.cc.Invoke(ctx, RPCService_KillAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCServiceClient) Clear(ctx context.Context, in *ClearRequest, opts ...grpc.CallOption) (*ProcessResponse, error) {
	out := new(ProcessResponse)
	err := c.cc.Invoke(ctx, RPCService_Clear_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCServiceClient) FreeSpace(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*FreeSpaceResponse, error) {
	out := new(FreeSpaceResponse)
	err := c.cc.Invoke(ctx, RPCService_FreeSpace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCServiceClient) DirectoryTree(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DirectoryTreeResponse, error) {
	out := new(DirectoryTreeResponse)
	err := c.cc.Invoke(ctx, RPCService_DirectoryTree_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCServiceClient) UpdateExecutable(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, RPCService_UpdateExecutable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServiceServer is the server API for RPCService service.
// All implementations must embed UnimplementedRPCServiceServer
// for forward compatibility
type RPCServiceServer interface {
	Exec(context.Context, *DownloadRequest) (*ProcessResponse, error)
	ExecPlaylist(context.Context, *DownloadRequest) (*ProcessResponse, error)
	Progress(context.Context, *Args) (*DownloadProgress, error)
	Formats(context.Context, *Args) (*DownloadFormats, error)
	Pending(context.Context, *Empty) (*PendingList, error)
	Running(context.Context, *Empty) (*RunningList, error)
	Kill(context.Context, *KillRequest) (*ProcessResponse, error)
	KillAll(context.Context, *Empty) (*ProcessResponse, error)
	Clear(context.Context, *ClearRequest) (*ProcessResponse, error)
	FreeSpace(context.Context, *Empty) (*FreeSpaceResponse, error)
	DirectoryTree(context.Context, *Empty) (*DirectoryTreeResponse, error)
	UpdateExecutable(context.Context, *Empty) (*UpdateResponse, error)
	mustEmbedUnimplementedRPCServiceServer()
}

// UnimplementedRPCServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRPCServiceServer struct {
}

func (UnimplementedRPCServiceServer) Exec(context.Context, *DownloadRequest) (*ProcessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exec not implemented")
}
func (UnimplementedRPCServiceServer) ExecPlaylist(context.Context, *DownloadRequest) (*ProcessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecPlaylist not implemented")
}
func (UnimplementedRPCServiceServer) Progress(context.Context, *Args) (*DownloadProgress, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Progress not implemented")
}
func (UnimplementedRPCServiceServer) Formats(context.Context, *Args) (*DownloadFormats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Formats not implemented")
}
func (UnimplementedRPCServiceServer) Pending(context.Context, *Empty) (*PendingList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pending not implemented")
}
func (UnimplementedRPCServiceServer) Running(context.Context, *Empty) (*RunningList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Running not implemented")
}
func (UnimplementedRPCServiceServer) Kill(context.Context, *KillRequest) (*ProcessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Kill not implemented")
}
func (UnimplementedRPCServiceServer) KillAll(context.Context, *Empty) (*ProcessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KillAll not implemented")
}
func (UnimplementedRPCServiceServer) Clear(context.Context, *ClearRequest) (*ProcessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Clear not implemented")
}
func (UnimplementedRPCServiceServer) FreeSpace(context.Context, *Empty) (*FreeSpaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FreeSpace not implemented")
}
func (UnimplementedRPCServiceServer) DirectoryTree(context.Context, *Empty) (*DirectoryTreeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DirectoryTree not implemented")
}
func (UnimplementedRPCServiceServer) UpdateExecutable(context.Context, *Empty) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateExecutable not implemented")
}
func (UnimplementedRPCServiceServer) mustEmbedUnimplementedRPCServiceServer() {}

// UnsafeRPCServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCServiceServer will
// result in compilation errors.
type UnsafeRPCServiceServer interface {
	mustEmbedUnimplementedRPCServiceServer()
}

func RegisterRPCServiceServer(s grpc.ServiceRegistrar, srv RPCServiceServer) {
	s.RegisterService(&RPCService_ServiceDesc, srv)
}

func _RPCService_Exec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).Exec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCService_Exec_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).Exec(ctx, req.(*DownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCService_ExecPlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).ExecPlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCService_ExecPlaylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).ExecPlaylist(ctx, req.(*DownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCService_Progress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Args)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).Progress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCService_Progress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).Progress(ctx, req.(*Args))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCService_Formats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Args)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).Formats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCService_Formats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).Formats(ctx, req.(*Args))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCService_Pending_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).Pending(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCService_Pending_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).Pending(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCService_Running_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).Running(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCService_Running_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).Running(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCService_Kill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).Kill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCService_Kill_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).Kill(ctx, req.(*KillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCService_KillAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).KillAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCService_KillAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).KillAll(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCService_Clear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).Clear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCService_Clear_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).Clear(ctx, req.(*ClearRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCService_FreeSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).FreeSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCService_FreeSpace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).FreeSpace(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCService_DirectoryTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).DirectoryTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCService_DirectoryTree_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).DirectoryTree(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCService_UpdateExecutable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).UpdateExecutable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCService_UpdateExecutable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).UpdateExecutable(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// RPCService_ServiceDesc is the grpc.ServiceDesc for RPCService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPCService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.RPCService",
	HandlerType: (*RPCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Exec",
			Handler:    _RPCService_Exec_Handler,
		},
		{
			MethodName: "ExecPlaylist",
			Handler:    _RPCService_ExecPlaylist_Handler,
		},
		{
			MethodName: "Progress",
			Handler:    _RPCService_Progress_Handler,
		},
		{
			MethodName: "Formats",
			Handler:    _RPCService_Formats_Handler,
		},
		{
			MethodName: "Pending",
			Handler:    _RPCService_Pending_Handler,
		},
		{
			MethodName: "Running",
			Handler:    _RPCService_Running_Handler,
		},
		{
			MethodName: "Kill",
			Handler:    _RPCService_Kill_Handler,
		},
		{
			MethodName: "KillAll",
			Handler:    _RPCService_KillAll_Handler,
		},
		{
			MethodName: "Clear",
			Handler:    _RPCService_Clear_Handler,
		},
		{
			MethodName: "FreeSpace",
			Handler:    _RPCService_FreeSpace_Handler,
		},
		{
			MethodName: "DirectoryTree",
			Handler:    _RPCService_DirectoryTree_Handler,
		},
		{
			MethodName: "UpdateExecutable",
			Handler:    _RPCService_UpdateExecutable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}
