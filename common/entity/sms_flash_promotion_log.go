package entity

// FlashPromotionLog 限时购通知记录
type FlashPromotionLog struct {
	ID            uint32 `gorm:"column:id;type:int(10);primary_key;auto_increment"`
	MemberID      uint32 `gorm:"column:member_id;type:int(10);unsigned;not null;default:0;comment:会员id"`
	ProductID     uint64 `gorm:"column:product_id;type:bigint;unsigned;not null;default:0;comment:商品id"`
	MemberPhone   string `gorm:"column:member_phone;type:varchar(64);not null;default:'';comment:会员电话"`
	ProductName   string `gorm:"column:product_name;type:varchar(100);not null;default:'';comment:商品名称"`
	SubscribeTime uint32 `gorm:"column:subscribe_time;type:int(10);unsigned;not null;default:0;comment:会员订阅时间"`
	SendTime      uint32 `gorm:"column:send_time;type:int(10);unsigned;not null;default:0;comment:发送时间"`
	// 公共字段
	BaseTime
}

func (s FlashPromotionLog) TableName() string {
	return "sms_flash_promotion_log"
}
