package entity

import "github.com/ccjshop/go-mall/common/util"

// CouponProductCategoryRelation 优惠券和商品分类关系表
// 用于存储优惠券与商品分类的关系，当优惠券的使用类型为指定分类时，优惠券与商品分类需要建立关系。
type CouponProductCategoryRelation struct {
	ID                  uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;"`
	CouponID            uint64 `gorm:"column:coupon_id;type:bigint;unsigned;not null;default:0;comment:优惠券id"`
	ProductCategoryID   uint64 `gorm:"column:product_category_id;type:bigint;unsigned;not null;default:0;comment:商品分类id"`
	ProductCategoryName string `gorm:"column:product_category_name;type:varchar(200);not null;default:'';comment:商品分类名称"`
	ParentCategoryName  string `gorm:"column:parent_category_name;type:varchar(200);not null;default:'';comment:父分类名称"`
	// 公共字段
	BaseTime
}

func (r CouponProductCategoryRelation) TableName() string {
	return "sms_coupon_product_category_relation"
}

type CouponProductCategoryRelations []*CouponProductCategoryRelation

// GetProductCategoryIDs 获取商品分类id
func (r CouponProductCategoryRelations) GetProductCategoryIDs() []uint64 {
	//productCategoryIDs := make([]uint64, 0)
	//for _, categoryRelation := range r {
	//	productCategoryIDs = append(productCategoryIDs, categoryRelation.ProductCategoryID)
	//}
	//return productCategoryIDs

	extractor := util.NewFieldExtractor[*CouponProductCategoryRelation, uint64]()
	return extractor.ExtractField(r, func(relation *CouponProductCategoryRelation) uint64 {
		return relation.ProductCategoryID
	})
}
