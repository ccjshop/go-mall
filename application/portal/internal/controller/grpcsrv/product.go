package grpcsrv

import (
	"context"

	pb "github.com/ccjshop/go-mall/proto/mall"
)

// SearchProduct 综合搜索商品
func (s PortalApiImpl) SearchProduct(ctx context.Context, req *pb.SearchProductReq) (*pb.SearchProductRsp, error) {
	var (
		res = &pb.SearchProductRsp{}
	)
	products, err := s.product.SearchProduct(ctx, req)
	if err != nil {
		return nil, err
	}

	res.Data = products

	return res, nil
}

// CategoryTreeList 以树形结构获取所有商品分类
func (s PortalApiImpl) CategoryTreeList(context.Context, *pb.CategoryTreeListReq) (*pb.CategoryTreeListRsp, error) {
	var (
		res = &pb.CategoryTreeListRsp{}
	)
	return res, nil
}

// ProductDetail 获取前台商品详情
func (s PortalApiImpl) ProductDetail(ctx context.Context, req *pb.ProductDetailReq) (*pb.ProductDetailRsp, error) {
	product, err := s.product.ProductDetail(ctx, req)
	if err != nil {
		return nil, err
	}
	return product, nil
}
