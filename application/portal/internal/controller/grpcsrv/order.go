package grpcsrv

import (
	"context"

	"github.com/ccjshop/go-mall/common/util"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// GenerateConfirmOrder 根据购物车信息生成确认单
func (s PortalApiImpl) GenerateConfirmOrder(ctx context.Context, req *pb.GenerateConfirmOrderReq) (*pb.GenerateConfirmOrderRsp, error) {
	memberID, err := util.CtxUtils.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	return s.orderUseCase.GenerateConfirmOrder(ctx, memberID, req)
}

// GenerateOrder 根据购物车信息生成订单
func (s PortalApiImpl) GenerateOrder(ctx context.Context, req *pb.GenerateOrderReq) (*pb.GenerateOrderRsp, error) {
	memberID, err := util.CtxUtils.GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	return s.orderUseCase.GenerateOrder(ctx, memberID, req)
}

// PaySuccess 用户支付成功的回调
func (s PortalApiImpl) PaySuccess(ctx context.Context, req *pb.PaySuccessReq) (*pb.PaySuccessRsp, error) {
	//TODO implement me
	panic("implement me")
}

// CancelTimeOutOrder 自动取消超时订单
func (s PortalApiImpl) CancelTimeOutOrder(ctx context.Context, req *pb.CancelTimeOutOrderReq) (*pb.CancelTimeOutOrderRsp, error) {
	//TODO implement me
	panic("implement me")
}

// CancelOrder 取消单个超时订单
func (s PortalApiImpl) CancelOrder(ctx context.Context, req *pb.CancelOrderReq) (*pb.CancelOrderRsp, error) {
	//TODO implement me
	panic("implement me")
}

// OrderList 按状态分页获取用户订单列表
func (s PortalApiImpl) OrderList(ctx context.Context, req *pb.OrderListReq) (*pb.OrderListRsp, error) {
	//TODO implement me
	panic("implement me")
}

// OrderDetail 根据ID获取订单详情
func (s PortalApiImpl) OrderDetail(ctx context.Context, req *pb.OrderDetailReq) (*pb.OrderDetailRsp, error) {
	//TODO implement me
	panic("implement me")
}

// CancelUserOrder 用户取消订单
func (s PortalApiImpl) CancelUserOrder(ctx context.Context, req *pb.CancelUserOrderReq) (*pb.CancelUserOrderRsp, error) {
	//TODO implement me
	panic("implement me")
}

// ConfirmReceiveOrder 用户确认收货
func (s PortalApiImpl) ConfirmReceiveOrder(ctx context.Context, req *pb.ConfirmReceiveOrderReq) (*pb.ConfirmReceiveOrderRsp, error) {
	//TODO implement me
	panic("implement me")
}

// DeleteOrder 用户删除订单
func (s PortalApiImpl) DeleteOrder(ctx context.Context, req *pb.PortalDeleteOrderReq) (*pb.PortalDeleteOrderRsp, error) {
	//TODO implement me
	panic("implement me")
}
