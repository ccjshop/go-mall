package grpcsrv

import (
	"context"

	"github.com/ccjshop/go-mall/common/retcode"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// CreateJsonDynamicConfig 添加JSON动态配置
func (s *AdminApiImpl) CreateJsonDynamicConfig(ctx context.Context, param *pb.AddOrUpdateJsonDynamicConfigParam) (*pb.CommonRsp, error) {
	var (
		res = &pb.CommonRsp{}
	)

	if err := s.jsonDynamicConfig.CreateJsonDynamicConfig(ctx, param); err != nil {
		return nil, err
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// UpdateJsonDynamicConfig 修改JSON动态配置
func (s *AdminApiImpl) UpdateJsonDynamicConfig(ctx context.Context, param *pb.AddOrUpdateJsonDynamicConfigParam) (*pb.CommonRsp, error) {
	var (
		res = &pb.CommonRsp{}
	)

	if err := s.jsonDynamicConfig.UpdateJsonDynamicConfig(ctx, param); err != nil {
		return nil, err
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// GetJsonDynamicConfigs 分页查询JSON动态配置
func (s *AdminApiImpl) GetJsonDynamicConfigs(ctx context.Context, param *pb.GetJsonDynamicConfigsParam) (*pb.GetJsonDynamicConfigsRsp, error) {
	var (
		res = &pb.GetJsonDynamicConfigsRsp{}
	)

	jsonDynamicConfigs, pageTotal, err := s.jsonDynamicConfig.GetJsonDynamicConfigs(ctx, param)
	if err != nil {
		return nil, err
	}

	res.Data = &pb.JsonDynamicConfigsData{
		Data:      jsonDynamicConfigs,
		PageTotal: pageTotal,
		PageSize:  param.GetPageSize(),
		PageNum:   param.GetPageNum(),
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// GetJsonDynamicConfig 根据id获取JSON动态配置
func (s *AdminApiImpl) GetJsonDynamicConfig(ctx context.Context, param *pb.GetJsonDynamicConfigReq) (*pb.GetJsonDynamicConfigRsp, error) {
	var (
		res = &pb.GetJsonDynamicConfigRsp{}
	)

	jsonDynamicConfig, err := s.jsonDynamicConfig.GetJsonDynamicConfig(ctx, param.GetId())
	if err != nil {
		return nil, err
	}
	res.Data = jsonDynamicConfig

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// DeleteJsonDynamicConfig 删除JSON动态配置
func (s *AdminApiImpl) DeleteJsonDynamicConfig(ctx context.Context, param *pb.DeleteJsonDynamicConfigReq) (*pb.CommonRsp, error) {
	var (
		res = &pb.CommonRsp{}
	)

	if err := s.jsonDynamicConfig.DeleteJsonDynamicConfig(ctx, param.GetId()); err != nil {
		return nil, err
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}
