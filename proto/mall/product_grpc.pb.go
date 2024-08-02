// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: portal/product.proto

package mall

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

const (
	PortalProductApi_SearchProduct_FullMethodName    = "/admin.PortalProductApi/SearchProduct"
	PortalProductApi_CategoryTreeList_FullMethodName = "/admin.PortalProductApi/CategoryTreeList"
	PortalProductApi_ProductDetail_FullMethodName    = "/admin.PortalProductApi/ProductDetail"
)

// PortalProductApiClient is the client API for PortalProductApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortalProductApiClient interface {
	// 综合搜索商品
	SearchProduct(ctx context.Context, in *SearchProductReq, opts ...grpc.CallOption) (*SearchProductRsp, error)
	// 以树形结构获取所有商品分类
	CategoryTreeList(ctx context.Context, in *CategoryTreeListReq, opts ...grpc.CallOption) (*CategoryTreeListRsp, error)
	// 获取前台商品详情
	ProductDetail(ctx context.Context, in *ProductDetailReq, opts ...grpc.CallOption) (*ProductDetailRsp, error)
}

type portalProductApiClient struct {
	cc grpc.ClientConnInterface
}

func NewPortalProductApiClient(cc grpc.ClientConnInterface) PortalProductApiClient {
	return &portalProductApiClient{cc}
}

func (c *portalProductApiClient) SearchProduct(ctx context.Context, in *SearchProductReq, opts ...grpc.CallOption) (*SearchProductRsp, error) {
	out := new(SearchProductRsp)
	err := c.cc.Invoke(ctx, PortalProductApi_SearchProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalProductApiClient) CategoryTreeList(ctx context.Context, in *CategoryTreeListReq, opts ...grpc.CallOption) (*CategoryTreeListRsp, error) {
	out := new(CategoryTreeListRsp)
	err := c.cc.Invoke(ctx, PortalProductApi_CategoryTreeList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalProductApiClient) ProductDetail(ctx context.Context, in *ProductDetailReq, opts ...grpc.CallOption) (*ProductDetailRsp, error) {
	out := new(ProductDetailRsp)
	err := c.cc.Invoke(ctx, PortalProductApi_ProductDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortalProductApiServer is the server API for PortalProductApi service.
// All implementations must embed UnimplementedPortalProductApiServer
// for forward compatibility
type PortalProductApiServer interface {
	// 综合搜索商品
	SearchProduct(context.Context, *SearchProductReq) (*SearchProductRsp, error)
	// 以树形结构获取所有商品分类
	CategoryTreeList(context.Context, *CategoryTreeListReq) (*CategoryTreeListRsp, error)
	// 获取前台商品详情
	ProductDetail(context.Context, *ProductDetailReq) (*ProductDetailRsp, error)
	mustEmbedUnimplementedPortalProductApiServer()
}

// UnimplementedPortalProductApiServer must be embedded to have forward compatible implementations.
type UnimplementedPortalProductApiServer struct {
}

func (UnimplementedPortalProductApiServer) SearchProduct(context.Context, *SearchProductReq) (*SearchProductRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchProduct not implemented")
}
func (UnimplementedPortalProductApiServer) CategoryTreeList(context.Context, *CategoryTreeListReq) (*CategoryTreeListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategoryTreeList not implemented")
}
func (UnimplementedPortalProductApiServer) ProductDetail(context.Context, *ProductDetailReq) (*ProductDetailRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductDetail not implemented")
}
func (UnimplementedPortalProductApiServer) mustEmbedUnimplementedPortalProductApiServer() {}

// UnsafePortalProductApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortalProductApiServer will
// result in compilation errors.
type UnsafePortalProductApiServer interface {
	mustEmbedUnimplementedPortalProductApiServer()
}

func RegisterPortalProductApiServer(s grpc.ServiceRegistrar, srv PortalProductApiServer) {
	s.RegisterService(&PortalProductApi_ServiceDesc, srv)
}

func _PortalProductApi_SearchProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalProductApiServer).SearchProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortalProductApi_SearchProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalProductApiServer).SearchProduct(ctx, req.(*SearchProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortalProductApi_CategoryTreeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryTreeListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalProductApiServer).CategoryTreeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortalProductApi_CategoryTreeList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalProductApiServer).CategoryTreeList(ctx, req.(*CategoryTreeListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortalProductApi_ProductDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalProductApiServer).ProductDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortalProductApi_ProductDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalProductApiServer).ProductDetail(ctx, req.(*ProductDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

// PortalProductApi_ServiceDesc is the grpc.ServiceDesc for PortalProductApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PortalProductApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.PortalProductApi",
	HandlerType: (*PortalProductApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchProduct",
			Handler:    _PortalProductApi_SearchProduct_Handler,
		},
		{
			MethodName: "CategoryTreeList",
			Handler:    _PortalProductApi_CategoryTreeList_Handler,
		},
		{
			MethodName: "ProductDetail",
			Handler:    _PortalProductApi_ProductDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "portal/product.proto",
}