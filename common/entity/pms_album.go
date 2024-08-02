package entity

// Album 相册表
type Album struct {
	ID          uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;comment:主键"`
	Name        string `gorm:"column:name;type:varchar(64);not null;default:'';comment:相册名称"`
	CoverPic    string `gorm:"column:cover_pic;type:varchar(1000);not null;default:'';comment:相册封面"`
	PicCount    uint32 `gorm:"column:pic_count;type:int(10);unsigned;not null;default:0;comment:图片数量"`
	Sort        uint32 `gorm:"column:sort;type:int(10);unsigned;not null;default:0;comment:排序"`
	Description string `gorm:"column:description;type:varchar(1000);not null;default:'';comment:描述"`
	// 公共字段
	BaseTime
}

func (a Album) TableName() string {
	return "pms_album"
}
