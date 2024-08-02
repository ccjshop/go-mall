package grpcsrv

import (
	"context"

	"github.com/ccjshop/go-mall/common/retcode"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// CreateOrderReturnReason 添加退货原因
func (s *AdminApiImpl) CreateOrderReturnReason(ctx context.Context, param *pb.AddOrUpdateOrderReturnReasonParam) (*pb.CommonRsp, error) {
	var (
		res = &pb.CommonRsp{}
	)

	if err := s.orderReturnReason.CreateOrderReturnReason(ctx, param); err != nil {
		return nil, err
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// UpdateOrderReturnReason 修改退货原因
func (s *AdminApiImpl) UpdateOrderReturnReason(ctx context.Context, param *pb.AddOrUpdateOrderReturnReasonParam) (*pb.CommonRsp, error) {
	var (
		res = &pb.CommonRsp{}
	)

	if err := s.orderReturnReason.UpdateOrderReturnReason(ctx, param); err != nil {
		return nil, err
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// GetOrderReturnReasons 分页查询退货原因
func (s *AdminApiImpl) GetOrderReturnReasons(ctx context.Context, param *pb.GetOrderReturnReasonsParam) (*pb.GetOrderReturnReasonsRsp, error) {
	var (
		res = &pb.GetOrderReturnReasonsRsp{}
	)

	orderReturnReasons, pageTotal, err := s.orderReturnReason.GetOrderReturnReasons(ctx, param)
	if err != nil {
		return nil, err
	}

	res.Data = &pb.OrderReturnReasonsData{
		Data:      orderReturnReasons,
		PageTotal: pageTotal,
		PageSize:  param.GetPageSize(),
		PageNum:   param.GetPageNum(),
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// GetOrderReturnReason 根据id获取退货原因
func (s *AdminApiImpl) GetOrderReturnReason(ctx context.Context, param *pb.GetOrderReturnReasonReq) (*pb.GetOrderReturnReasonRsp, error) {
	var (
		res = &pb.GetOrderReturnReasonRsp{}
	)

	orderReturnReason, err := s.orderReturnReason.GetOrderReturnReason(ctx, param.GetId())
	if err != nil {
		return nil, err
	}
	res.Data = orderReturnReason

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// DeleteOrderReturnReason 删除退货原因
func (s *AdminApiImpl) DeleteOrderReturnReason(ctx context.Context, param *pb.DeleteOrderReturnReasonReq) (*pb.CommonRsp, error) {
	var (
		res = &pb.CommonRsp{}
	)

	if err := s.orderReturnReason.DeleteOrderReturnReason(ctx, param.GetId()); err != nil {
		return nil, err
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}
