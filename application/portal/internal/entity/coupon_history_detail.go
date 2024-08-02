package entity

import g_entity "github.com/ccjshop/go-mall/common/entity"

// CouponHistoryDetail 优惠券领取历史详情（包括优惠券信息和关联关系）
type CouponHistoryDetail struct {
	*g_entity.CouponHistory
	Coupon            *g_entity.Coupon                        // 相关优惠券信息
	ProductRelations  g_entity.CouponProductRelations         // 优惠券关联商品
	CategoryRelations g_entity.CouponProductCategoryRelations // 优惠券关联商品分类
}
