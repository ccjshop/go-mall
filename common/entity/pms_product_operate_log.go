package entity

import "github.com/shopspring/decimal"

// ProductOperateLog 商品操作记录表
// 用于记录商品操作记录
type ProductOperateLog struct {
	ID               uint64          `gorm:"column:id;type:bigint;primary_key;auto_increment;comment:主键"`
	ProductID        uint64          `gorm:"column:product_id;type:bigint;not null;default:0;comment:商品id"`
	PriceOld         decimal.Decimal `gorm:"column:price_old;type:decimal(10,2);not null;default:0.00;comment:改变前价格"`
	PriceNew         decimal.Decimal `gorm:"column:price_new;type:decimal(10,2);not null;default:0.00;comment:改变后价格"`
	SalePriceOld     decimal.Decimal `gorm:"column:sale_price_old;type:decimal(10,2);not null;default:0.00;comment:改变前优惠价"`
	SalePriceNew     decimal.Decimal `gorm:"column:sale_price_new;type:decimal(10,2);not null;default:0.00;comment:改变后优惠价"`
	GiftPointOld     uint32          `gorm:"column:gift_point_old;type:int(10);unsigned;not null;default:0;comment:改变前积分"`
	GiftPointNew     uint32          `gorm:"column:gift_point_new;type:int(10);unsigned;not null;default:0;comment:改变后积分"`
	UsePointLimitOld uint32          `gorm:"column:use_point_limit_old;type:int(10);unsigned;not null;default:0;comment:改变前积分使用限制"`
	UsePointLimitNew uint32          `gorm:"column:use_point_limit_new;type:int(10);unsigned;not null;default:0;comment:改变后积分使用限制"`
	OperateMan       string          `gorm:"column:operate_man;type:varchar(64);not null;default:'';comment:操作人"`
	CreateTime       uint32          `gorm:"column:create_time;type:int(10);unsigned;not null;default:0;comment:创建时间"`
	// 公共字段
	BaseTime
}

func (c ProductOperateLog) TableName() string {
	return "pms_product_operate_log"
}
