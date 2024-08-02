package entity

import "github.com/shopspring/decimal"

// ProductFullReduction 商品满减表
// 商品优惠相关表，购买同商品满足一定金额后，可以减免一定金额。如：买满1000减100元。
type ProductFullReduction struct {
	ID        uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;comment:主键"`
	ProductID uint64 `gorm:"column:product_id;type:bigint;unsigned;not null;default:0;comment:商品id"` // pms_product#id
	// 优惠
	FullPrice   decimal.Decimal `gorm:"column:full_price;type:decimal(10,2);not null;default:0.00;comment:商品满足金额"`
	ReducePrice decimal.Decimal `gorm:"column:reduce_price;type:decimal(10,2);not null;default:0.00;comment:商品减少金额"`
	// 公共字段
	BaseTime
}

func (c ProductFullReduction) TableName() string {
	return "pms_product_full_reduction"
}

type ProductFullReductions []*ProductFullReduction
