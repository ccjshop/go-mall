package entity

// PrefrenceArea 优选专区
type PrefrenceArea struct {
	ID         uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;comment:主键"`
	Name       string `gorm:"column:name;type:varchar(255);not null;default:'';comment:名称"`
	SubTitle   string `gorm:"column:sub_title;type:varchar(255);not null;default:'';comment:子标题"`
	Pic        string `gorm:"column:pic;type:varchar(255);not null;default:'';comment:展示图片"`
	Sort       uint32 `gorm:"column:sort;type:int(10);unsigned;not null;default:0;comment:排序"`
	ShowStatus uint8  `gorm:"column:show_status;type:tinyint(4);unsigned;not null;default:0;comment:显示状态"`
	// 公共字段
	BaseTime
}

func (c PrefrenceArea) TableName() string {
	return "cms_prefrence_area"
}
