// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: log.proto

package pb

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
	LogService_ListLog_FullMethodName   = "/log.v1.LogService/ListLog"
	LogService_CreateLog_FullMethodName = "/log.v1.LogService/CreateLog"
	LogService_UpdateLog_FullMethodName = "/log.v1.LogService/UpdateLog"
	LogService_DeleteLog_FullMethodName = "/log.v1.LogService/DeleteLog"
)

// LogServiceClient is the client API for LogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogServiceClient interface {
	ListLog(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryLogResp, error)
	CreateLog(ctx context.Context, in *Log, opts ...grpc.CallOption) (*RetInfo, error)
	UpdateLog(ctx context.Context, in *Log, opts ...grpc.CallOption) (*RetInfo, error)
	DeleteLog(ctx context.Context, in *DeleteIds, opts ...grpc.CallOption) (*RetInfo, error)
}

type logServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogServiceClient(cc grpc.ClientConnInterface) LogServiceClient {
	return &logServiceClient{cc}
}

func (c *logServiceClient) ListLog(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryLogResp, error) {
	out := new(QueryLogResp)
	err := c.cc.Invoke(ctx, LogService_ListLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) CreateLog(ctx context.Context, in *Log, opts ...grpc.CallOption) (*RetInfo, error) {
	out := new(RetInfo)
	err := c.cc.Invoke(ctx, LogService_CreateLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) UpdateLog(ctx context.Context, in *Log, opts ...grpc.CallOption) (*RetInfo, error) {
	out := new(RetInfo)
	err := c.cc.Invoke(ctx, LogService_UpdateLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) DeleteLog(ctx context.Context, in *DeleteIds, opts ...grpc.CallOption) (*RetInfo, error) {
	out := new(RetInfo)
	err := c.cc.Invoke(ctx, LogService_DeleteLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogServiceServer is the server API for LogService service.
// All implementations must embed UnimplementedLogServiceServer
// for forward compatibility
type LogServiceServer interface {
	ListLog(context.Context, *QueryRequest) (*QueryLogResp, error)
	CreateLog(context.Context, *Log) (*RetInfo, error)
	UpdateLog(context.Context, *Log) (*RetInfo, error)
	DeleteLog(context.Context, *DeleteIds) (*RetInfo, error)
	mustEmbedUnimplementedLogServiceServer()
}

// UnimplementedLogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLogServiceServer struct {
}

func (UnimplementedLogServiceServer) ListLog(context.Context, *QueryRequest) (*QueryLogResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLog not implemented")
}
func (UnimplementedLogServiceServer) CreateLog(context.Context, *Log) (*RetInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLog not implemented")
}
func (UnimplementedLogServiceServer) UpdateLog(context.Context, *Log) (*RetInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLog not implemented")
}
func (UnimplementedLogServiceServer) DeleteLog(context.Context, *DeleteIds) (*RetInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLog not implemented")
}
func (UnimplementedLogServiceServer) mustEmbedUnimplementedLogServiceServer() {}

// UnsafeLogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogServiceServer will
// result in compilation errors.
type UnsafeLogServiceServer interface {
	mustEmbedUnimplementedLogServiceServer()
}

func RegisterLogServiceServer(s grpc.ServiceRegistrar, srv LogServiceServer) {
	s.RegisterService(&LogService_ServiceDesc, srv)
}

func _LogService_ListLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).ListLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogService_ListLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).ListLog(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_CreateLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Log)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).CreateLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogService_CreateLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).CreateLog(ctx, req.(*Log))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_UpdateLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Log)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).UpdateLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogService_UpdateLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).UpdateLog(ctx, req.(*Log))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_DeleteLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).DeleteLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogService_DeleteLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).DeleteLog(ctx, req.(*DeleteIds))
	}
	return interceptor(ctx, in, info, handler)
}

// LogService_ServiceDesc is the grpc.ServiceDesc for LogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "log.v1.LogService",
	HandlerType: (*LogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListLog",
			Handler:    _LogService_ListLog_Handler,
		},
		{
			MethodName: "CreateLog",
			Handler:    _LogService_CreateLog_Handler,
		},
		{
			MethodName: "UpdateLog",
			Handler:    _LogService_UpdateLog_Handler,
		},
		{
			MethodName: "DeleteLog",
			Handler:    _LogService_DeleteLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "log.proto",
}
