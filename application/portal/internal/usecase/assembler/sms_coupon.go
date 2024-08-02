package assembler

import (
	portal_entity "github.com/ccjshop/go-mall/application/portal/internal/entity"
	"github.com/ccjshop/go-mall/common/entity"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// CouponEntityToModel entity转pb
func CouponEntityToModel(coupon *entity.Coupon) *pb.Coupon {
	return &pb.Coupon{}
}

// CouponHistoryDetailToModel entity转pb
func CouponHistoryDetailToModel(couponHistoryDetails []*portal_entity.CouponHistoryDetail) []*pb.CouponHistoryDetail {
	res := make([]*pb.CouponHistoryDetail, 0)
	for _, detail := range couponHistoryDetails {
		item := &pb.CouponHistoryDetail{
			// 一、CouponHistory
			Id:             detail.CouponHistory.ID,
			MemberId:       detail.CouponHistory.MemberID,
			CouponId:       detail.CouponHistory.CouponID,
			CouponCode:     detail.CouponHistory.CouponCode,
			OrderId:        detail.CouponHistory.OrderID,
			OrderSn:        detail.CouponHistory.OrderSN,
			GetType:        uint32(detail.CouponHistory.GetType),
			UseStatus:      uint32(detail.CouponHistory.UseStatus),
			CreateTime:     detail.CouponHistory.CreateTime,
			UseTime:        detail.CouponHistory.UseTime,
			MemberNickname: detail.CouponHistory.MemberNickname,
		}
		item.Coupon = &pb.CouponHistoryDetail_Coupon{
			Id:      detail.Coupon.ID,
			Name:    detail.Coupon.Name,
			Amount:  detail.Coupon.Amount.String(),
			Note:    detail.Coupon.Note,
			Code:    detail.Coupon.Code,
			Type:    uint32(detail.Coupon.Type),
			UseType: detail.Coupon.UseType,
			EndTime: detail.Coupon.EndTime,
		}
		for _, relation := range detail.ProductRelations {
			item.ProductRelations = append(item.ProductRelations, &pb.CouponHistoryDetail_CouponProductRelation{
				CouponId:    relation.CouponID,
				ProductId:   relation.ProductID,
				ProductName: relation.ProductName,
				ProductSn:   relation.ProductSN,
			})
		}
		for _, relation := range detail.CategoryRelations {
			item.CategoryRelations = append(item.CategoryRelations, &pb.CouponHistoryDetail_CouponProductCategoryRelation{
				CouponId:            relation.CouponID,
				ProductCategoryId:   relation.ProductCategoryID,
				ProductCategoryName: relation.ProductCategoryName,
				ParentCategoryName:  relation.ParentCategoryName,
			})
		}
		res = append(res, item)
	}
	return res
}
