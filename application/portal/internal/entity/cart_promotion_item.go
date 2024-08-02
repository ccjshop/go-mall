package entity

import (
	g_entity "github.com/ccjshop/go-mall/common/entity"
	"github.com/shopspring/decimal"
)

// CartPromotionItem 购物车中促销信息的封装
type CartPromotionItem struct {
	// 购物车信息
	g_entity.CartItem
	// 促销字段
	PromotionMessage string          // 促销活动信息
	ReduceAmount     decimal.Decimal // 促销活动减去的金额，针对每个商品
	RealStock        uint32          // 剩余库存-锁定库存
	Integration      uint32          // 购买商品赠送积分
	Growth           uint32          // 购买商品赠送成长值
}

type CartPromotionItems []*CartPromotionItem
