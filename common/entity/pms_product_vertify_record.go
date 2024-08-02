package entity

// ProductVertifyRecord 商品审核记录表
// 用于记录商品审核记录
type ProductVertifyRecord struct {
	ID         uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;comment:主键"`
	ProductID  uint64 `gorm:"column:product_id;type:bigint;not null;default:0;comment:商品id"`
	CreateTime uint32 `gorm:"column:create_time;type:int(10);unsigned;not null;default:0;comment:创建时间"`
	VertifyMan string `gorm:"column:vertify_man;type:varchar(64);not null;default:'';comment:审核人"`
	Status     uint32 `gorm:"column:status;type:int(1);unsigned;not null;default:0;comment:审核后的状态：0->未通过；2->已通过"`
	Detail     string `gorm:"column:detail;type:varchar(255);not null;default:'';comment:反馈详情"`
	// 公共字段
	BaseTime
}

func (c ProductVertifyRecord) TableName() string {
	return "pms_product_vertify_record"
}
