// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: services.proto

package pb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UserService_CreateUser_FullMethodName   = "/services.UserService/CreateUser"
	UserService_ActivateUser_FullMethodName = "/services.UserService/ActivateUser"
	UserService_Login_FullMethodName        = "/services.UserService/Login"
	UserService_GetUser_FullMethodName      = "/services.UserService/GetUser"
	UserService_UpdateUser_FullMethodName   = "/services.UserService/UpdateUser"
	UserService_DeleteUser_FullMethodName   = "/services.UserService/DeleteUser"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// ----- SERVICE DEFINITIONS -----
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	ActivateUser(ctx context.Context, in *ActivateUserRequest, opts ...grpc.CallOption) (*ActivateUserResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, UserService_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ActivateUser(ctx context.Context, in *ActivateUserRequest, opts ...grpc.CallOption) (*ActivateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ActivateUserResponse)
	err := c.cc.Invoke(ctx, UserService_ActivateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, UserService_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, UserService_GetUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, UserService_UpdateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, UserService_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility.
//
// ----- SERVICE DEFINITIONS -----
type UserServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	ActivateUser(context.Context, *ActivateUserRequest) (*ActivateUserResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*empty.Empty, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) ActivateUser(context.Context, *ActivateUserRequest) (*ActivateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateUser not implemented")
}
func (UnimplementedUserServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}
func (UnimplementedUserServiceServer) testEmbeddedByValue()                     {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ActivateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ActivateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ActivateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ActivateUser(ctx, req.(*ActivateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "ActivateUser",
			Handler:    _UserService_ActivateUser_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}

const (
	FarmerService_GetFarmerInfo_FullMethodName = "/services.FarmerService/GetFarmerInfo"
)

// FarmerServiceClient is the client API for FarmerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FarmerServiceClient interface {
	GetFarmerInfo(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*FarmerInfo, error)
}

type farmerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFarmerServiceClient(cc grpc.ClientConnInterface) FarmerServiceClient {
	return &farmerServiceClient{cc}
}

func (c *farmerServiceClient) GetFarmerInfo(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*FarmerInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FarmerInfo)
	err := c.cc.Invoke(ctx, FarmerService_GetFarmerInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FarmerServiceServer is the server API for FarmerService service.
// All implementations must embed UnimplementedFarmerServiceServer
// for forward compatibility.
type FarmerServiceServer interface {
	GetFarmerInfo(context.Context, *GetUserRequest) (*FarmerInfo, error)
	mustEmbedUnimplementedFarmerServiceServer()
}

// UnimplementedFarmerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFarmerServiceServer struct{}

func (UnimplementedFarmerServiceServer) GetFarmerInfo(context.Context, *GetUserRequest) (*FarmerInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFarmerInfo not implemented")
}
func (UnimplementedFarmerServiceServer) mustEmbedUnimplementedFarmerServiceServer() {}
func (UnimplementedFarmerServiceServer) testEmbeddedByValue()                       {}

// UnsafeFarmerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FarmerServiceServer will
// result in compilation errors.
type UnsafeFarmerServiceServer interface {
	mustEmbedUnimplementedFarmerServiceServer()
}

func RegisterFarmerServiceServer(s grpc.ServiceRegistrar, srv FarmerServiceServer) {
	// If the following call pancis, it indicates UnimplementedFarmerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FarmerService_ServiceDesc, srv)
}

func _FarmerService_GetFarmerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FarmerServiceServer).GetFarmerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FarmerService_GetFarmerInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FarmerServiceServer).GetFarmerInfo(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FarmerService_ServiceDesc is the grpc.ServiceDesc for FarmerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FarmerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.FarmerService",
	HandlerType: (*FarmerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFarmerInfo",
			Handler:    _FarmerService_GetFarmerInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}

const (
	FarmService_CreateFarm_FullMethodName  = "/services.FarmService/CreateFarm"
	FarmService_GetFarmByID_FullMethodName = "/services.FarmService/GetFarmByID"
	FarmService_ListFarms_FullMethodName   = "/services.FarmService/ListFarms"
)

// FarmServiceClient is the client API for FarmService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FarmServiceClient interface {
	CreateFarm(ctx context.Context, in *Farm, opts ...grpc.CallOption) (*Farm, error)
	GetFarmByID(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*Farm, error)
	ListFarms(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Farm], error)
}

type farmServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFarmServiceClient(cc grpc.ClientConnInterface) FarmServiceClient {
	return &farmServiceClient{cc}
}

func (c *farmServiceClient) CreateFarm(ctx context.Context, in *Farm, opts ...grpc.CallOption) (*Farm, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Farm)
	err := c.cc.Invoke(ctx, FarmService_CreateFarm_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *farmServiceClient) GetFarmByID(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*Farm, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Farm)
	err := c.cc.Invoke(ctx, FarmService_GetFarmByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *farmServiceClient) ListFarms(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Farm], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &FarmService_ServiceDesc.Streams[0], FarmService_ListFarms_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[empty.Empty, Farm]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FarmService_ListFarmsClient = grpc.ServerStreamingClient[Farm]

// FarmServiceServer is the server API for FarmService service.
// All implementations must embed UnimplementedFarmServiceServer
// for forward compatibility.
type FarmServiceServer interface {
	CreateFarm(context.Context, *Farm) (*Farm, error)
	GetFarmByID(context.Context, *GetUserRequest) (*Farm, error)
	ListFarms(*empty.Empty, grpc.ServerStreamingServer[Farm]) error
	mustEmbedUnimplementedFarmServiceServer()
}

// UnimplementedFarmServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFarmServiceServer struct{}

func (UnimplementedFarmServiceServer) CreateFarm(context.Context, *Farm) (*Farm, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFarm not implemented")
}
func (UnimplementedFarmServiceServer) GetFarmByID(context.Context, *GetUserRequest) (*Farm, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFarmByID not implemented")
}
func (UnimplementedFarmServiceServer) ListFarms(*empty.Empty, grpc.ServerStreamingServer[Farm]) error {
	return status.Errorf(codes.Unimplemented, "method ListFarms not implemented")
}
func (UnimplementedFarmServiceServer) mustEmbedUnimplementedFarmServiceServer() {}
func (UnimplementedFarmServiceServer) testEmbeddedByValue()                     {}

// UnsafeFarmServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FarmServiceServer will
// result in compilation errors.
type UnsafeFarmServiceServer interface {
	mustEmbedUnimplementedFarmServiceServer()
}

func RegisterFarmServiceServer(s grpc.ServiceRegistrar, srv FarmServiceServer) {
	// If the following call pancis, it indicates UnimplementedFarmServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FarmService_ServiceDesc, srv)
}

func _FarmService_CreateFarm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Farm)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FarmServiceServer).CreateFarm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FarmService_CreateFarm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FarmServiceServer).CreateFarm(ctx, req.(*Farm))
	}
	return interceptor(ctx, in, info, handler)
}

func _FarmService_GetFarmByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FarmServiceServer).GetFarmByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FarmService_GetFarmByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FarmServiceServer).GetFarmByID(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FarmService_ListFarms_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FarmServiceServer).ListFarms(m, &grpc.GenericServerStream[empty.Empty, Farm]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FarmService_ListFarmsServer = grpc.ServerStreamingServer[Farm]

// FarmService_ServiceDesc is the grpc.ServiceDesc for FarmService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FarmService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.FarmService",
	HandlerType: (*FarmServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFarm",
			Handler:    _FarmService_CreateFarm_Handler,
		},
		{
			MethodName: "GetFarmByID",
			Handler:    _FarmService_GetFarmByID_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListFarms",
			Handler:       _FarmService_ListFarms_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "services.proto",
}

const (
	ApplicationService_CreateApplication_FullMethodName  = "/services.ApplicationService/CreateApplication"
	ApplicationService_GetApplicationByID_FullMethodName = "/services.ApplicationService/GetApplicationByID"
	ApplicationService_ListApplications_FullMethodName   = "/services.ApplicationService/ListApplications"
)

// ApplicationServiceClient is the client API for ApplicationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApplicationServiceClient interface {
	CreateApplication(ctx context.Context, in *Application, opts ...grpc.CallOption) (*empty.Empty, error)
	GetApplicationByID(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*Application, error)
	ListApplications(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Application], error)
}

type applicationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApplicationServiceClient(cc grpc.ClientConnInterface) ApplicationServiceClient {
	return &applicationServiceClient{cc}
}

func (c *applicationServiceClient) CreateApplication(ctx context.Context, in *Application, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, ApplicationService_CreateApplication_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) GetApplicationByID(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*Application, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Application)
	err := c.cc.Invoke(ctx, ApplicationService_GetApplicationByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) ListApplications(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Application], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ApplicationService_ServiceDesc.Streams[0], ApplicationService_ListApplications_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[empty.Empty, Application]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ApplicationService_ListApplicationsClient = grpc.ServerStreamingClient[Application]

// ApplicationServiceServer is the server API for ApplicationService service.
// All implementations must embed UnimplementedApplicationServiceServer
// for forward compatibility.
type ApplicationServiceServer interface {
	CreateApplication(context.Context, *Application) (*empty.Empty, error)
	GetApplicationByID(context.Context, *GetUserRequest) (*Application, error)
	ListApplications(*empty.Empty, grpc.ServerStreamingServer[Application]) error
	mustEmbedUnimplementedApplicationServiceServer()
}

// UnimplementedApplicationServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedApplicationServiceServer struct{}

func (UnimplementedApplicationServiceServer) CreateApplication(context.Context, *Application) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApplication not implemented")
}
func (UnimplementedApplicationServiceServer) GetApplicationByID(context.Context, *GetUserRequest) (*Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApplicationByID not implemented")
}
func (UnimplementedApplicationServiceServer) ListApplications(*empty.Empty, grpc.ServerStreamingServer[Application]) error {
	return status.Errorf(codes.Unimplemented, "method ListApplications not implemented")
}
func (UnimplementedApplicationServiceServer) mustEmbedUnimplementedApplicationServiceServer() {}
func (UnimplementedApplicationServiceServer) testEmbeddedByValue()                            {}

// UnsafeApplicationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApplicationServiceServer will
// result in compilation errors.
type UnsafeApplicationServiceServer interface {
	mustEmbedUnimplementedApplicationServiceServer()
}

func RegisterApplicationServiceServer(s grpc.ServiceRegistrar, srv ApplicationServiceServer) {
	// If the following call pancis, it indicates UnimplementedApplicationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ApplicationService_ServiceDesc, srv)
}

func _ApplicationService_CreateApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Application)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).CreateApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApplicationService_CreateApplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).CreateApplication(ctx, req.(*Application))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_GetApplicationByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).GetApplicationByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApplicationService_GetApplicationByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).GetApplicationByID(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_ListApplications_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ApplicationServiceServer).ListApplications(m, &grpc.GenericServerStream[empty.Empty, Application]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ApplicationService_ListApplicationsServer = grpc.ServerStreamingServer[Application]

// ApplicationService_ServiceDesc is the grpc.ServiceDesc for ApplicationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApplicationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.ApplicationService",
	HandlerType: (*ApplicationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateApplication",
			Handler:    _ApplicationService_CreateApplication_Handler,
		},
		{
			MethodName: "GetApplicationByID",
			Handler:    _ApplicationService_GetApplicationByID_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListApplications",
			Handler:       _ApplicationService_ListApplications_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "services.proto",
}