// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package danbooru_v1

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

// GalleryClient is the client API for Gallery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GalleryClient interface {
	Posts(ctx context.Context, in *PostsRequest, opts ...grpc.CallOption) (*PostsResponse, error)
}

type galleryClient struct {
	cc grpc.ClientConnInterface
}

func NewGalleryClient(cc grpc.ClientConnInterface) GalleryClient {
	return &galleryClient{cc}
}

func (c *galleryClient) Posts(ctx context.Context, in *PostsRequest, opts ...grpc.CallOption) (*PostsResponse, error) {
	out := new(PostsResponse)
	err := c.cc.Invoke(ctx, "/danbooru.v1.Gallery/Posts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GalleryServer is the server API for Gallery service.
// All implementations must embed UnimplementedGalleryServer
// for forward compatibility
type GalleryServer interface {
	Posts(context.Context, *PostsRequest) (*PostsResponse, error)
	mustEmbedUnimplementedGalleryServer()
}

// UnimplementedGalleryServer must be embedded to have forward compatible implementations.
type UnimplementedGalleryServer struct {
}

func (UnimplementedGalleryServer) Posts(context.Context, *PostsRequest) (*PostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Posts not implemented")
}
func (UnimplementedGalleryServer) mustEmbedUnimplementedGalleryServer() {}

// UnsafeGalleryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GalleryServer will
// result in compilation errors.
type UnsafeGalleryServer interface {
	mustEmbedUnimplementedGalleryServer()
}

func RegisterGalleryServer(s grpc.ServiceRegistrar, srv GalleryServer) {
	s.RegisterService(&Gallery_ServiceDesc, srv)
}

func _Gallery_Posts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GalleryServer).Posts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/danbooru.v1.Gallery/Posts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GalleryServer).Posts(ctx, req.(*PostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gallery_ServiceDesc is the grpc.ServiceDesc for Gallery service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gallery_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "danbooru.v1.Gallery",
	HandlerType: (*GalleryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Posts",
			Handler:    _Gallery_Posts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apis/danbooru.proto",
}
