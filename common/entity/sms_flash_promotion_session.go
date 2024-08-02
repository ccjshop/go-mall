package entity

// FlashPromotionSession 限时购场次表
// 用于存储限时购场次的信息，在一天中，一个限时购活动会有多个不同的活动时间段。
type FlashPromotionSession struct {
	ID         uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;comment:编号"`
	Name       string `gorm:"column:name;type:varchar(200);not null;default:'';comment:场次名称"`
	Status     uint8  `gorm:"column:status;type:tinyint(4);unsigned;not null;default:0;comment:启用状态：0->不启用；1->启用"`
	StartTime  uint32 `gorm:"column:start_time;type:int(10);unsigned;not null;default:0;comment:每日开始时间"`
	EndTime    uint32 `gorm:"column:end_time;type:int(10);unsigned;not null;default:0;comment:每日结束时间"`
	CreateTime uint32 `gorm:"column:create_time;type:int(10);unsigned;not null;default:0;comment:创建时间"`
	// 公共字段
	BaseTime
}

func (s FlashPromotionSession) TableName() string {
	return "sms_flash_promotion_session"
}
