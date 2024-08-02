package entity

// FlashPromotion 限时购表
// 用于存储限时购活动的信息，包括开始时间、结束时间以及上下线状态。
type FlashPromotion struct {
	ID         uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment"`
	Title      string `gorm:"column:title;type:varchar(200);not null;default:'';comment:标题"`
	Status     uint8  `gorm:"column:status;type:tinyint(4);unsigned;not null;default:0;comment:上下线状态"`
	StartDate  uint32 `gorm:"column:start_date;type:int(10);unsigned;not null;default:0;comment:开始日期"`
	EndDate    uint32 `gorm:"column:end_date;type:int(10);unsigned;not null;default:0;comment:结束日期"`
	CreateTime uint32 `gorm:"column:create_time;type:int(10);unsigned;not null;default:0;comment:创建时间"`
	// 公共字段
	BaseTime
}

func (s FlashPromotion) TableName() string {
	return "sms_flash_promotion"
}
