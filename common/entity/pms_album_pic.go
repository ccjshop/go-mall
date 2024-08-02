package entity

// AlbumPic 画册图片表
type AlbumPic struct {
	ID      uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;comment:主键"`
	AlbumID uint64 `gorm:"column:album_id;type:bigint;unsigned;not null;default:0;comment:相册id"`
	Pic     string `gorm:"column:pic;type:varchar(1000);not null;default:'';comment:图片"`
	// 公共字段
	BaseTime
}

func (a AlbumPic) TableName() string {
	return "pms_album_pic"
}
