// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package shapecalc

import (
	context "context"
	shape "github.com/muzammilar/geometric-shapes/protos/shape"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GeometryClient is the client API for Geometry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeometryClient interface {
	// A simple RPC.
	//
	// Obtains the feature at a given position.
	//
	// A mesurement with an empty name is returned if there's no field at the given
	// position.
	ComputeArea(ctx context.Context, in *shape.Rectangle, opts ...grpc.CallOption) (*shape.ShapeInfo_Mesurement, error)
	GetRectangleInfo(ctx context.Context, in *shape.Rectangle, opts ...grpc.CallOption) (*shape.ShapeInfo, error)
	// A server-to-client streaming RPC.
	//
	// Obtains all the Planar Coordinate (dimensions) available within the given Rectangle.  Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a
	// huge number of features.
	ListRectangleCoordinates(ctx context.Context, in *shape.Rectangle, opts ...grpc.CallOption) (Geometry_ListRectangleCoordinatesClient, error)
}

type geometryClient struct {
	cc grpc.ClientConnInterface
}

func NewGeometryClient(cc grpc.ClientConnInterface) GeometryClient {
	return &geometryClient{cc}
}

func (c *geometryClient) ComputeArea(ctx context.Context, in *shape.Rectangle, opts ...grpc.CallOption) (*shape.ShapeInfo_Mesurement, error) {
	out := new(shape.ShapeInfo_Mesurement)
	err := c.cc.Invoke(ctx, "/shapecalc.Geometry/ComputeArea", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geometryClient) GetRectangleInfo(ctx context.Context, in *shape.Rectangle, opts ...grpc.CallOption) (*shape.ShapeInfo, error) {
	out := new(shape.ShapeInfo)
	err := c.cc.Invoke(ctx, "/shapecalc.Geometry/GetRectangleInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geometryClient) ListRectangleCoordinates(ctx context.Context, in *shape.Rectangle, opts ...grpc.CallOption) (Geometry_ListRectangleCoordinatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Geometry_ServiceDesc.Streams[0], "/shapecalc.Geometry/ListRectangleCoordinates", opts...)
	if err != nil {
		return nil, err
	}
	x := &geometryListRectangleCoordinatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Geometry_ListRectangleCoordinatesClient interface {
	Recv() (*shape.PlanarCoordinates, error)
	grpc.ClientStream
}

type geometryListRectangleCoordinatesClient struct {
	grpc.ClientStream
}

func (x *geometryListRectangleCoordinatesClient) Recv() (*shape.PlanarCoordinates, error) {
	m := new(shape.PlanarCoordinates)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GeometryServer is the server API for Geometry service.
// All implementations must embed UnimplementedGeometryServer
// for forward compatibility
type GeometryServer interface {
	// A simple RPC.
	//
	// Obtains the feature at a given position.
	//
	// A mesurement with an empty name is returned if there's no field at the given
	// position.
	ComputeArea(context.Context, *shape.Rectangle) (*shape.ShapeInfo_Mesurement, error)
	GetRectangleInfo(context.Context, *shape.Rectangle) (*shape.ShapeInfo, error)
	// A server-to-client streaming RPC.
	//
	// Obtains all the Planar Coordinate (dimensions) available within the given Rectangle.  Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a
	// huge number of features.
	ListRectangleCoordinates(*shape.Rectangle, Geometry_ListRectangleCoordinatesServer) error
	mustEmbedUnimplementedGeometryServer()
}

// UnimplementedGeometryServer must be embedded to have forward compatible implementations.
type UnimplementedGeometryServer struct {
}

func (UnimplementedGeometryServer) ComputeArea(context.Context, *shape.Rectangle) (*shape.ShapeInfo_Mesurement, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComputeArea not implemented")
}
func (UnimplementedGeometryServer) GetRectangleInfo(context.Context, *shape.Rectangle) (*shape.ShapeInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRectangleInfo not implemented")
}
func (UnimplementedGeometryServer) ListRectangleCoordinates(*shape.Rectangle, Geometry_ListRectangleCoordinatesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListRectangleCoordinates not implemented")
}
func (UnimplementedGeometryServer) mustEmbedUnimplementedGeometryServer() {}

// UnsafeGeometryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeometryServer will
// result in compilation errors.
type UnsafeGeometryServer interface {
	mustEmbedUnimplementedGeometryServer()
}

func RegisterGeometryServer(s grpc.ServiceRegistrar, srv GeometryServer) {
	s.RegisterService(&Geometry_ServiceDesc, srv)
}

func _Geometry_ComputeArea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(shape.Rectangle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeometryServer).ComputeArea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shapecalc.Geometry/ComputeArea",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeometryServer).ComputeArea(ctx, req.(*shape.Rectangle))
	}
	return interceptor(ctx, in, info, handler)
}

func _Geometry_GetRectangleInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(shape.Rectangle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeometryServer).GetRectangleInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shapecalc.Geometry/GetRectangleInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeometryServer).GetRectangleInfo(ctx, req.(*shape.Rectangle))
	}
	return interceptor(ctx, in, info, handler)
}

func _Geometry_ListRectangleCoordinates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(shape.Rectangle)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GeometryServer).ListRectangleCoordinates(m, &geometryListRectangleCoordinatesServer{stream})
}

type Geometry_ListRectangleCoordinatesServer interface {
	Send(*shape.PlanarCoordinates) error
	grpc.ServerStream
}

type geometryListRectangleCoordinatesServer struct {
	grpc.ServerStream
}

func (x *geometryListRectangleCoordinatesServer) Send(m *shape.PlanarCoordinates) error {
	return x.ServerStream.SendMsg(m)
}

// Geometry_ServiceDesc is the grpc.ServiceDesc for Geometry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Geometry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shapecalc.Geometry",
	HandlerType: (*GeometryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ComputeArea",
			Handler:    _Geometry_ComputeArea_Handler,
		},
		{
			MethodName: "GetRectangleInfo",
			Handler:    _Geometry_GetRectangleInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListRectangleCoordinates",
			Handler:       _Geometry_ListRectangleCoordinates_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "shapecalc.proto",
}