package grpcsrv

import (
	"context"

	"github.com/ccjshop/go-mall/common/retcode"
	"github.com/ccjshop/go-mall/common/util"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// CartItemAdd 添加商品到购物车
func (s PortalApiImpl) CartItemAdd(ctx context.Context, req *pb.CartItemAddReq) (*pb.CartItemAddRsp, error) {
	res := &pb.CartItemAddRsp{}
	memberID, err := util.CtxUtils.GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	if err := s.cartItemUseCase.CartItemAdd(ctx, memberID, req); err != nil {
		return nil, err
	}
	return res, nil
}

// CartItemList 获取当前会员的购物车列表
func (s PortalApiImpl) CartItemList(ctx context.Context, req *pb.CartItemListReq) (*pb.CartItemListRsp, error) {
	memberID, err := util.CtxUtils.GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	list, err := s.cartItemUseCase.CartItemList(ctx, memberID)
	if err != nil {
		return nil, err
	}

	res := &pb.CartItemListRsp{
		Data: list,
	}
	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, err
}

// CartItemListPromotion 获取当前会员的购物车列表，包括促销信息
func (s PortalApiImpl) CartItemListPromotion(ctx context.Context, req *pb.CartItemListPromotionReq) (*pb.CartItemListPromotionRsp, error) {
	//TODO implement me
	panic("implement me")
}

// CartItemUpdateQuantity 修改购物车中指定商品的数量
func (s PortalApiImpl) CartItemUpdateQuantity(ctx context.Context, req *pb.CartItemUpdateQuantityReq) (*pb.CartItemUpdateQuantityRsp, error) {
	res := &pb.CartItemUpdateQuantityRsp{}
	memberID, err := util.CtxUtils.GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	if err := s.cartItemUseCase.CartItemUpdateQuantity(ctx, memberID, req); err != nil {
		return nil, err
	}
	return res, nil
}

// CartItemGetCartProduct 获取购物车中指定商品的规格，用于重选规格
func (s PortalApiImpl) CartItemGetCartProduct(ctx context.Context, req *pb.CartItemGetCartProductReq) (*pb.CartItemGetCartProductRsp, error) {
	//TODO implement me
	panic("implement me")
}

// CartItemUpdateAttr 修改购物车中商品的规格
func (s PortalApiImpl) CartItemUpdateAttr(ctx context.Context, req *pb.CartItemUpdateAttrReq) (*pb.CartItemUpdateAttrRsp, error) {
	//TODO implement me
	panic("implement me")
}

// CartItemDelete 删除购物车中的指定商品
func (s PortalApiImpl) CartItemDelete(ctx context.Context, req *pb.CartItemDeleteReq) (*pb.CartItemDeleteRsp, error) {
	res := &pb.CartItemDeleteRsp{}
	memberID, err := util.CtxUtils.GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	if err := s.cartItemUseCase.CartItemDelete(ctx, memberID, req); err != nil {
		return nil, err
	}
	return res, nil
}

// CartItemClear 清空当前会员的购物车
func (s PortalApiImpl) CartItemClear(ctx context.Context, req *pb.CartItemClearReq) (*pb.CartItemClearRsp, error) {
	res := &pb.CartItemClearRsp{}
	memberID, err := util.CtxUtils.GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	if err := s.cartItemUseCase.CartItemClear(ctx, memberID); err != nil {
		return nil, err
	}
	return res, nil
}
