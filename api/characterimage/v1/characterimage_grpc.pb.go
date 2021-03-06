// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: characterimage/v1/characterimage.proto

package characterimage

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

// ImagesClient is the client API for Images service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImagesClient interface {
	// AddBody adds a body for use with an animation. This is the basic component of any animation.
	AddBody(ctx context.Context, in *AddBodyRequest, opts ...grpc.CallOption) (*AddBodyResponse, error)
	// ListBodies gets a list of all of the bodies available.
	// TODO: Add filter.
	ListBodies(ctx context.Context, in *ListBodiesRequest, opts ...grpc.CallOption) (*ListBodiesResponse, error)
	// AddDynamicMapping adds a new dynamic element mapping for use with an animation. This image
	// should be comprised of a series of pixels that are each uniquely colored. These pixels are used
	// to map onto an animation frame, and their positional data correlates with dynamics images in
	// the same shape.
	AddDynamicMapping(ctx context.Context, in *AddDynamicMappingRequest, opts ...grpc.CallOption) (*AddDynamicMappingResponse, error)
	// AddStatic creates a new static part for use with an animation. Note that these parts only
	// change in that they either move laterally or rotate.
	AddStatic(ctx context.Context, in *AddStaticRequest, opts ...grpc.CallOption) (*AddStaticResponse, error)
	// ListStatics returns a list of all of the statics available.
	// TODO: Add filter.
	ListStatics(ctx context.Context, in *ListStaticRequest, opts ...grpc.CallOption) (*ListStaticsResponse, error)
	// AddDynamic adds a new dynamic for use with an animation. This dynamic should be comprised of a
	// set of pixels where the position of each pixel correlates to the position of a pixel on the
	// related mapping type for the associated body type.
	AddDynamic(ctx context.Context, in *AddDynamicRequest, opts ...grpc.CallOption) (*AddDynamicResponse, error)
	// ListDynamics returns a list of all of the dynamics available.
	ListDynamics(ctx context.Context, in *ListDynamicsRequest, opts ...grpc.CallOption) (*ListDynamicsResponse, error)
	// AddAnimation adds a new animation. Note that any pose is also an animation (just one with a
	// single frame).
	AddAnimation(ctx context.Context, in *AddAnimationRequest, opts ...grpc.CallOption) (*AddAnimationResponse, error)
	// ListAnimations returns all of the animations available.
	// TODO: Add filter.
	ListAnimations(ctx context.Context, in *ListAnimationsRequest, opts ...grpc.CallOption) (*ListAnimationsResponse, error)
	// AddFrame adds a new frame for an animation.
	AddFrame(ctx context.Context, in *AddFrameRequest, opts ...grpc.CallOption) (*AddFrameResponse, error)
	// AddProp adds a new prop.
	AddProp(ctx context.Context, in *AddPropRequest, opts ...grpc.CallOption) (*AddPropResponse, error)
	// ListProps returns a list of all available props.
	// TODO: Add filter.
	ListProps(ctx context.Context, in *ListPropsRequest, opts ...grpc.CallOption) (*ListPropsResponse, error)
	// GenerateAnimation accepts a series of parameters to return all of the data necessary to render
	// a complete animation. This is the "selection" phase after all of the components have been added
	// to the system.
	GenerateAnimation(ctx context.Context, in *GenerateAnimationRequest, opts ...grpc.CallOption) (*GenerateAnimationResponse, error)
}

type imagesClient struct {
	cc grpc.ClientConnInterface
}

func NewImagesClient(cc grpc.ClientConnInterface) ImagesClient {
	return &imagesClient{cc}
}

func (c *imagesClient) AddBody(ctx context.Context, in *AddBodyRequest, opts ...grpc.CallOption) (*AddBodyResponse, error) {
	out := new(AddBodyResponse)
	err := c.cc.Invoke(ctx, "/lantspants.lootloadout.characterimage.Images/AddBody", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) ListBodies(ctx context.Context, in *ListBodiesRequest, opts ...grpc.CallOption) (*ListBodiesResponse, error) {
	out := new(ListBodiesResponse)
	err := c.cc.Invoke(ctx, "/lantspants.lootloadout.characterimage.Images/ListBodies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) AddDynamicMapping(ctx context.Context, in *AddDynamicMappingRequest, opts ...grpc.CallOption) (*AddDynamicMappingResponse, error) {
	out := new(AddDynamicMappingResponse)
	err := c.cc.Invoke(ctx, "/lantspants.lootloadout.characterimage.Images/AddDynamicMapping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) AddStatic(ctx context.Context, in *AddStaticRequest, opts ...grpc.CallOption) (*AddStaticResponse, error) {
	out := new(AddStaticResponse)
	err := c.cc.Invoke(ctx, "/lantspants.lootloadout.characterimage.Images/AddStatic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) ListStatics(ctx context.Context, in *ListStaticRequest, opts ...grpc.CallOption) (*ListStaticsResponse, error) {
	out := new(ListStaticsResponse)
	err := c.cc.Invoke(ctx, "/lantspants.lootloadout.characterimage.Images/ListStatics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) AddDynamic(ctx context.Context, in *AddDynamicRequest, opts ...grpc.CallOption) (*AddDynamicResponse, error) {
	out := new(AddDynamicResponse)
	err := c.cc.Invoke(ctx, "/lantspants.lootloadout.characterimage.Images/AddDynamic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) ListDynamics(ctx context.Context, in *ListDynamicsRequest, opts ...grpc.CallOption) (*ListDynamicsResponse, error) {
	out := new(ListDynamicsResponse)
	err := c.cc.Invoke(ctx, "/lantspants.lootloadout.characterimage.Images/ListDynamics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) AddAnimation(ctx context.Context, in *AddAnimationRequest, opts ...grpc.CallOption) (*AddAnimationResponse, error) {
	out := new(AddAnimationResponse)
	err := c.cc.Invoke(ctx, "/lantspants.lootloadout.characterimage.Images/AddAnimation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) ListAnimations(ctx context.Context, in *ListAnimationsRequest, opts ...grpc.CallOption) (*ListAnimationsResponse, error) {
	out := new(ListAnimationsResponse)
	err := c.cc.Invoke(ctx, "/lantspants.lootloadout.characterimage.Images/ListAnimations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) AddFrame(ctx context.Context, in *AddFrameRequest, opts ...grpc.CallOption) (*AddFrameResponse, error) {
	out := new(AddFrameResponse)
	err := c.cc.Invoke(ctx, "/lantspants.lootloadout.characterimage.Images/AddFrame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) AddProp(ctx context.Context, in *AddPropRequest, opts ...grpc.CallOption) (*AddPropResponse, error) {
	out := new(AddPropResponse)
	err := c.cc.Invoke(ctx, "/lantspants.lootloadout.characterimage.Images/AddProp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) ListProps(ctx context.Context, in *ListPropsRequest, opts ...grpc.CallOption) (*ListPropsResponse, error) {
	out := new(ListPropsResponse)
	err := c.cc.Invoke(ctx, "/lantspants.lootloadout.characterimage.Images/ListProps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) GenerateAnimation(ctx context.Context, in *GenerateAnimationRequest, opts ...grpc.CallOption) (*GenerateAnimationResponse, error) {
	out := new(GenerateAnimationResponse)
	err := c.cc.Invoke(ctx, "/lantspants.lootloadout.characterimage.Images/GenerateAnimation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImagesServer is the server API for Images service.
// All implementations must embed UnimplementedImagesServer
// for forward compatibility
type ImagesServer interface {
	// AddBody adds a body for use with an animation. This is the basic component of any animation.
	AddBody(context.Context, *AddBodyRequest) (*AddBodyResponse, error)
	// ListBodies gets a list of all of the bodies available.
	// TODO: Add filter.
	ListBodies(context.Context, *ListBodiesRequest) (*ListBodiesResponse, error)
	// AddDynamicMapping adds a new dynamic element mapping for use with an animation. This image
	// should be comprised of a series of pixels that are each uniquely colored. These pixels are used
	// to map onto an animation frame, and their positional data correlates with dynamics images in
	// the same shape.
	AddDynamicMapping(context.Context, *AddDynamicMappingRequest) (*AddDynamicMappingResponse, error)
	// AddStatic creates a new static part for use with an animation. Note that these parts only
	// change in that they either move laterally or rotate.
	AddStatic(context.Context, *AddStaticRequest) (*AddStaticResponse, error)
	// ListStatics returns a list of all of the statics available.
	// TODO: Add filter.
	ListStatics(context.Context, *ListStaticRequest) (*ListStaticsResponse, error)
	// AddDynamic adds a new dynamic for use with an animation. This dynamic should be comprised of a
	// set of pixels where the position of each pixel correlates to the position of a pixel on the
	// related mapping type for the associated body type.
	AddDynamic(context.Context, *AddDynamicRequest) (*AddDynamicResponse, error)
	// ListDynamics returns a list of all of the dynamics available.
	ListDynamics(context.Context, *ListDynamicsRequest) (*ListDynamicsResponse, error)
	// AddAnimation adds a new animation. Note that any pose is also an animation (just one with a
	// single frame).
	AddAnimation(context.Context, *AddAnimationRequest) (*AddAnimationResponse, error)
	// ListAnimations returns all of the animations available.
	// TODO: Add filter.
	ListAnimations(context.Context, *ListAnimationsRequest) (*ListAnimationsResponse, error)
	// AddFrame adds a new frame for an animation.
	AddFrame(context.Context, *AddFrameRequest) (*AddFrameResponse, error)
	// AddProp adds a new prop.
	AddProp(context.Context, *AddPropRequest) (*AddPropResponse, error)
	// ListProps returns a list of all available props.
	// TODO: Add filter.
	ListProps(context.Context, *ListPropsRequest) (*ListPropsResponse, error)
	// GenerateAnimation accepts a series of parameters to return all of the data necessary to render
	// a complete animation. This is the "selection" phase after all of the components have been added
	// to the system.
	GenerateAnimation(context.Context, *GenerateAnimationRequest) (*GenerateAnimationResponse, error)
	mustEmbedUnimplementedImagesServer()
}

// UnimplementedImagesServer must be embedded to have forward compatible implementations.
type UnimplementedImagesServer struct {
}

func (UnimplementedImagesServer) AddBody(context.Context, *AddBodyRequest) (*AddBodyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBody not implemented")
}
func (UnimplementedImagesServer) ListBodies(context.Context, *ListBodiesRequest) (*ListBodiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBodies not implemented")
}
func (UnimplementedImagesServer) AddDynamicMapping(context.Context, *AddDynamicMappingRequest) (*AddDynamicMappingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDynamicMapping not implemented")
}
func (UnimplementedImagesServer) AddStatic(context.Context, *AddStaticRequest) (*AddStaticResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStatic not implemented")
}
func (UnimplementedImagesServer) ListStatics(context.Context, *ListStaticRequest) (*ListStaticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStatics not implemented")
}
func (UnimplementedImagesServer) AddDynamic(context.Context, *AddDynamicRequest) (*AddDynamicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDynamic not implemented")
}
func (UnimplementedImagesServer) ListDynamics(context.Context, *ListDynamicsRequest) (*ListDynamicsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDynamics not implemented")
}
func (UnimplementedImagesServer) AddAnimation(context.Context, *AddAnimationRequest) (*AddAnimationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAnimation not implemented")
}
func (UnimplementedImagesServer) ListAnimations(context.Context, *ListAnimationsRequest) (*ListAnimationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAnimations not implemented")
}
func (UnimplementedImagesServer) AddFrame(context.Context, *AddFrameRequest) (*AddFrameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFrame not implemented")
}
func (UnimplementedImagesServer) AddProp(context.Context, *AddPropRequest) (*AddPropResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProp not implemented")
}
func (UnimplementedImagesServer) ListProps(context.Context, *ListPropsRequest) (*ListPropsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProps not implemented")
}
func (UnimplementedImagesServer) GenerateAnimation(context.Context, *GenerateAnimationRequest) (*GenerateAnimationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateAnimation not implemented")
}
func (UnimplementedImagesServer) mustEmbedUnimplementedImagesServer() {}

// UnsafeImagesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImagesServer will
// result in compilation errors.
type UnsafeImagesServer interface {
	mustEmbedUnimplementedImagesServer()
}

func RegisterImagesServer(s grpc.ServiceRegistrar, srv ImagesServer) {
	s.RegisterService(&Images_ServiceDesc, srv)
}

func _Images_AddBody_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBodyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).AddBody(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lantspants.lootloadout.characterimage.Images/AddBody",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).AddBody(ctx, req.(*AddBodyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_ListBodies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBodiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).ListBodies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lantspants.lootloadout.characterimage.Images/ListBodies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).ListBodies(ctx, req.(*ListBodiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_AddDynamicMapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDynamicMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).AddDynamicMapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lantspants.lootloadout.characterimage.Images/AddDynamicMapping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).AddDynamicMapping(ctx, req.(*AddDynamicMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_AddStatic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddStaticRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).AddStatic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lantspants.lootloadout.characterimage.Images/AddStatic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).AddStatic(ctx, req.(*AddStaticRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_ListStatics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStaticRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).ListStatics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lantspants.lootloadout.characterimage.Images/ListStatics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).ListStatics(ctx, req.(*ListStaticRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_AddDynamic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDynamicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).AddDynamic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lantspants.lootloadout.characterimage.Images/AddDynamic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).AddDynamic(ctx, req.(*AddDynamicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_ListDynamics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDynamicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).ListDynamics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lantspants.lootloadout.characterimage.Images/ListDynamics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).ListDynamics(ctx, req.(*ListDynamicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_AddAnimation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAnimationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).AddAnimation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lantspants.lootloadout.characterimage.Images/AddAnimation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).AddAnimation(ctx, req.(*AddAnimationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_ListAnimations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAnimationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).ListAnimations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lantspants.lootloadout.characterimage.Images/ListAnimations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).ListAnimations(ctx, req.(*ListAnimationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_AddFrame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFrameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).AddFrame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lantspants.lootloadout.characterimage.Images/AddFrame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).AddFrame(ctx, req.(*AddFrameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_AddProp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPropRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).AddProp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lantspants.lootloadout.characterimage.Images/AddProp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).AddProp(ctx, req.(*AddPropRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_ListProps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPropsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).ListProps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lantspants.lootloadout.characterimage.Images/ListProps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).ListProps(ctx, req.(*ListPropsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_GenerateAnimation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateAnimationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).GenerateAnimation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lantspants.lootloadout.characterimage.Images/GenerateAnimation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).GenerateAnimation(ctx, req.(*GenerateAnimationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Images_ServiceDesc is the grpc.ServiceDesc for Images service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Images_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lantspants.lootloadout.characterimage.Images",
	HandlerType: (*ImagesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBody",
			Handler:    _Images_AddBody_Handler,
		},
		{
			MethodName: "ListBodies",
			Handler:    _Images_ListBodies_Handler,
		},
		{
			MethodName: "AddDynamicMapping",
			Handler:    _Images_AddDynamicMapping_Handler,
		},
		{
			MethodName: "AddStatic",
			Handler:    _Images_AddStatic_Handler,
		},
		{
			MethodName: "ListStatics",
			Handler:    _Images_ListStatics_Handler,
		},
		{
			MethodName: "AddDynamic",
			Handler:    _Images_AddDynamic_Handler,
		},
		{
			MethodName: "ListDynamics",
			Handler:    _Images_ListDynamics_Handler,
		},
		{
			MethodName: "AddAnimation",
			Handler:    _Images_AddAnimation_Handler,
		},
		{
			MethodName: "ListAnimations",
			Handler:    _Images_ListAnimations_Handler,
		},
		{
			MethodName: "AddFrame",
			Handler:    _Images_AddFrame_Handler,
		},
		{
			MethodName: "AddProp",
			Handler:    _Images_AddProp_Handler,
		},
		{
			MethodName: "ListProps",
			Handler:    _Images_ListProps_Handler,
		},
		{
			MethodName: "GenerateAnimation",
			Handler:    _Images_GenerateAnimation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "characterimage/v1/characterimage.proto",
}
