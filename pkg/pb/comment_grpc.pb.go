// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: comment.proto

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

// CommentServiceClient is the client API for CommentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentServiceClient interface {
	CommentTopAction(ctx context.Context, in *CommentActionTopReq, opts ...grpc.CallOption) (*CommentActionRsp, error)
	CommentOtherAction(ctx context.Context, in *CommentActionOtherReq, opts ...grpc.CallOption) (*CommentActionRsp, error)
	GetTopCommentList(ctx context.Context, in *GetTopCommentListReq, opts ...grpc.CallOption) (*GetTopCommentListRsp, error)
	GetOtherCommentList(ctx context.Context, in *GetOtherCommentListReq, opts ...grpc.CallOption) (*GetOtherCommentListRsp, error)
	GetCommentSum(ctx context.Context, in *GetCommentNumReq, opts ...grpc.CallOption) (*GetCommentNumRsp, error)
}

type commentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentServiceClient(cc grpc.ClientConnInterface) CommentServiceClient {
	return &commentServiceClient{cc}
}

func (c *commentServiceClient) CommentTopAction(ctx context.Context, in *CommentActionTopReq, opts ...grpc.CallOption) (*CommentActionRsp, error) {
	out := new(CommentActionRsp)
	err := c.cc.Invoke(ctx, "/CommentService/CommentTopAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) CommentOtherAction(ctx context.Context, in *CommentActionOtherReq, opts ...grpc.CallOption) (*CommentActionRsp, error) {
	out := new(CommentActionRsp)
	err := c.cc.Invoke(ctx, "/CommentService/CommentOtherAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) GetTopCommentList(ctx context.Context, in *GetTopCommentListReq, opts ...grpc.CallOption) (*GetTopCommentListRsp, error) {
	out := new(GetTopCommentListRsp)
	err := c.cc.Invoke(ctx, "/CommentService/GetTopCommentList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) GetOtherCommentList(ctx context.Context, in *GetOtherCommentListReq, opts ...grpc.CallOption) (*GetOtherCommentListRsp, error) {
	out := new(GetOtherCommentListRsp)
	err := c.cc.Invoke(ctx, "/CommentService/GetOtherCommentList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) GetCommentSum(ctx context.Context, in *GetCommentNumReq, opts ...grpc.CallOption) (*GetCommentNumRsp, error) {
	out := new(GetCommentNumRsp)
	err := c.cc.Invoke(ctx, "/CommentService/GetCommentSum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServiceServer is the server API for CommentService service.
// All implementations must embed UnimplementedCommentServiceServer
// for forward compatibility
type CommentServiceServer interface {
	CommentTopAction(context.Context, *CommentActionTopReq) (*CommentActionRsp, error)
	CommentOtherAction(context.Context, *CommentActionOtherReq) (*CommentActionRsp, error)
	GetTopCommentList(context.Context, *GetTopCommentListReq) (*GetTopCommentListRsp, error)
	GetOtherCommentList(context.Context, *GetOtherCommentListReq) (*GetOtherCommentListRsp, error)
	GetCommentSum(context.Context, *GetCommentNumReq) (*GetCommentNumRsp, error)
	mustEmbedUnimplementedCommentServiceServer()
}

// UnimplementedCommentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommentServiceServer struct {
}

func (UnimplementedCommentServiceServer) CommentTopAction(context.Context, *CommentActionTopReq) (*CommentActionRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentTopAction not implemented")
}
func (UnimplementedCommentServiceServer) CommentOtherAction(context.Context, *CommentActionOtherReq) (*CommentActionRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentOtherAction not implemented")
}
func (UnimplementedCommentServiceServer) GetTopCommentList(context.Context, *GetTopCommentListReq) (*GetTopCommentListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopCommentList not implemented")
}
func (UnimplementedCommentServiceServer) GetOtherCommentList(context.Context, *GetOtherCommentListReq) (*GetOtherCommentListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOtherCommentList not implemented")
}
func (UnimplementedCommentServiceServer) GetCommentSum(context.Context, *GetCommentNumReq) (*GetCommentNumRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentSum not implemented")
}
func (UnimplementedCommentServiceServer) mustEmbedUnimplementedCommentServiceServer() {}

// UnsafeCommentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentServiceServer will
// result in compilation errors.
type UnsafeCommentServiceServer interface {
	mustEmbedUnimplementedCommentServiceServer()
}

func RegisterCommentServiceServer(s grpc.ServiceRegistrar, srv CommentServiceServer) {
	s.RegisterService(&CommentService_ServiceDesc, srv)
}

func _CommentService_CommentTopAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentActionTopReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).CommentTopAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CommentService/CommentTopAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).CommentTopAction(ctx, req.(*CommentActionTopReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_CommentOtherAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentActionOtherReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).CommentOtherAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CommentService/CommentOtherAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).CommentOtherAction(ctx, req.(*CommentActionOtherReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_GetTopCommentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopCommentListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetTopCommentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CommentService/GetTopCommentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetTopCommentList(ctx, req.(*GetTopCommentListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_GetOtherCommentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOtherCommentListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetOtherCommentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CommentService/GetOtherCommentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetOtherCommentList(ctx, req.(*GetOtherCommentListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_GetCommentSum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentNumReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetCommentSum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CommentService/GetCommentSum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetCommentSum(ctx, req.(*GetCommentNumReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CommentService_ServiceDesc is the grpc.ServiceDesc for CommentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CommentService",
	HandlerType: (*CommentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CommentTopAction",
			Handler:    _CommentService_CommentTopAction_Handler,
		},
		{
			MethodName: "CommentOtherAction",
			Handler:    _CommentService_CommentOtherAction_Handler,
		},
		{
			MethodName: "GetTopCommentList",
			Handler:    _CommentService_GetTopCommentList_Handler,
		},
		{
			MethodName: "GetOtherCommentList",
			Handler:    _CommentService_GetOtherCommentList_Handler,
		},
		{
			MethodName: "GetCommentSum",
			Handler:    _CommentService_GetCommentSum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment.proto",
}
