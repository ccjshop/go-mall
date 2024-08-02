package grpcsrv

import (
	"context"

	"github.com/ccjshop/go-mall/common/retcode"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// CreateHomeAdvertise 添加首页轮播广告表
func (s *AdminApiImpl) CreateHomeAdvertise(ctx context.Context, param *pb.AddOrUpdateHomeAdvertiseParam) (*pb.CommonRsp, error) {
	var (
		res = &pb.CommonRsp{}
	)

	if err := s.homeAdvertise.CreateHomeAdvertise(ctx, param); err != nil {
		return nil, err
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// UpdateHomeAdvertise 修改首页轮播广告表
func (s *AdminApiImpl) UpdateHomeAdvertise(ctx context.Context, param *pb.AddOrUpdateHomeAdvertiseParam) (*pb.CommonRsp, error) {
	var (
		res = &pb.CommonRsp{}
	)

	if err := s.homeAdvertise.UpdateHomeAdvertise(ctx, param); err != nil {
		return nil, err
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// GetHomeAdvertises 分页查询首页轮播广告表
func (s *AdminApiImpl) GetHomeAdvertises(ctx context.Context, param *pb.GetHomeAdvertisesParam) (*pb.GetHomeAdvertisesRsp, error) {
	var (
		res = &pb.GetHomeAdvertisesRsp{}
	)

	homeAdvertises, pageTotal, err := s.homeAdvertise.GetHomeAdvertises(ctx, param)
	if err != nil {
		return nil, err
	}

	res.Data = &pb.HomeAdvertisesData{
		Data:      homeAdvertises,
		PageTotal: pageTotal,
		PageSize:  param.GetPageSize(),
		PageNum:   param.GetPageNum(),
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// GetHomeAdvertise 根据id获取首页轮播广告表
func (s *AdminApiImpl) GetHomeAdvertise(ctx context.Context, param *pb.GetHomeAdvertiseReq) (*pb.GetHomeAdvertiseRsp, error) {
	var (
		res = &pb.GetHomeAdvertiseRsp{}
	)

	homeAdvertise, err := s.homeAdvertise.GetHomeAdvertise(ctx, param.GetId())
	if err != nil {
		return nil, err
	}
	res.Data = homeAdvertise

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// DeleteHomeAdvertise 删除首页轮播广告表
func (s *AdminApiImpl) DeleteHomeAdvertise(ctx context.Context, param *pb.DeleteHomeAdvertiseReq) (*pb.CommonRsp, error) {
	var (
		res = &pb.CommonRsp{}
	)

	if err := s.homeAdvertise.DeleteHomeAdvertise(ctx, param.GetId()); err != nil {
		return nil, err
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}
