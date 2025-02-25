// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package types

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

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	MakePool(ctx context.Context, in *MsgMakePoolRequest, opts ...grpc.CallOption) (*MsgMakePoolResponse, error)
	TakePool(ctx context.Context, in *MsgTakePoolRequest, opts ...grpc.CallOption) (*MsgTakePoolResponse, error)
	SingleAssetDeposit(ctx context.Context, in *MsgSingleAssetDepositRequest, opts ...grpc.CallOption) (*MsgSingleAssetDepositResponse, error)
	MakeMultiAssetDeposit(ctx context.Context, in *MsgMakeMultiAssetDepositRequest, opts ...grpc.CallOption) (*MsgMultiAssetDepositResponse, error)
	TakeMultiAssetDeposit(ctx context.Context, in *MsgTakeMultiAssetDepositRequest, opts ...grpc.CallOption) (*MsgMultiAssetDepositResponse, error)
	MultiAssetWithdraw(ctx context.Context, in *MsgMultiAssetWithdrawRequest, opts ...grpc.CallOption) (*MsgMultiAssetWithdrawResponse, error)
	Swap(ctx context.Context, in *MsgSwapRequest, opts ...grpc.CallOption) (*MsgSwapResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) MakePool(ctx context.Context, in *MsgMakePoolRequest, opts ...grpc.CallOption) (*MsgMakePoolResponse, error) {
	out := new(MsgMakePoolResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.interchain_swap.v1.Msg/MakePool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) TakePool(ctx context.Context, in *MsgTakePoolRequest, opts ...grpc.CallOption) (*MsgTakePoolResponse, error) {
	out := new(MsgTakePoolResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.interchain_swap.v1.Msg/TakePool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SingleAssetDeposit(ctx context.Context, in *MsgSingleAssetDepositRequest, opts ...grpc.CallOption) (*MsgSingleAssetDepositResponse, error) {
	out := new(MsgSingleAssetDepositResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.interchain_swap.v1.Msg/SingleAssetDeposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MakeMultiAssetDeposit(ctx context.Context, in *MsgMakeMultiAssetDepositRequest, opts ...grpc.CallOption) (*MsgMultiAssetDepositResponse, error) {
	out := new(MsgMultiAssetDepositResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.interchain_swap.v1.Msg/MakeMultiAssetDeposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) TakeMultiAssetDeposit(ctx context.Context, in *MsgTakeMultiAssetDepositRequest, opts ...grpc.CallOption) (*MsgMultiAssetDepositResponse, error) {
	out := new(MsgMultiAssetDepositResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.interchain_swap.v1.Msg/TakeMultiAssetDeposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MultiAssetWithdraw(ctx context.Context, in *MsgMultiAssetWithdrawRequest, opts ...grpc.CallOption) (*MsgMultiAssetWithdrawResponse, error) {
	out := new(MsgMultiAssetWithdrawResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.interchain_swap.v1.Msg/MultiAssetWithdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Swap(ctx context.Context, in *MsgSwapRequest, opts ...grpc.CallOption) (*MsgSwapResponse, error) {
	out := new(MsgSwapResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.interchain_swap.v1.Msg/Swap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations should embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	MakePool(context.Context, *MsgMakePoolRequest) (*MsgMakePoolResponse, error)
	TakePool(context.Context, *MsgTakePoolRequest) (*MsgTakePoolResponse, error)
	SingleAssetDeposit(context.Context, *MsgSingleAssetDepositRequest) (*MsgSingleAssetDepositResponse, error)
	MakeMultiAssetDeposit(context.Context, *MsgMakeMultiAssetDepositRequest) (*MsgMultiAssetDepositResponse, error)
	TakeMultiAssetDeposit(context.Context, *MsgTakeMultiAssetDepositRequest) (*MsgMultiAssetDepositResponse, error)
	MultiAssetWithdraw(context.Context, *MsgMultiAssetWithdrawRequest) (*MsgMultiAssetWithdrawResponse, error)
	Swap(context.Context, *MsgSwapRequest) (*MsgSwapResponse, error)
}

// UnimplementedMsgServer should be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) MakePool(context.Context, *MsgMakePoolRequest) (*MsgMakePoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakePool not implemented")
}
func (UnimplementedMsgServer) TakePool(context.Context, *MsgTakePoolRequest) (*MsgTakePoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TakePool not implemented")
}
func (UnimplementedMsgServer) SingleAssetDeposit(context.Context, *MsgSingleAssetDepositRequest) (*MsgSingleAssetDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SingleAssetDeposit not implemented")
}
func (UnimplementedMsgServer) MakeMultiAssetDeposit(context.Context, *MsgMakeMultiAssetDepositRequest) (*MsgMultiAssetDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeMultiAssetDeposit not implemented")
}
func (UnimplementedMsgServer) TakeMultiAssetDeposit(context.Context, *MsgTakeMultiAssetDepositRequest) (*MsgMultiAssetDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TakeMultiAssetDeposit not implemented")
}
func (UnimplementedMsgServer) MultiAssetWithdraw(context.Context, *MsgMultiAssetWithdrawRequest) (*MsgMultiAssetWithdrawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MultiAssetWithdraw not implemented")
}
func (UnimplementedMsgServer) Swap(context.Context, *MsgSwapRequest) (*MsgSwapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Swap not implemented")
}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_MakePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMakePoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MakePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.interchain_swap.v1.Msg/MakePool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MakePool(ctx, req.(*MsgMakePoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_TakePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTakePoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).TakePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.interchain_swap.v1.Msg/TakePool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).TakePool(ctx, req.(*MsgTakePoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SingleAssetDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSingleAssetDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SingleAssetDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.interchain_swap.v1.Msg/SingleAssetDeposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SingleAssetDeposit(ctx, req.(*MsgSingleAssetDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MakeMultiAssetDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMakeMultiAssetDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MakeMultiAssetDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.interchain_swap.v1.Msg/MakeMultiAssetDeposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MakeMultiAssetDeposit(ctx, req.(*MsgMakeMultiAssetDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_TakeMultiAssetDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTakeMultiAssetDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).TakeMultiAssetDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.interchain_swap.v1.Msg/TakeMultiAssetDeposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).TakeMultiAssetDeposit(ctx, req.(*MsgTakeMultiAssetDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MultiAssetWithdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMultiAssetWithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MultiAssetWithdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.interchain_swap.v1.Msg/MultiAssetWithdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MultiAssetWithdraw(ctx, req.(*MsgMultiAssetWithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Swap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSwapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Swap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.interchain_swap.v1.Msg/Swap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Swap(ctx, req.(*MsgSwapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ibc.applications.interchain_swap.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MakePool",
			Handler:    _Msg_MakePool_Handler,
		},
		{
			MethodName: "TakePool",
			Handler:    _Msg_TakePool_Handler,
		},
		{
			MethodName: "SingleAssetDeposit",
			Handler:    _Msg_SingleAssetDeposit_Handler,
		},
		{
			MethodName: "MakeMultiAssetDeposit",
			Handler:    _Msg_MakeMultiAssetDeposit_Handler,
		},
		{
			MethodName: "TakeMultiAssetDeposit",
			Handler:    _Msg_TakeMultiAssetDeposit_Handler,
		},
		{
			MethodName: "MultiAssetWithdraw",
			Handler:    _Msg_MultiAssetWithdraw_Handler,
		},
		{
			MethodName: "Swap",
			Handler:    _Msg_Swap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc/applications/interchain_swap/v1/tx.proto",
}
