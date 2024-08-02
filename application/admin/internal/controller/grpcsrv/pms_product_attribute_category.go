package grpcsrv

import (
	"context"

	"github.com/ccjshop/go-mall/common/retcode"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// CreateProductAttributeCategory 添加产品属性分类表
func (s *AdminApiImpl) CreateProductAttributeCategory(ctx context.Context, param *pb.AddOrUpdateProductAttributeCategoryParam) (*pb.CommonRsp, error) {
	if err := s.productAttributeCategory.CreateProductAttributeCategory(ctx, param); err != nil {
		return nil, err
	}

	res := &pb.CommonRsp{}
	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// UpdateProductAttributeCategory 修改产品属性分类表
func (s *AdminApiImpl) UpdateProductAttributeCategory(ctx context.Context, param *pb.AddOrUpdateProductAttributeCategoryParam) (*pb.CommonRsp, error) {
	if err := s.productAttributeCategory.UpdateProductAttributeCategory(ctx, param); err != nil {
		return nil, err
	}

	res := &pb.CommonRsp{}
	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// GetProductAttributeCategories 分页查询产品属性分类表
func (s *AdminApiImpl) GetProductAttributeCategories(ctx context.Context, param *pb.GetProductAttributeCategoriesParam) (*pb.GetProductAttributeCategoriesRsp, error) {
	productAttributeCategories, pageTotal, err := s.productAttributeCategory.GetProductAttributeCategories(ctx, param)
	if err != nil {
		return nil, err
	}
	res := &pb.GetProductAttributeCategoriesRsp{
		Data: &pb.ProductAttributeCategoriesData{
			Data:      productAttributeCategories,
			PageTotal: pageTotal,
			PageSize:  param.GetPageSize(),
			PageNum:   param.GetPageNum(),
		},
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// GetProductAttributeCategory 根据id获取产品属性分类表
func (s *AdminApiImpl) GetProductAttributeCategory(ctx context.Context, param *pb.GetProductAttributeCategoryReq) (*pb.GetProductAttributeCategoryRsp, error) {
	productAttributeCategory, err := s.productAttributeCategory.GetProductAttributeCategory(ctx, param.GetId())
	if err != nil {
		return nil, err
	}

	res := &pb.GetProductAttributeCategoryRsp{
		Data: productAttributeCategory,
	}
	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// DeleteProductAttributeCategory 删除产品属性分类表
func (s *AdminApiImpl) DeleteProductAttributeCategory(ctx context.Context, param *pb.DeleteProductAttributeCategoryReq) (*pb.CommonRsp, error) {
	if err := s.productAttributeCategory.DeleteProductAttributeCategory(ctx, param.GetId()); err != nil {
		return nil, err
	}

	res := &pb.CommonRsp{}
	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}
