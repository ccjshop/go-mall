package grpcsrv

import (
	"context"

	"github.com/ccjshop/go-mall/common/retcode"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// CreateProductAttribute 添加商品属性参数表
func (s *AdminApiImpl) CreateProductAttribute(ctx context.Context, param *pb.AddOrUpdateProductAttributeParam) (*pb.CommonRsp, error) {
	if err := s.productAttribute.CreateProductAttribute(ctx, param); err != nil {
		return nil, err
	}

	res := &pb.CommonRsp{}
	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// UpdateProductAttribute 修改商品属性参数表
func (s *AdminApiImpl) UpdateProductAttribute(ctx context.Context, param *pb.AddOrUpdateProductAttributeParam) (*pb.CommonRsp, error) {
	if err := s.productAttribute.UpdateProductAttribute(ctx, param); err != nil {
		return nil, err
	}

	res := &pb.CommonRsp{}
	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// GetProductAttributes 分页查询商品属性参数表
func (s *AdminApiImpl) GetProductAttributes(ctx context.Context, param *pb.GetProductAttributesParam) (*pb.GetProductAttributesRsp, error) {
	productAttributes, pageTotal, err := s.productAttribute.GetProductAttributes(ctx, param)
	if err != nil {
		return nil, err
	}
	res := &pb.GetProductAttributesRsp{
		Data: &pb.ProductAttributesData{
			Data:      productAttributes,
			PageTotal: pageTotal,
			PageSize:  param.GetPageSize(),
			PageNum:   param.GetPageNum(),
		},
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// GetProductAttribute 根据id获取商品属性参数表
func (s *AdminApiImpl) GetProductAttribute(ctx context.Context, param *pb.GetProductAttributeReq) (*pb.GetProductAttributeRsp, error) {
	productAttribute, err := s.productAttribute.GetProductAttribute(ctx, param.GetId())
	if err != nil {
		return nil, err
	}

	res := &pb.GetProductAttributeRsp{
		Data: productAttribute,
	}
	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// DeleteProductAttribute 删除商品属性参数表
func (s *AdminApiImpl) DeleteProductAttribute(ctx context.Context, param *pb.DeleteProductAttributeReq) (*pb.CommonRsp, error) {
	if err := s.productAttribute.DeleteProductAttribute(ctx, param.GetId()); err != nil {
		return nil, err
	}

	res := &pb.CommonRsp{}
	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}
