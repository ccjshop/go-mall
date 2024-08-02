package entity

// ProductAttributeCategory 产品属性分类表
// 商品属性分类表，便于管理商品属性。
type ProductAttributeCategory struct {
	ID   uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;comment:编号"`
	Name string `gorm:"column:name;type:varchar(64);not null;default:'';comment:类型名称"`
	// 冗余字段
	AttributeCount int `gorm:"column:attribute_count;type:int(10);unsigned;not null;default:0;comment:属性数量"` // pms_product_attribute type=0
	ParamCount     int `gorm:"column:param_count;type:int(10);unsigned;not null;default:0;comment:参数数量"`     // pms_product_attribute type=1
	// 公共字段
	BaseTime
}

func (c ProductAttributeCategory) TableName() string {
	return "pms_product_attribute_category"
}
