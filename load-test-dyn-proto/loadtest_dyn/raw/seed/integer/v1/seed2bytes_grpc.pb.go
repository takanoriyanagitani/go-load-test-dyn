// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: loadtest_dyn/raw/seed/integer/v1/seed2bytes.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Seed64IService_BytesFrom64I_FullMethodName = "/loadtest_dyn.raw.seed.integer.v1.Seed64iService/BytesFrom64i"
)

// Seed64IServiceClient is the client API for Seed64IService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Seed64IServiceClient interface {
	BytesFrom64I(ctx context.Context, in *BytesFrom64IRequest, opts ...grpc.CallOption) (*BytesFrom64IResponse, error)
}

type seed64IServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSeed64IServiceClient(cc grpc.ClientConnInterface) Seed64IServiceClient {
	return &seed64IServiceClient{cc}
}

func (c *seed64IServiceClient) BytesFrom64I(ctx context.Context, in *BytesFrom64IRequest, opts ...grpc.CallOption) (*BytesFrom64IResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BytesFrom64IResponse)
	err := c.cc.Invoke(ctx, Seed64IService_BytesFrom64I_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Seed64IServiceServer is the server API for Seed64IService service.
// All implementations must embed UnimplementedSeed64IServiceServer
// for forward compatibility.
type Seed64IServiceServer interface {
	BytesFrom64I(context.Context, *BytesFrom64IRequest) (*BytesFrom64IResponse, error)
	mustEmbedUnimplementedSeed64IServiceServer()
}

// UnimplementedSeed64IServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSeed64IServiceServer struct{}

func (UnimplementedSeed64IServiceServer) BytesFrom64I(context.Context, *BytesFrom64IRequest) (*BytesFrom64IResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BytesFrom64I not implemented")
}
func (UnimplementedSeed64IServiceServer) mustEmbedUnimplementedSeed64IServiceServer() {}
func (UnimplementedSeed64IServiceServer) testEmbeddedByValue()                        {}

// UnsafeSeed64IServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Seed64IServiceServer will
// result in compilation errors.
type UnsafeSeed64IServiceServer interface {
	mustEmbedUnimplementedSeed64IServiceServer()
}

func RegisterSeed64IServiceServer(s grpc.ServiceRegistrar, srv Seed64IServiceServer) {
	// If the following call pancis, it indicates UnimplementedSeed64IServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Seed64IService_ServiceDesc, srv)
}

func _Seed64IService_BytesFrom64I_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BytesFrom64IRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Seed64IServiceServer).BytesFrom64I(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Seed64IService_BytesFrom64I_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Seed64IServiceServer).BytesFrom64I(ctx, req.(*BytesFrom64IRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Seed64IService_ServiceDesc is the grpc.ServiceDesc for Seed64IService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Seed64IService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "loadtest_dyn.raw.seed.integer.v1.Seed64iService",
	HandlerType: (*Seed64IServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BytesFrom64i",
			Handler:    _Seed64IService_BytesFrom64I_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "loadtest_dyn/raw/seed/integer/v1/seed2bytes.proto",
}
