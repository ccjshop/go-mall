package grpcsrv

import (
	"context"

	"github.com/ccjshop/go-mall/common/retcode"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// CreateBrand 添加商品品牌
func (s *AdminApiImpl) CreateBrand(ctx context.Context, param *pb.AddOrUpdateBrandParam) (*pb.CommonRsp, error) {
	if err := s.brand.CreateBrand(ctx, param); err != nil {
		return nil, err
	}

	res := &pb.CommonRsp{}
	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// UpdateBrand 修改商品品牌
func (s *AdminApiImpl) UpdateBrand(ctx context.Context, param *pb.AddOrUpdateBrandParam) (*pb.CommonRsp, error) {
	if err := s.brand.UpdateBrand(ctx, param); err != nil {
		return nil, err
	}

	res := &pb.CommonRsp{}
	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// GetBrands 分页查询商品品牌
func (s *AdminApiImpl) GetBrands(ctx context.Context, param *pb.GetBrandsParam) (*pb.GetBrandsRsp, error) {
	brands, pageTotal, err := s.brand.GetBrands(ctx, param)
	if err != nil {
		return nil, err
	}

	res := &pb.GetBrandsRsp{
		Data: &pb.GetBrandsData{
			Data:      brands,
			PageTotal: pageTotal,
			PageNum:   param.PageNum,
			PageSize:  param.PageSize,
		},
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// GetBrand 根据id获取商品品牌
func (s *AdminApiImpl) GetBrand(ctx context.Context, param *pb.GetBrandReq) (*pb.GetBrandRsp, error) {
	brand, err := s.brand.GetBrand(ctx, param.GetId())
	if err != nil {
		return nil, err
	}

	res := &pb.GetBrandRsp{
		Data: brand,
	}
	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// DeleteBrand 删除商品品牌
func (s *AdminApiImpl) DeleteBrand(ctx context.Context, param *pb.DeleteBrandReq) (*pb.CommonRsp, error) {
	err := s.brand.DeleteBrand(ctx, param.GetId())
	if err != nil {
		return nil, err
	}

	res := &pb.CommonRsp{}
	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}
