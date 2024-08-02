package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// OrderItemsToModel entity转pb
func OrderItemsToModel(orderItems []*entity.OrderItem) []*pb.OrderItem {
	res := make([]*pb.OrderItem, 0)
	for _, orderItem := range orderItems {
		res = append(res, OrderItemToModel(orderItem))
	}
	return res
}

// OrderItemToModel entity转pb
func OrderItemToModel(orderItem *entity.OrderItem) *pb.OrderItem {
	return &pb.OrderItem{
		Id:      orderItem.ID,
		OrderId: orderItem.OrderID,
		OrderSn: orderItem.OrderSN,
		// 商品信息
		ProductId:         orderItem.ProductID,
		ProductPic:        util.ImgUtils.GetFullUrl(orderItem.ProductPic),
		ProductName:       orderItem.ProductName,
		ProductBrand:      orderItem.ProductBrand,
		ProductPrice:      orderItem.ProductPrice.String(),
		ProductSn:         orderItem.ProductSN,
		ProductAttr:       orderItem.ProductAttr,
		ProductQuantity:   orderItem.ProductQuantity,
		ProductCategoryId: orderItem.ProductCategoryID,
		ProductSkuId:      orderItem.ProductSkuID,
		ProductSkuCode:    orderItem.ProductSkuCode,
		PromotionName:     orderItem.PromotionName,
		PromotionAmount:   orderItem.PromotionAmount.String(),
		CouponAmount:      orderItem.CouponAmount.String(),
		IntegrationAmount: orderItem.IntegrationAmount.String(),
		RealAmount:        orderItem.RealAmount.String(),
		GiftIntegration:   orderItem.GiftIntegration,
		GiftGrowth:        orderItem.GiftGrowth,
	}
}
