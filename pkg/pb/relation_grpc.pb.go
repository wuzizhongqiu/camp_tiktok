// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: relationsvr.proto

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

// RelationServiceClient is the client API for RelationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RelationServiceClient interface {
	RelationAction(ctx context.Context, in *RelationActionReq, opts ...grpc.CallOption) (*RelationActionRsp, error)
	GetRelationFollowList(ctx context.Context, in *GetRelationFollowListReq, opts ...grpc.CallOption) (*GetRelationFollowListRsp, error)
	GetRelationFollowerList(ctx context.Context, in *GetRelationFollowerListReq, opts ...grpc.CallOption) (*GetRelationFollowerListRsp, error)
	IsFollowDict(ctx context.Context, in *IsFollowDictReq, opts ...grpc.CallOption) (*IsFollowDictRsp, error)
	GetUserFollowerNum(ctx context.Context, in *UserFollowerNumReq, opts ...grpc.CallOption) (*UserFollowerNumRsp, error)
	GetUserFollowNum(ctx context.Context, in *UserFollowNumReq, opts ...grpc.CallOption) (*UserFollowNumRsp, error)
}

type relationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRelationServiceClient(cc grpc.ClientConnInterface) RelationServiceClient {
	return &relationServiceClient{cc}
}

func (c *relationServiceClient) RelationAction(ctx context.Context, in *RelationActionReq, opts ...grpc.CallOption) (*RelationActionRsp, error) {
	out := new(RelationActionRsp)
	err := c.cc.Invoke(ctx, "/RelationService/RelationAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) GetRelationFollowList(ctx context.Context, in *GetRelationFollowListReq, opts ...grpc.CallOption) (*GetRelationFollowListRsp, error) {
	out := new(GetRelationFollowListRsp)
	err := c.cc.Invoke(ctx, "/RelationService/GetRelationFollowList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) GetRelationFollowerList(ctx context.Context, in *GetRelationFollowerListReq, opts ...grpc.CallOption) (*GetRelationFollowerListRsp, error) {
	out := new(GetRelationFollowerListRsp)
	err := c.cc.Invoke(ctx, "/RelationService/GetRelationFollowerList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) IsFollowDict(ctx context.Context, in *IsFollowDictReq, opts ...grpc.CallOption) (*IsFollowDictRsp, error) {
	out := new(IsFollowDictRsp)
	err := c.cc.Invoke(ctx, "/RelationService/IsFollowDict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) GetUserFollowerNum(ctx context.Context, in *UserFollowerNumReq, opts ...grpc.CallOption) (*UserFollowerNumRsp, error) {
	out := new(UserFollowerNumRsp)
	err := c.cc.Invoke(ctx, "/RelationService/GetUserFollowerNum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) GetUserFollowNum(ctx context.Context, in *UserFollowNumReq, opts ...grpc.CallOption) (*UserFollowNumRsp, error) {
	out := new(UserFollowNumRsp)
	err := c.cc.Invoke(ctx, "/RelationService/GetUserFollowNum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelationServiceServer is the server API for RelationService service.
// All implementations must embed UnimplementedRelationServiceServer
// for forward compatibility
type RelationServiceServer interface {
	RelationAction(context.Context, *RelationActionReq) (*RelationActionRsp, error)
	GetRelationFollowList(context.Context, *GetRelationFollowListReq) (*GetRelationFollowListRsp, error)
	GetRelationFollowerList(context.Context, *GetRelationFollowerListReq) (*GetRelationFollowerListRsp, error)
	IsFollowDict(context.Context, *IsFollowDictReq) (*IsFollowDictRsp, error)
	GetUserFollowerNum(context.Context, *UserFollowerNumReq) (*UserFollowerNumRsp, error)
	GetUserFollowNum(context.Context, *UserFollowNumReq) (*UserFollowNumRsp, error)
	mustEmbedUnimplementedRelationServiceServer()
}

// UnimplementedRelationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRelationServiceServer struct {
}

func (UnimplementedRelationServiceServer) RelationAction(context.Context, *RelationActionReq) (*RelationActionRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RelationAction not implemented")
}
func (UnimplementedRelationServiceServer) GetRelationFollowList(context.Context, *GetRelationFollowListReq) (*GetRelationFollowListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRelationFollowList not implemented")
}
func (UnimplementedRelationServiceServer) GetRelationFollowerList(context.Context, *GetRelationFollowerListReq) (*GetRelationFollowerListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRelationFollowerList not implemented")
}
func (UnimplementedRelationServiceServer) IsFollowDict(context.Context, *IsFollowDictReq) (*IsFollowDictRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsFollowDict not implemented")
}
func (UnimplementedRelationServiceServer) GetUserFollowerNum(context.Context, *UserFollowerNumReq) (*UserFollowerNumRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFollowerNum not implemented")
}
func (UnimplementedRelationServiceServer) GetUserFollowNum(context.Context, *UserFollowNumReq) (*UserFollowNumRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFollowNum not implemented")
}
func (UnimplementedRelationServiceServer) mustEmbedUnimplementedRelationServiceServer() {}

// UnsafeRelationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RelationServiceServer will
// result in compilation errors.
type UnsafeRelationServiceServer interface {
	mustEmbedUnimplementedRelationServiceServer()
}

func RegisterRelationServiceServer(s grpc.ServiceRegistrar, srv RelationServiceServer) {
	s.RegisterService(&RelationService_ServiceDesc, srv)
}

func _RelationService_RelationAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelationActionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).RelationAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RelationService/RelationAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).RelationAction(ctx, req.(*RelationActionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_GetRelationFollowList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRelationFollowListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).GetRelationFollowList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RelationService/GetRelationFollowList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).GetRelationFollowList(ctx, req.(*GetRelationFollowListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_GetRelationFollowerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRelationFollowerListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).GetRelationFollowerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RelationService/GetRelationFollowerList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).GetRelationFollowerList(ctx, req.(*GetRelationFollowerListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_IsFollowDict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsFollowDictReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).IsFollowDict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RelationService/IsFollowDict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).IsFollowDict(ctx, req.(*IsFollowDictReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_GetUserFollowerNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserFollowerNumReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).GetUserFollowerNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RelationService/GetUserFollowerNum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).GetUserFollowerNum(ctx, req.(*UserFollowerNumReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_GetUserFollowNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserFollowNumReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).GetUserFollowNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RelationService/GetUserFollowNum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).GetUserFollowNum(ctx, req.(*UserFollowNumReq))
	}
	return interceptor(ctx, in, info, handler)
}

// RelationService_ServiceDesc is the grpc.ServiceDesc for RelationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RelationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RelationService",
	HandlerType: (*RelationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RelationAction",
			Handler:    _RelationService_RelationAction_Handler,
		},
		{
			MethodName: "GetRelationFollowList",
			Handler:    _RelationService_GetRelationFollowList_Handler,
		},
		{
			MethodName: "GetRelationFollowerList",
			Handler:    _RelationService_GetRelationFollowerList_Handler,
		},
		{
			MethodName: "IsFollowDict",
			Handler:    _RelationService_IsFollowDict_Handler,
		},
		{
			MethodName: "GetUserFollowerNum",
			Handler:    _RelationService_GetUserFollowerNum_Handler,
		},
		{
			MethodName: "GetUserFollowNum",
			Handler:    _RelationService_GetUserFollowNum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "relationsvr.proto",
}
