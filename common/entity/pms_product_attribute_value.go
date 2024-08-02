package entity

// ProductAttributeValue 商品属性值表
// 1、如果对应的商品属性是规格且规格支持手动添加(pms_product_attribute#type=0#input_type=0)，那么该表用于存储手动新增的值；
// 2、如果对应的商品属性是参数(pms_product_attribute#type=1)，那么该表用于存储参数的值。
type ProductAttributeValue struct {
	ID                 uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;comment:主键ID"`
	ProductID          uint64 `gorm:"column:product_id;type:bigint;unsigned;not null;default:0;comment:商品ID"`             // pms_product#id
	ProductAttributeID uint64 `gorm:"column:product_attribute_id;type:bigint;unsigned;not null;default:0;comment:商品属性ID"` // pms_product_attribute#id
	// 值
	Value string `gorm:"column:value;type:varchar(64);not null;default:'';comment:手动添加规格或参数的值，参数单值，规格有多个时以逗号隔开"`
	// 公共字段
	BaseTime
}

// TableName sets the insert table name for this struct type
func (ProductAttributeValue) TableName() string {
	return "pms_product_attribute_value"
}

type ProductAttributeValues []*ProductAttributeValue
