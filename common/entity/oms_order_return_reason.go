package entity

// OrderReturnReason 退货原因表
type OrderReturnReason struct {
	ID         uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;comment:主键"`
	Name       string `gorm:"column:name;type:varchar(100);not null;default:'';comment:退货类型"`
	Sort       uint32 `gorm:"column:sort;type:int(10);unsigned;not null;default:0;comment:排序"`
	Status     uint8  `gorm:"column:status;type:tinyint(4);unsigned;not null;default:0;comment:状态：0->不启用；1->启用"`
	CreateTime uint32 `gorm:"column:create_time;type:int(10);unsigned;not null;default:0;comment:添加时间"`
	// 公共字段
	BaseTime
}

func (o OrderReturnReason) TableName() string {
	return "oms_order_return_reason"
}
