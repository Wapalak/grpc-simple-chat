// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: chat.proto

package api

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

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatClient interface {
	AddChat(ctx context.Context, in *ChatId, opts ...grpc.CallOption) (*ChatId, error)
	AddMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MessageResponce, error)
	GetChat(ctx context.Context, in *ChatId, opts ...grpc.CallOption) (*GetChatResponse, error)
	StreamChat(ctx context.Context, in *ChatId, opts ...grpc.CallOption) (Chat_StreamChatClient, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) AddChat(ctx context.Context, in *ChatId, opts ...grpc.CallOption) (*ChatId, error) {
	out := new(ChatId)
	err := c.cc.Invoke(ctx, "/api.Chat/AddChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) AddMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MessageResponce, error) {
	out := new(MessageResponce)
	err := c.cc.Invoke(ctx, "/api.Chat/AddMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) GetChat(ctx context.Context, in *ChatId, opts ...grpc.CallOption) (*GetChatResponse, error) {
	out := new(GetChatResponse)
	err := c.cc.Invoke(ctx, "/api.Chat/GetChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) StreamChat(ctx context.Context, in *ChatId, opts ...grpc.CallOption) (Chat_StreamChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chat_ServiceDesc.Streams[0], "/api.Chat/StreamChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatStreamChatClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chat_StreamChatClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type chatStreamChatClient struct {
	grpc.ClientStream
}

func (x *chatStreamChatClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServer is the server API for Chat service.
// All implementations must embed UnimplementedChatServer
// for forward compatibility
type ChatServer interface {
	AddChat(context.Context, *ChatId) (*ChatId, error)
	AddMessage(context.Context, *Message) (*MessageResponce, error)
	GetChat(context.Context, *ChatId) (*GetChatResponse, error)
	StreamChat(*ChatId, Chat_StreamChatServer) error
	mustEmbedUnimplementedChatServer()
}

// UnimplementedChatServer must be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (UnimplementedChatServer) AddChat(context.Context, *ChatId) (*ChatId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddChat not implemented")
}
func (UnimplementedChatServer) AddMessage(context.Context, *Message) (*MessageResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMessage not implemented")
}
func (UnimplementedChatServer) GetChat(context.Context, *ChatId) (*GetChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChat not implemented")
}
func (UnimplementedChatServer) StreamChat(*ChatId, Chat_StreamChatServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamChat not implemented")
}
func (UnimplementedChatServer) mustEmbedUnimplementedChatServer() {}

// UnsafeChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServer will
// result in compilation errors.
type UnsafeChatServer interface {
	mustEmbedUnimplementedChatServer()
}

func RegisterChatServer(s grpc.ServiceRegistrar, srv ChatServer) {
	s.RegisterService(&Chat_ServiceDesc, srv)
}

func _Chat_AddChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).AddChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Chat/AddChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).AddChat(ctx, req.(*ChatId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_AddMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).AddMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Chat/AddMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).AddMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_GetChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).GetChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Chat/GetChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).GetChat(ctx, req.(*ChatId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_StreamChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ChatId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServer).StreamChat(m, &chatStreamChatServer{stream})
}

type Chat_StreamChatServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type chatStreamChatServer struct {
	grpc.ServerStream
}

func (x *chatStreamChatServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

// Chat_ServiceDesc is the grpc.ServiceDesc for Chat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddChat",
			Handler:    _Chat_AddChat_Handler,
		},
		{
			MethodName: "AddMessage",
			Handler:    _Chat_AddMessage_Handler,
		},
		{
			MethodName: "GetChat",
			Handler:    _Chat_GetChat_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamChat",
			Handler:       _Chat_StreamChat_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chat.proto",
}
