package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// OrderEntityToModel entity转pb
func OrderEntityToModel(order *entity.Order) *pb.Order {
	return &pb.Order{
		// 基本信息
		Id:              order.ID,
		OrderSn:         order.OrderSN,
		MemberId:        order.MemberID,
		PayType:         uint32(order.PayType),
		SourceType:      uint32(order.SourceType),
		OrderType:       uint32(order.OrderType),
		DeliveryCompany: order.DeliveryCompany,
		DeliverySn:      order.DeliverySN,
		AutoConfirmDay:  order.AutoConfirmDay,
		ReceiveTime:     order.ReceiveTime,
		Integration:     order.Integration,
		Growth:          order.Growth,
		PromotionInfo:   order.PromotionInfo,
		Note:            order.Note,
		// 收货人信息
		ReceiverName:          order.ReceiverName,
		ReceiverPhone:         order.ReceiverPhone,
		ReceiverPostCode:      order.ReceiverPostCode,
		ReceiverProvince:      order.ReceiverProvince,
		ReceiverCity:          order.ReceiverCity,
		ReceiverRegion:        order.ReceiverRegion,
		ReceiverDetailAddress: order.ReceiverDetailAddress,
		// 费用信息
		TotalAmount:       order.TotalAmount.String(),
		FreightAmount:     order.FreightAmount.String(),
		CouponId:          order.CouponID,
		CouponAmount:      order.CouponAmount.String(),
		UseIntegration:    order.UseIntegration,
		IntegrationAmount: order.IntegrationAmount.String(),
		PromotionAmount:   order.PromotionAmount.String(),
		DiscountAmount:    order.DiscountAmount.String(),
		PayAmount:         order.PayAmount.String(),
		// 发票信息
		BillType:          uint32(order.BillType),
		BillHeader:        order.BillHeader,
		BillContent:       order.BillContent,
		BillReceiverPhone: order.BillReceiverPhone,
		BillReceiverEmail: order.BillReceiverEmail,
		// 状态
		Status:        uint32(order.Status),
		ConfirmStatus: uint32(order.ConfirmStatus),
		DeleteStatus:  uint32(order.DeleteStatus),
		// 时间
		PaymentTime:  order.PaymentTime,
		DeliveryTime: order.DeliveryTime,
		CommentTime:  order.CommentTime,
		CreateTime:   order.CreateTime,
		ModifyTime:   order.ModifyTime,
		// 冗余字段
		MemberUsername: order.MemberUsername,
	}
}
