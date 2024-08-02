package entity

// OrderOperateHistory 订单操作历史记录表
// 当订单状态发生改变时，用于记录订单的操作信息。
type OrderOperateHistory struct {
	ID          uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment"`
	OrderID     uint64 `gorm:"column:order_id;type:bigint;unsigned;not null;default:0;comment:订单id"` // oms_order#id
	OperateMan  string `gorm:"column:operate_man;type:varchar(100);not null;default:'';comment:操作人：用户；系统；后台管理员"`
	OrderStatus uint8  `gorm:"column:order_status;type:tinyint(4);unsigned;not null;default:0;comment:订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单"`
	Note        string `gorm:"column:note;type:varchar(500);not null;default:'';comment:备注"`
	CreateTime  uint32 `gorm:"column:create_time;type:int(10);unsigned;not null;default:0;comment:操作时间"`
	// 公共字段
	BaseTime
}

func (o OrderOperateHistory) TableName() string {
	return "oms_order_operate_history"
}

type OrderOperateHistories []*OrderOperateHistory
