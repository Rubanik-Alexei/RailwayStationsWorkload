// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protobuff

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

// WorkloadServiceClient is the client API for WorkloadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkloadServiceClient interface {
	GetStationWorkload(ctx context.Context, in *GetStationWorkloadRequest, opts ...grpc.CallOption) (WorkloadService_GetStationWorkloadClient, error)
}

type workloadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkloadServiceClient(cc grpc.ClientConnInterface) WorkloadServiceClient {
	return &workloadServiceClient{cc}
}

func (c *workloadServiceClient) GetStationWorkload(ctx context.Context, in *GetStationWorkloadRequest, opts ...grpc.CallOption) (WorkloadService_GetStationWorkloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &WorkloadService_ServiceDesc.Streams[0], "/WorkloadService/GetStationWorkload", opts...)
	if err != nil {
		return nil, err
	}
	x := &workloadServiceGetStationWorkloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WorkloadService_GetStationWorkloadClient interface {
	Recv() (*StationData, error)
	grpc.ClientStream
}

type workloadServiceGetStationWorkloadClient struct {
	grpc.ClientStream
}

func (x *workloadServiceGetStationWorkloadClient) Recv() (*StationData, error) {
	m := new(StationData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WorkloadServiceServer is the server API for WorkloadService service.
// All implementations must embed UnimplementedWorkloadServiceServer
// for forward compatibility
type WorkloadServiceServer interface {
	GetStationWorkload(*GetStationWorkloadRequest, WorkloadService_GetStationWorkloadServer) error
	mustEmbedUnimplementedWorkloadServiceServer()
}

// UnimplementedWorkloadServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkloadServiceServer struct {
}

func (UnimplementedWorkloadServiceServer) GetStationWorkload(*GetStationWorkloadRequest, WorkloadService_GetStationWorkloadServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStationWorkload not implemented")
}
func (UnimplementedWorkloadServiceServer) mustEmbedUnimplementedWorkloadServiceServer() {}

// UnsafeWorkloadServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkloadServiceServer will
// result in compilation errors.
type UnsafeWorkloadServiceServer interface {
	mustEmbedUnimplementedWorkloadServiceServer()
}

func RegisterWorkloadServiceServer(s grpc.ServiceRegistrar, srv WorkloadServiceServer) {
	s.RegisterService(&WorkloadService_ServiceDesc, srv)
}

func _WorkloadService_GetStationWorkload_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetStationWorkloadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WorkloadServiceServer).GetStationWorkload(m, &workloadServiceGetStationWorkloadServer{stream})
}

type WorkloadService_GetStationWorkloadServer interface {
	Send(*StationData) error
	grpc.ServerStream
}

type workloadServiceGetStationWorkloadServer struct {
	grpc.ServerStream
}

func (x *workloadServiceGetStationWorkloadServer) Send(m *StationData) error {
	return x.ServerStream.SendMsg(m)
}

// WorkloadService_ServiceDesc is the grpc.ServiceDesc for WorkloadService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkloadService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "WorkloadService",
	HandlerType: (*WorkloadServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStationWorkload",
			Handler:       _WorkloadService_GetStationWorkload_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "workload_service/protobuff/wl.proto",
}
