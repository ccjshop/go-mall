package grpcsrv

import (
	"context"

	"github.com/ccjshop/go-mall/common/retcode"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// GetOrders 分页查询订单
func (s *AdminApiImpl) GetOrders(ctx context.Context, param *pb.GetOrdersParam) (*pb.GetOrdersRsp, error) {
	var (
		res = &pb.GetOrdersRsp{}
	)

	orders, pageTotal, err := s.order.GetOrders(ctx, param)
	if err != nil {
		return nil, err
	}

	res.Data = &pb.OrdersData{
		Data:      orders,
		PageTotal: pageTotal,
		PageSize:  param.GetPageSize(),
		PageNum:   param.GetPageNum(),
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// GetOrder 根据id获取订单
func (s *AdminApiImpl) GetOrder(ctx context.Context, param *pb.GetOrderReq) (*pb.GetOrderRsp, error) {
	var (
		res = &pb.GetOrderRsp{}
	)

	order, err := s.order.GetOrder(ctx, param.GetId())
	if err != nil {
		return nil, err
	}
	res.Data = order

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// DeleteOrder 删除订单
func (s *AdminApiImpl) DeleteOrder(ctx context.Context, param *pb.DeleteOrderReq) (*pb.CommonRsp, error) {
	var (
		res = &pb.CommonRsp{}
	)

	if err := s.order.DeleteOrder(ctx, param.GetId()); err != nil {
		return nil, err
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}
