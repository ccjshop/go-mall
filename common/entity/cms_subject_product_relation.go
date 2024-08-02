package entity

// SubjectProductRelation 专题商品关系表
type SubjectProductRelation struct {
	ID        uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;comment:主键"`
	SubjectID uint64 `gorm:"column:subject_id;type:bigint;unsigned;not null;default:0"` // cms_subject#id
	ProductID uint64 `gorm:"column:product_id;type:bigint;unsigned;not null;default:0"` // pms_product#id
}

func (SubjectProductRelation) TableName() string {
	return "cms_subject_product_relation"
}
