// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/user.proto

package proto

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	//rpc Signup(SignupRequest) returns (google.protobuf.Empty);
	//unary
	Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error)
	//server streaming
	GetPosts(ctx context.Context, in *GetPostsRequest, opts ...grpc.CallOption) (UserService_GetPostsClient, error)
	//client streaming
	CreatePost(ctx context.Context, opts ...grpc.CallOption) (UserService_CreatePostClient, error)
	//bidirectional streaming
	GreetEveryone(ctx context.Context, opts ...grpc.CallOption) (UserService_GreetEveryoneClient, error)
	//rpc for metadata
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	//rpc for jwt
	Jwt(ctx context.Context, in *JwtRequest, opts ...grpc.CallOption) (*JwtResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error) {
	out := new(SignupResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/Signup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetPosts(ctx context.Context, in *GetPostsRequest, opts ...grpc.CallOption) (UserService_GetPostsClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[0], "/proto.UserService/GetPosts", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGetPostsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserService_GetPostsClient interface {
	Recv() (*GetPostsResponse, error)
	grpc.ClientStream
}

type userServiceGetPostsClient struct {
	grpc.ClientStream
}

func (x *userServiceGetPostsClient) Recv() (*GetPostsResponse, error) {
	m := new(GetPostsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) CreatePost(ctx context.Context, opts ...grpc.CallOption) (UserService_CreatePostClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[1], "/proto.UserService/CreatePost", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceCreatePostClient{stream}
	return x, nil
}

type UserService_CreatePostClient interface {
	Send(*CreatePostRequest) error
	CloseAndRecv() (*CreatePostResponse, error)
	grpc.ClientStream
}

type userServiceCreatePostClient struct {
	grpc.ClientStream
}

func (x *userServiceCreatePostClient) Send(m *CreatePostRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userServiceCreatePostClient) CloseAndRecv() (*CreatePostResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CreatePostResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) GreetEveryone(ctx context.Context, opts ...grpc.CallOption) (UserService_GreetEveryoneClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[2], "/proto.UserService/GreetEveryone", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGreetEveryoneClient{stream}
	return x, nil
}

type UserService_GreetEveryoneClient interface {
	Send(*GreetEveryoneRequest) error
	Recv() (*GreetEveryoneResponse, error)
	grpc.ClientStream
}

type userServiceGreetEveryoneClient struct {
	grpc.ClientStream
}

func (x *userServiceGreetEveryoneClient) Send(m *GreetEveryoneRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userServiceGreetEveryoneClient) Recv() (*GreetEveryoneResponse, error) {
	m := new(GreetEveryoneResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Jwt(ctx context.Context, in *JwtRequest, opts ...grpc.CallOption) (*JwtResponse, error) {
	out := new(JwtResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/Jwt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	//rpc Signup(SignupRequest) returns (google.protobuf.Empty);
	//unary
	Signup(context.Context, *SignupRequest) (*SignupResponse, error)
	//server streaming
	GetPosts(*GetPostsRequest, UserService_GetPostsServer) error
	//client streaming
	CreatePost(UserService_CreatePostServer) error
	//bidirectional streaming
	GreetEveryone(UserService_GreetEveryoneServer) error
	//rpc for metadata
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	//rpc for jwt
	Jwt(context.Context, *JwtRequest) (*JwtResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Signup(context.Context, *SignupRequest) (*SignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (UnimplementedUserServiceServer) GetPosts(*GetPostsRequest, UserService_GetPostsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPosts not implemented")
}
func (UnimplementedUserServiceServer) CreatePost(UserService_CreatePostServer) error {
	return status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedUserServiceServer) GreetEveryone(UserService_GreetEveryoneServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetEveryone not implemented")
}
func (UnimplementedUserServiceServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedUserServiceServer) Jwt(context.Context, *JwtRequest) (*JwtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Jwt not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/Signup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Signup(ctx, req.(*SignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetPosts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetPostsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServiceServer).GetPosts(m, &userServiceGetPostsServer{stream})
}

type UserService_GetPostsServer interface {
	Send(*GetPostsResponse) error
	grpc.ServerStream
}

type userServiceGetPostsServer struct {
	grpc.ServerStream
}

func (x *userServiceGetPostsServer) Send(m *GetPostsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _UserService_CreatePost_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserServiceServer).CreatePost(&userServiceCreatePostServer{stream})
}

type UserService_CreatePostServer interface {
	SendAndClose(*CreatePostResponse) error
	Recv() (*CreatePostRequest, error)
	grpc.ServerStream
}

type userServiceCreatePostServer struct {
	grpc.ServerStream
}

func (x *userServiceCreatePostServer) SendAndClose(m *CreatePostResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userServiceCreatePostServer) Recv() (*CreatePostRequest, error) {
	m := new(CreatePostRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _UserService_GreetEveryone_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserServiceServer).GreetEveryone(&userServiceGreetEveryoneServer{stream})
}

type UserService_GreetEveryoneServer interface {
	Send(*GreetEveryoneResponse) error
	Recv() (*GreetEveryoneRequest, error)
	grpc.ServerStream
}

type userServiceGreetEveryoneServer struct {
	grpc.ServerStream
}

func (x *userServiceGreetEveryoneServer) Send(m *GreetEveryoneResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userServiceGreetEveryoneServer) Recv() (*GreetEveryoneRequest, error) {
	m := new(GreetEveryoneRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _UserService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Jwt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JwtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Jwt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/Jwt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Jwt(ctx, req.(*JwtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signup",
			Handler:    _UserService_Signup_Handler,
		},
		{
			MethodName: "Hello",
			Handler:    _UserService_Hello_Handler,
		},
		{
			MethodName: "Jwt",
			Handler:    _UserService_Jwt_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPosts",
			Handler:       _UserService_GetPosts_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CreatePost",
			Handler:       _UserService_CreatePost_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GreetEveryone",
			Handler:       _UserService_GreetEveryone_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/user.proto",
}
