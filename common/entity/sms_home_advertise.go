package entity

// HomeAdvertise 首页轮播广告表
// 用于管理首页显示的轮播广告信息。
type HomeAdvertise struct {
	ID   uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;"`
	Name string `gorm:"column:name;type:varchar(100);not null;default:'';comment:名称"`
	Pic  string `gorm:"column:pic;type:varchar(500);not null;default:'';comment:图片地址"`
	URL  string `gorm:"column:url;type:varchar(500);not null;default:'';comment:链接地址"`
	Sort uint32 `gorm:"column:sort;type:int(10);unsigned;not null;default:0;comment:排序"`
	Note string `gorm:"column:note;type:varchar(500);not null;default:'';comment:备注"`
	// 类型
	Type uint8 `gorm:"column:type;type:tinyint(4);unsigned;not null;default:0;comment:轮播位置：0->PC首页轮播；1->app首页轮播"`
	// 时间
	StartTime uint32 `gorm:"column:start_time;type:int(10);unsigned;not null;default:0;comment:开始时间"`
	EndTime   uint32 `gorm:"column:end_time;type:int(10);unsigned;not null;default:0;comment:结束时间"`
	// 状态
	Status uint8 `gorm:"column:status;type:tinyint(4);unsigned;not null;default:0;comment:上下线状态：0->下线；1->上线"`
	// 统计
	ClickCount uint32 `gorm:"column:click_count;type:int(10);unsigned;not null;default:0;comment:点击数"`
	OrderCount uint32 `gorm:"column:order_count;type:int(10);unsigned;not null;default:0;comment:下单数"`
	// 公共字段
	BaseTime
}

func (a HomeAdvertise) TableName() string {
	return "sms_home_advertise"
}
