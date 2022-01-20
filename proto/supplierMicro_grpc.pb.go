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

// SupplierMicroServiceClient is the client API for SupplierMicroService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SupplierMicroServiceClient interface {
	CreateStation(ctx context.Context, in *Station, opts ...grpc.CallOption) (*Response, error)
	GetLocations(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	GetStations(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	CreateStationInLocation(ctx context.Context, in *StationLocation, opts ...grpc.CallOption) (*Response, error)
}

type supplierMicroServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSupplierMicroServiceClient(cc grpc.ClientConnInterface) SupplierMicroServiceClient {
	return &supplierMicroServiceClient{cc}
}

func (c *supplierMicroServiceClient) CreateStation(ctx context.Context, in *Station, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/supplierMicro.SupplierMicroService/CreateStation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierMicroServiceClient) GetLocations(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/supplierMicro.SupplierMicroService/GetLocations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierMicroServiceClient) GetStations(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/supplierMicro.SupplierMicroService/GetStations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierMicroServiceClient) CreateStationInLocation(ctx context.Context, in *StationLocation, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/supplierMicro.SupplierMicroService/CreateStationInLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SupplierMicroServiceServer is the server API for SupplierMicroService service.
// All implementations must embed UnimplementedSupplierMicroServiceServer
// for forward compatibility
type SupplierMicroServiceServer interface {
	CreateStation(context.Context, *Station) (*Response, error)
	GetLocations(context.Context, *Request) (*Response, error)
	GetStations(context.Context, *Request) (*Response, error)
	CreateStationInLocation(context.Context, *StationLocation) (*Response, error)
	mustEmbedUnimplementedSupplierMicroServiceServer()
}

// UnimplementedSupplierMicroServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSupplierMicroServiceServer struct {
}

func (UnimplementedSupplierMicroServiceServer) CreateStation(context.Context, *Station) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStation not implemented")
}
func (UnimplementedSupplierMicroServiceServer) GetLocations(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocations not implemented")
}
func (UnimplementedSupplierMicroServiceServer) GetStations(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStations not implemented")
}
func (UnimplementedSupplierMicroServiceServer) CreateStationInLocation(context.Context, *StationLocation) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStationInLocation not implemented")
}
func (UnimplementedSupplierMicroServiceServer) mustEmbedUnimplementedSupplierMicroServiceServer() {}

// UnsafeSupplierMicroServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SupplierMicroServiceServer will
// result in compilation errors.
type UnsafeSupplierMicroServiceServer interface {
	mustEmbedUnimplementedSupplierMicroServiceServer()
}

func RegisterSupplierMicroServiceServer(s grpc.ServiceRegistrar, srv SupplierMicroServiceServer) {
	s.RegisterService(&SupplierMicroService_ServiceDesc, srv)
}

func _SupplierMicroService_CreateStation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Station)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierMicroServiceServer).CreateStation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supplierMicro.SupplierMicroService/CreateStation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierMicroServiceServer).CreateStation(ctx, req.(*Station))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierMicroService_GetLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierMicroServiceServer).GetLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supplierMicro.SupplierMicroService/GetLocations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierMicroServiceServer).GetLocations(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierMicroService_GetStations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierMicroServiceServer).GetStations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supplierMicro.SupplierMicroService/GetStations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierMicroServiceServer).GetStations(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierMicroService_CreateStationInLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StationLocation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierMicroServiceServer).CreateStationInLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supplierMicro.SupplierMicroService/CreateStationInLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierMicroServiceServer).CreateStationInLocation(ctx, req.(*StationLocation))
	}
	return interceptor(ctx, in, info, handler)
}

// SupplierMicroService_ServiceDesc is the grpc.ServiceDesc for SupplierMicroService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SupplierMicroService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "supplierMicro.SupplierMicroService",
	HandlerType: (*SupplierMicroServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStation",
			Handler:    _SupplierMicroService_CreateStation_Handler,
		},
		{
			MethodName: "GetLocations",
			Handler:    _SupplierMicroService_GetLocations_Handler,
		},
		{
			MethodName: "GetStations",
			Handler:    _SupplierMicroService_GetStations_Handler,
		},
		{
			MethodName: "CreateStationInLocation",
			Handler:    _SupplierMicroService_CreateStationInLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "supplierMicro.proto",
}
