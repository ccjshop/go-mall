package entity

import "github.com/shopspring/decimal"

// FeightTemplate 运费模版
type FeightTemplate struct {
	ID             uint64          `gorm:"column:id;type:bigint;primary_key;auto_increment;comment:主键"`
	Name           string          `gorm:"column:name;type:varchar(64);not null;default:'';comment:模版名称"`
	ChargeType     uint8           `gorm:"column:charge_type;type:tinyint(4);unsigned;not null;default:0;comment:计费类型:0->按重量；1->按件数"`
	FirstWeight    decimal.Decimal `gorm:"column:first_weight;type:decimal(10,2);not null;default:0.00;comment:首重kg"`
	FirstFee       decimal.Decimal `gorm:"column:first_fee;type:decimal(10,2);not null;default:0.00;comment:首费（元）"`
	ContinueWeight decimal.Decimal `gorm:"column:continue_weight;type:decimal(10,2);not null;default:0.00;comment:续重kg"`
	ContinmeFee    decimal.Decimal `gorm:"column:continme_fee;type:decimal(10,2);not null;default:0.00;comment:续费（元）"`
	Dest           string          `gorm:"column:dest;type:varchar(255);not null;default:'';comment:目的地（省、市）"`
	// 公共字段
	BaseTime
}

func (f FeightTemplate) TableName() string {
	return "pms_feight_template"
}
