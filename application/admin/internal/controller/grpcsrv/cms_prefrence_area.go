package grpcsrv

import (
	"context"

	"github.com/ccjshop/go-mall/common/retcode"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// CreatePrefrenceArea 添加优选专区
func (s *AdminApiImpl) CreatePrefrenceArea(ctx context.Context, param *pb.AddOrUpdatePrefrenceAreaParam) (*pb.CommonRsp, error) {
	var (
		res = &pb.CommonRsp{}
	)

	if err := s.prefrenceArea.CreatePrefrenceArea(ctx, param); err != nil {
		return nil, err
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// UpdatePrefrenceArea 修改优选专区
func (s *AdminApiImpl) UpdatePrefrenceArea(ctx context.Context, param *pb.AddOrUpdatePrefrenceAreaParam) (*pb.CommonRsp, error) {
	var (
		res = &pb.CommonRsp{}
	)

	if err := s.prefrenceArea.UpdatePrefrenceArea(ctx, param); err != nil {
		return nil, err
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// GetPrefrenceAreas 分页查询优选专区
func (s *AdminApiImpl) GetPrefrenceAreas(ctx context.Context, param *pb.GetPrefrenceAreasParam) (*pb.GetPrefrenceAreasRsp, error) {
	var (
		res = &pb.GetPrefrenceAreasRsp{}
	)

	prefrenceAreas, pageTotal, err := s.prefrenceArea.GetPrefrenceAreas(ctx, param)
	if err != nil {
		return nil, err
	}

	res.Data = &pb.PrefrenceAreasData{
		Data:      prefrenceAreas,
		PageTotal: pageTotal,
		PageSize:  param.GetPageSize(),
		PageNum:   param.GetPageNum(),
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// GetPrefrenceArea 根据id获取优选专区
func (s *AdminApiImpl) GetPrefrenceArea(ctx context.Context, param *pb.GetPrefrenceAreaReq) (*pb.GetPrefrenceAreaRsp, error) {
	var (
		res = &pb.GetPrefrenceAreaRsp{}
	)

	prefrenceArea, err := s.prefrenceArea.GetPrefrenceArea(ctx, param.GetId())
	if err != nil {
		return nil, err
	}
	res.Data = prefrenceArea

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// DeletePrefrenceArea 删除优选专区
func (s *AdminApiImpl) DeletePrefrenceArea(ctx context.Context, param *pb.DeletePrefrenceAreaReq) (*pb.CommonRsp, error) {
	var (
		res = &pb.CommonRsp{}
	)

	if err := s.prefrenceArea.DeletePrefrenceArea(ctx, param.GetId()); err != nil {
		return nil, err
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}
