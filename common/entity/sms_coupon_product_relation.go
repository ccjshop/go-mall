package entity

// CouponProductRelation 优惠券和商品的关系表
// 用于存储优惠券与商品的关系，当优惠券的使用类型为指定商品时，优惠券与商品需要建立关系。
type CouponProductRelation struct {
	ID          uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;"`
	CouponID    uint64 `gorm:"column:coupon_id;type:bigint;unsigned;not null;default:0;comment:优惠券id"`
	ProductID   uint64 `gorm:"column:product_id;type:bigint;unsigned;not null;default:0;comment:商品id"`
	ProductName string `gorm:"column:product_name;type:varchar(500);not null;default:'';comment:商品名称"`
	ProductSN   string `gorm:"column:product_sn;type:varchar(200);not null;default:'';comment:商品条码"`
	// 公共字段
	BaseTime
}

func (r CouponProductRelation) TableName() string {
	return "sms_coupon_product_relation"
}

type CouponProductRelations []*CouponProductRelation

func (r CouponProductRelations) GetProductIDs() []uint64 {
	productIDs := make([]uint64, 0)
	for _, productRelation := range r {
		productIDs = append(productIDs, productRelation.ProductID)
	}

	return productIDs
}
