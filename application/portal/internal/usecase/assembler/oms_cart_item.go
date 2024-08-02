package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// CartItemModelToEntity pb转entity
func CartItemModelToEntity(cartItemPb *pb.CartItemAddReq) *entity.CartItem {
	if cartItemPb == nil {
		return nil
	}
	return &entity.CartItem{
		// 商品信息
		ProductID: cartItemPb.ProductId,
		// 商品属性
		ProductSkuID:   cartItemPb.ProductSkuId,
		ProductSkuCode: cartItemPb.ProductSkuCode,
		ProductAttr:    cartItemPb.ProductAttr,
		// 价格数量
		Quantity: cartItemPb.Quantity,
	}
}

// CartItemEntityToModel entity转pb
func CartItemEntityToModel(cartItem *entity.CartItem, memberMap map[uint64]*entity.Member) *pb.CartItem {
	if cartItem == nil {
		return nil
	}
	res := &pb.CartItem{
		// 主键
		Id: cartItem.ID,
		// 用户信息
		MemberId: cartItem.MemberID,
		// 商品信息
		ProductId:         cartItem.ProductID,
		ProductName:       cartItem.ProductName,
		ProductPic:        util.ImgUtils.GetFullUrl(cartItem.ProductPic),
		ProductSubTitle:   cartItem.ProductSubTitle,
		ProductSn:         cartItem.ProductSN,
		ProductBrand:      cartItem.ProductBrand,
		ProductCategoryId: cartItem.ProductCategoryID,
		// 商品属性
		ProductSkuId:   cartItem.ProductSkuID,
		ProductSkuCode: cartItem.ProductSkuCode,
		ProductAttr:    cartItem.ProductAttr,
		// 价格数量
		Price:    cartItem.Price.String(),
		Quantity: cartItem.Quantity,
		// 状态
		CreateDate:   cartItem.CreateDate,
		ModifyDate:   cartItem.ModifyDate,
		DeleteStatus: uint32(cartItem.DeleteStatus),
		// 冗余字段
	}
	if member, exist := memberMap[cartItem.MemberID]; exist {
		res.MemberNickname = member.Nickname
	}
	return res
}
