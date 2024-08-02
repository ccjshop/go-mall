package entity

import "github.com/shopspring/decimal"

// SkuStock 商品SKU表
// SKU（Stock Keeping Unit）是指库存量单位，SPU（Standard Product Unit）是指标准产品单位。举个例子：iphone xs是一个SPU，而iphone xs 公开版 64G 银色是一个SKU。
type SkuStock struct {
	ID        uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;comment:主键"`
	SkuCode   string `gorm:"column:sku_code;type:varchar(64);not null;default:'';comment:sku编码"`
	Pic       string `gorm:"column:pic;type:varchar(255);not null;default:'';comment:展示图片"`
	Sale      uint32 `gorm:"column:sale;type:int(10);unsigned;not null;default:0;comment:销量"`
	SpData    string `gorm:"column:sp_data;type:varchar(500);not null;default:'';comment:商品销售属性，json格式"` // [{"key":"颜色","value":"黑色"},{"key":"容量","value":"32G"}]
	ProductID uint64 `gorm:"column:product_id;type:bigint;unsigned;not null;default:0;comment:产品ID"`     // pms_product#id
	// 价格
	Price          decimal.Decimal `gorm:"column:price;type:decimal(10,2);not null;default:0.00;comment:价格"`
	PromotionPrice decimal.Decimal `gorm:"column:promotion_price;type:decimal(10,2);not null;default:0.00;comment:单品促销价格"`
	// 库存
	Stock     uint32 `gorm:"column:stock;type:int(10);unsigned;not null;default:0;comment:库存"`
	LockStock uint32 `gorm:"column:lock_stock;type:int(10);unsigned;not null;default:0;comment:锁定库存"`
	LowStock  uint32 `gorm:"column:low_stock;type:int(10);unsigned;not null;default:0;comment:预警库存"`
	// 公共字段
	BaseTime
}

func (c SkuStock) TableName() string {
	return "pms_sku_stock"
}

// SkuStocks 商品SKU集合
type SkuStocks []*SkuStock
