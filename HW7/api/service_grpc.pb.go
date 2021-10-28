// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserList, error)
	Get(ctx context.Context, in *UserRequestId, opts ...grpc.CallOption) (*User, error)
	Insert(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	Update(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	Remove(ctx context.Context, in *UserRequestId, opts ...grpc.CallOption) (*Empty, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserList, error) {
	out := new(UserList)
	err := c.cc.Invoke(ctx, "/api.UserService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Get(ctx context.Context, in *UserRequestId, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/api.UserService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Insert(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/api.UserService/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/api.UserService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Remove(ctx context.Context, in *UserRequestId, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/api.UserService/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	GetAll(context.Context, *Empty) (*UserList, error)
	Get(context.Context, *UserRequestId) (*User, error)
	Insert(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*User, error)
	Remove(context.Context, *UserRequestId) (*Empty, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetAll(context.Context, *Empty) (*UserList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedUserServiceServer) Get(context.Context, *UserRequestId) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedUserServiceServer) Insert(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedUserServiceServer) Update(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUserServiceServer) Remove(context.Context, *UserRequestId) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
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

func _UserService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAll(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Get(ctx, req.(*UserRequestId))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Insert(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Update(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Remove(ctx, req.(*UserRequestId))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _UserService_GetAll_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UserService_Get_Handler,
		},
		{
			MethodName: "Insert",
			Handler:    _UserService_Insert_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserService_Update_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _UserService_Remove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/service.proto",
}

// SubmissionServiceClient is the client API for SubmissionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubmissionServiceClient interface {
	GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SubmissionList, error)
	Get(ctx context.Context, in *SubmissionRequestId, opts ...grpc.CallOption) (*Submission, error)
	Insert(ctx context.Context, in *Submission, opts ...grpc.CallOption) (*Submission, error)
	Update(ctx context.Context, in *Submission, opts ...grpc.CallOption) (*Submission, error)
	Remove(ctx context.Context, in *SubmissionRequestId, opts ...grpc.CallOption) (*Empty, error)
}

type submissionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubmissionServiceClient(cc grpc.ClientConnInterface) SubmissionServiceClient {
	return &submissionServiceClient{cc}
}

func (c *submissionServiceClient) GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SubmissionList, error) {
	out := new(SubmissionList)
	err := c.cc.Invoke(ctx, "/api.SubmissionService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *submissionServiceClient) Get(ctx context.Context, in *SubmissionRequestId, opts ...grpc.CallOption) (*Submission, error) {
	out := new(Submission)
	err := c.cc.Invoke(ctx, "/api.SubmissionService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *submissionServiceClient) Insert(ctx context.Context, in *Submission, opts ...grpc.CallOption) (*Submission, error) {
	out := new(Submission)
	err := c.cc.Invoke(ctx, "/api.SubmissionService/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *submissionServiceClient) Update(ctx context.Context, in *Submission, opts ...grpc.CallOption) (*Submission, error) {
	out := new(Submission)
	err := c.cc.Invoke(ctx, "/api.SubmissionService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *submissionServiceClient) Remove(ctx context.Context, in *SubmissionRequestId, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/api.SubmissionService/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubmissionServiceServer is the server API for SubmissionService service.
// All implementations must embed UnimplementedSubmissionServiceServer
// for forward compatibility
type SubmissionServiceServer interface {
	GetAll(context.Context, *Empty) (*SubmissionList, error)
	Get(context.Context, *SubmissionRequestId) (*Submission, error)
	Insert(context.Context, *Submission) (*Submission, error)
	Update(context.Context, *Submission) (*Submission, error)
	Remove(context.Context, *SubmissionRequestId) (*Empty, error)
	mustEmbedUnimplementedSubmissionServiceServer()
}

// UnimplementedSubmissionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSubmissionServiceServer struct {
}

func (UnimplementedSubmissionServiceServer) GetAll(context.Context, *Empty) (*SubmissionList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedSubmissionServiceServer) Get(context.Context, *SubmissionRequestId) (*Submission, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSubmissionServiceServer) Insert(context.Context, *Submission) (*Submission, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedSubmissionServiceServer) Update(context.Context, *Submission) (*Submission, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSubmissionServiceServer) Remove(context.Context, *SubmissionRequestId) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedSubmissionServiceServer) mustEmbedUnimplementedSubmissionServiceServer() {}

// UnsafeSubmissionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubmissionServiceServer will
// result in compilation errors.
type UnsafeSubmissionServiceServer interface {
	mustEmbedUnimplementedSubmissionServiceServer()
}

func RegisterSubmissionServiceServer(s grpc.ServiceRegistrar, srv SubmissionServiceServer) {
	s.RegisterService(&SubmissionService_ServiceDesc, srv)
}

func _SubmissionService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmissionServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SubmissionService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmissionServiceServer).GetAll(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubmissionService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmissionRequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmissionServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SubmissionService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmissionServiceServer).Get(ctx, req.(*SubmissionRequestId))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubmissionService_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Submission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmissionServiceServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SubmissionService/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmissionServiceServer).Insert(ctx, req.(*Submission))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubmissionService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Submission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmissionServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SubmissionService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmissionServiceServer).Update(ctx, req.(*Submission))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubmissionService_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmissionRequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmissionServiceServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SubmissionService/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmissionServiceServer).Remove(ctx, req.(*SubmissionRequestId))
	}
	return interceptor(ctx, in, info, handler)
}

// SubmissionService_ServiceDesc is the grpc.ServiceDesc for SubmissionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubmissionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.SubmissionService",
	HandlerType: (*SubmissionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _SubmissionService_GetAll_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _SubmissionService_Get_Handler,
		},
		{
			MethodName: "Insert",
			Handler:    _SubmissionService_Insert_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SubmissionService_Update_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _SubmissionService_Remove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/service.proto",
}
