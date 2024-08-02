package entity

// PrefrenceAreaProductRelation 优选专区和产品关系表
type PrefrenceAreaProductRelation struct {
	ID              uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;comment:主键"`
	PrefrenceAreaID uint64 `gorm:"column:prefrence_area_id;type:bigint;unsigned;not null;default:0"`
	ProductID       uint64 `gorm:"column:product_id;type:bigint;unsigned;not null;default:0"`
}

func (PrefrenceAreaProductRelation) TableName() string {
	return "cms_prefrence_area_product_relation"
}
