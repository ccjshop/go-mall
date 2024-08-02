package entity

import "github.com/shopspring/decimal"

// OrderItem 订单商品信息表
// 订单中包含的商品信息，一个订单中会有多个订单商品信息。
type OrderItem struct {
	ID      uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment"`
	OrderID uint64 `gorm:"column:order_id;type:bigint;unsigned;not null;default:0;comment:订单id"` // oms_order#id
	OrderSN string `gorm:"column:order_sn;type:varchar(64);not null;default:'';comment:订单编号"`
	// 商品信息
	ProductID         uint64 `gorm:"column:product_id;type:bigint;unsigned;not null;default:0;comment:商品id"`
	ProductCategoryID uint64 `gorm:"column:product_category_id;type:bigint;unsigned;not null;default:0;comment:商品分类id"`
	ProductPic        string `gorm:"column:product_pic;type:varchar(500);not null;default:'';comment:商品图片"`
	ProductName       string `gorm:"column:product_name;type:varchar(200);not null;default:'';comment:商品名称"`
	ProductBrand      string `gorm:"column:product_brand;type:varchar(200);not null;default:'';comment:商品品牌"`
	ProductSN         string `gorm:"column:product_sn;type:varchar(64);not null;default:'';comment:商品条码"`
	ProductAttr       string `gorm:"column:product_attr;type:varchar(500);not null;default:'';comment:商品销售属性:[{\"key\":\"颜色\",\"value\":\"颜色\"},{\"key\":\"容量\",\"value\":\"4G\"}]"`
	PromotionName     string `gorm:"column:promotion_name;type:varchar(200);not null;default:'';comment:商品促销名称"`
	// sku
	ProductSkuID   uint64 `gorm:"column:product_sku_id;type:bigint;unsigned;not null;default:0;comment:商品sku编号"`
	ProductSkuCode string `gorm:"column:product_sku_code;type:varchar(50);not null;default:'';comment:商品sku条码"`
	// 数量
	ProductQuantity uint32 `gorm:"column:product_quantity;type:int(10);unsigned;not null;default:0;comment:购买数量"`
	// 价格&价格优惠&最终价格 realAmount = productPrice - promotionAmount - couponAmount - integrationAmount
	ProductPrice      decimal.Decimal `gorm:"column:product_price;type:decimal(10,2);not null;default:0.00;comment:销售价格"`
	PromotionAmount   decimal.Decimal `gorm:"column:promotion_amount;type:decimal(10,2);not null;default:0.00;comment:商品促销分解金额"`
	CouponAmount      decimal.Decimal `gorm:"column:coupon_amount;type:decimal(10,2);not null;default:0.00;comment:优惠券优惠分解金额"`
	IntegrationAmount decimal.Decimal `gorm:"column:integration_amount;type:decimal(10,2);not null;default:0.00;comment:积分优惠分解金额"`
	RealAmount        decimal.Decimal `gorm:"column:real_amount;type:decimal(10,2);not null;default:0.00;comment:该商品经过优惠后的分解金额"`
	// 其他奖励
	GiftIntegration uint32 `gorm:"column:gift_integration;type:int(10);unsigned;not null;default:0;comment:商品赠送积分"`
	GiftGrowth      uint32 `gorm:"column:gift_growth;type:int(10);unsigned;not null;default:0;comment:商品赠送成长值"`
	// 公共字段
	BaseTime
}

func (o OrderItem) TableName() string {
	return "oms_order_item"
}

type OrderItems []*OrderItem
