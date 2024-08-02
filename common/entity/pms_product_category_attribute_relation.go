package entity

// ProductCategoryAttributeRelation 商品分类和属性的关系表
// 用于前台商城选中分类后搜索时生成筛选属性（暂未使用）。
type ProductCategoryAttributeRelation struct {
	ID                 uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;comment:主键"`                           //
	ProductCategoryID  uint64 `gorm:"column:product_category_id;type:bigint;unsigned;not null;default:0;comment:产品分类表ID"`    // pms_product_category#id
	ProductAttributeID uint64 `gorm:"column:product_attribute_id;type:bigint;unsigned;not null;default:0;comment:商品属性参数表ID"` // pms_product_attribute#id type=1
	// 公共字段
	BaseTime
}

func (c ProductCategoryAttributeRelation) TableName() string {
	return "pms_product_category_attribute_relation"
}
