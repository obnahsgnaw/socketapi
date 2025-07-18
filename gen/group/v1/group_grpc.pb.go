// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: group/v1/group.proto

package groupv1

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
	GroupService_JoinGroup_FullMethodName      = "/group.v1.GroupService/JoinGroup"
	GroupService_LeaveGroup_FullMethodName     = "/group.v1.GroupService/LeaveGroup"
	GroupService_BroadcastGroup_FullMethodName = "/group.v1.GroupService/BroadcastGroup"
)

// GroupServiceClient is the client API for GroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupServiceClient interface {
	// join a group
	JoinGroup(ctx context.Context, in *JoinGroupRequest, opts ...grpc.CallOption) (*JoinGroupResponse, error)
	// leave a group
	LeaveGroup(ctx context.Context, in *LeaveGroupRequest, opts ...grpc.CallOption) (*LeaveGroupResponse, error)
	// broadcast in a group
	BroadcastGroup(ctx context.Context, in *BroadcastGroupRequest, opts ...grpc.CallOption) (*BroadcastGroupResponse, error)
}

type groupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupServiceClient(cc grpc.ClientConnInterface) GroupServiceClient {
	return &groupServiceClient{cc}
}

func (c *groupServiceClient) JoinGroup(ctx context.Context, in *JoinGroupRequest, opts ...grpc.CallOption) (*JoinGroupResponse, error) {
	out := new(JoinGroupResponse)
	err := c.cc.Invoke(ctx, GroupService_JoinGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) LeaveGroup(ctx context.Context, in *LeaveGroupRequest, opts ...grpc.CallOption) (*LeaveGroupResponse, error) {
	out := new(LeaveGroupResponse)
	err := c.cc.Invoke(ctx, GroupService_LeaveGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) BroadcastGroup(ctx context.Context, in *BroadcastGroupRequest, opts ...grpc.CallOption) (*BroadcastGroupResponse, error) {
	out := new(BroadcastGroupResponse)
	err := c.cc.Invoke(ctx, GroupService_BroadcastGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupServiceServer is the server API for GroupService service.
// All implementations must embed UnimplementedGroupServiceServer
// for forward compatibility
type GroupServiceServer interface {
	// join a group
	JoinGroup(context.Context, *JoinGroupRequest) (*JoinGroupResponse, error)
	// leave a group
	LeaveGroup(context.Context, *LeaveGroupRequest) (*LeaveGroupResponse, error)
	// broadcast in a group
	BroadcastGroup(context.Context, *BroadcastGroupRequest) (*BroadcastGroupResponse, error)
	mustEmbedUnimplementedGroupServiceServer()
}

// UnimplementedGroupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGroupServiceServer struct {
}

func (UnimplementedGroupServiceServer) JoinGroup(context.Context, *JoinGroupRequest) (*JoinGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinGroup not implemented")
}
func (UnimplementedGroupServiceServer) LeaveGroup(context.Context, *LeaveGroupRequest) (*LeaveGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveGroup not implemented")
}
func (UnimplementedGroupServiceServer) BroadcastGroup(context.Context, *BroadcastGroupRequest) (*BroadcastGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastGroup not implemented")
}
func (UnimplementedGroupServiceServer) mustEmbedUnimplementedGroupServiceServer() {}

// UnsafeGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupServiceServer will
// result in compilation errors.
type UnsafeGroupServiceServer interface {
	mustEmbedUnimplementedGroupServiceServer()
}

func RegisterGroupServiceServer(s grpc.ServiceRegistrar, srv GroupServiceServer) {
	s.RegisterService(&GroupService_ServiceDesc, srv)
}

func _GroupService_JoinGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).JoinGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_JoinGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).JoinGroup(ctx, req.(*JoinGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_LeaveGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).LeaveGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_LeaveGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).LeaveGroup(ctx, req.(*LeaveGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_BroadcastGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).BroadcastGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_BroadcastGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).BroadcastGroup(ctx, req.(*BroadcastGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GroupService_ServiceDesc is the grpc.ServiceDesc for GroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "group.v1.GroupService",
	HandlerType: (*GroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinGroup",
			Handler:    _GroupService_JoinGroup_Handler,
		},
		{
			MethodName: "LeaveGroup",
			Handler:    _GroupService_LeaveGroup_Handler,
		},
		{
			MethodName: "BroadcastGroup",
			Handler:    _GroupService_BroadcastGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "group/v1/group.proto",
}
