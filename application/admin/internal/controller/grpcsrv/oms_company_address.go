package grpcsrv

import (
	"context"

	"github.com/ccjshop/go-mall/common/retcode"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// CreateCompanyAddress 添加公司收发货地址
func (s *AdminApiImpl) CreateCompanyAddress(ctx context.Context, param *pb.AddOrUpdateCompanyAddressParam) (*pb.CommonRsp, error) {
	var (
		res = &pb.CommonRsp{}
	)

	if err := s.companyAddress.CreateCompanyAddress(ctx, param); err != nil {
		return nil, err
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// UpdateCompanyAddress 修改公司收发货地址
func (s *AdminApiImpl) UpdateCompanyAddress(ctx context.Context, param *pb.AddOrUpdateCompanyAddressParam) (*pb.CommonRsp, error) {
	var (
		res = &pb.CommonRsp{}
	)

	if err := s.companyAddress.UpdateCompanyAddress(ctx, param); err != nil {
		return nil, err
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// GetCompanyAddresses 分页查询公司收发货地址
func (s *AdminApiImpl) GetCompanyAddresses(ctx context.Context, param *pb.GetCompanyAddressesParam) (*pb.GetCompanyAddressesRsp, error) {
	var (
		res = &pb.GetCompanyAddressesRsp{}
	)

	companyAddresses, pageTotal, err := s.companyAddress.GetCompanyAddresses(ctx, param)
	if err != nil {
		return nil, err
	}

	res.Data = &pb.CompanyAddressesData{
		Data:      companyAddresses,
		PageTotal: pageTotal,
		PageSize:  param.GetPageSize(),
		PageNum:   param.GetPageNum(),
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// GetCompanyAddress 根据id获取公司收发货地址
func (s *AdminApiImpl) GetCompanyAddress(ctx context.Context, param *pb.GetCompanyAddressReq) (*pb.GetCompanyAddressRsp, error) {
	var (
		res = &pb.GetCompanyAddressRsp{}
	)

	companyAddress, err := s.companyAddress.GetCompanyAddress(ctx, param.GetId())
	if err != nil {
		return nil, err
	}
	res.Data = companyAddress

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// DeleteCompanyAddress 删除公司收发货地址
func (s *AdminApiImpl) DeleteCompanyAddress(ctx context.Context, param *pb.DeleteCompanyAddressReq) (*pb.CommonRsp, error) {
	var (
		res = &pb.CommonRsp{}
	)

	if err := s.companyAddress.DeleteCompanyAddress(ctx, param.GetId()); err != nil {
		return nil, err
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}
