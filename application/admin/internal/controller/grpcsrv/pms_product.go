package grpcsrv

import (
	"context"

	"github.com/ccjshop/go-mall/common/retcode"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// CreateProduct 添加商品
func (s *AdminApiImpl) CreateProduct(ctx context.Context, param *pb.AddOrUpdateProductParam) (*pb.CommonRsp, error) {
	if err := s.product.CreateProduct(ctx, param); err != nil {
		return nil, err
	}

	res := &pb.CommonRsp{}
	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// UpdateProduct 修改商品
func (s *AdminApiImpl) UpdateProduct(ctx context.Context, param *pb.AddOrUpdateProductParam) (*pb.CommonRsp, error) {
	if err := s.product.UpdateProduct(ctx, param); err != nil {
		return nil, err
	}

	res := &pb.CommonRsp{}
	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// GetProducts 分页查询商品
func (s *AdminApiImpl) GetProducts(ctx context.Context, param *pb.GetProductsParam) (*pb.GetProductsRsp, error) {
	products, pageTotal, err := s.product.GetProducts(ctx, param)
	if err != nil {
		return nil, err
	}

	res := &pb.GetProductsRsp{
		Data: &pb.ProductsData{
			Data:      products,
			PageTotal: pageTotal,
			PageSize:  param.GetPageSize(),
			PageNum:   param.GetPageNum(),
		},
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// GetProduct 根据id获取商品
func (s *AdminApiImpl) GetProduct(ctx context.Context, param *pb.GetProductReq) (*pb.GetProductRsp, error) {
	product, err := s.product.GetProduct(ctx, param.GetId())
	if err != nil {
		return nil, err
	}

	res := &pb.GetProductRsp{
		Data: product,
	}
	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// DeleteProduct 删除商品
func (s *AdminApiImpl) DeleteProduct(ctx context.Context, param *pb.DeleteProductReq) (*pb.CommonRsp, error) {
	if err := s.product.DeleteProduct(ctx, param.GetId()); err != nil {
		return nil, err
	}

	res := &pb.CommonRsp{}
	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}
