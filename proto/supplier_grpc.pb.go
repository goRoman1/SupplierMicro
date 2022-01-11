// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// SupplierServiceClient is the client API for SupplierService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SupplierServiceClient interface {
	CreateStation(ctx context.Context, in *ScooterStation, opts ...grpc.CallOption) (*Response, error)
	GetLocations(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Location, error)
	CreateStationInLocation(ctx context.Context, in *StationLocation, opts ...grpc.CallOption) (*Response, error)
}

type supplierServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSupplierServiceClient(cc grpc.ClientConnInterface) SupplierServiceClient {
	return &supplierServiceClient{cc}
}

func (c *supplierServiceClient) CreateStation(ctx context.Context, in *ScooterStation, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/supplier.SupplierService/CreateStation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierServiceClient) GetLocations(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Location, error) {
	out := new(Location)
	err := c.cc.Invoke(ctx, "/supplier.SupplierService/GetLocations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierServiceClient) CreateStationInLocation(ctx context.Context, in *StationLocation, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/supplier.SupplierService/CreateStationInLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SupplierServiceServer is the server API for SupplierService service.
// All implementations must embed UnimplementedSupplierServiceServer
// for forward compatibility
type SupplierServiceServer interface {
	CreateStation(context.Context, *ScooterStation) (*Response, error)
	GetLocations(context.Context, *Request) (*Location, error)
	CreateStationInLocation(context.Context, *StationLocation) (*Response, error)
	mustEmbedUnimplementedSupplierServiceServer()
}

// UnimplementedSupplierServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSupplierServiceServer struct {
}

func (UnimplementedSupplierServiceServer) CreateStation(context.Context, *ScooterStation) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStation not implemented")
}
func (UnimplementedSupplierServiceServer) GetLocations(context.Context, *Request) (*Location, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocations not implemented")
}
func (UnimplementedSupplierServiceServer) CreateStationInLocation(context.Context, *StationLocation) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStationInLocation not implemented")
}
func (UnimplementedSupplierServiceServer) mustEmbedUnimplementedSupplierServiceServer() {}

// UnsafeSupplierServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SupplierServiceServer will
// result in compilation errors.
type UnsafeSupplierServiceServer interface {
	mustEmbedUnimplementedSupplierServiceServer()
}

func RegisterSupplierServiceServer(s grpc.ServiceRegistrar, srv SupplierServiceServer) {
	s.RegisterService(&SupplierService_ServiceDesc, srv)
}

func _SupplierService_CreateStation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScooterStation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).CreateStation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supplier.SupplierService/CreateStation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).CreateStation(ctx, req.(*ScooterStation))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierService_GetLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).GetLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supplier.SupplierService/GetLocations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).GetLocations(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierService_CreateStationInLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StationLocation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).CreateStationInLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supplier.SupplierService/CreateStationInLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).CreateStationInLocation(ctx, req.(*StationLocation))
	}
	return interceptor(ctx, in, info, handler)
}

// SupplierService_ServiceDesc is the grpc.ServiceDesc for SupplierService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SupplierService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "supplier.SupplierService",
	HandlerType: (*SupplierServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStation",
			Handler:    _SupplierService_CreateStation_Handler,
		},
		{
			MethodName: "GetLocations",
			Handler:    _SupplierService_GetLocations_Handler,
		},
		{
			MethodName: "CreateStationInLocation",
			Handler:    _SupplierService_CreateStationInLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "supplier.proto",
}
