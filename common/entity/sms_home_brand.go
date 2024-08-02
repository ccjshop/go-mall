package entity

// HomeBrand 首页品牌推荐表
// 用于管理首页显示的品牌制造商直供信息。
type HomeBrand struct {
	ID              uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;"`
	BrandID         uint64 `gorm:"column:brand_id;type:bigint;unsigned;not null;default:0;comment:商品品牌id"`
	BrandName       string `gorm:"column:brand_name;type:varchar(64);not null;default:'';comment:商品品牌名称"`
	RecommendStatus uint8  `gorm:"column:recommend_status;type:tinyint(4);unsigned;not null;default:0;comment:推荐状态：0->不推荐;1->推荐"`
	Sort            uint32 `gorm:"column:sort;type:int(10);unsigned;not null;default:0;comment:排序"`
	// 公共字段
	BaseTime
}

func (b HomeBrand) TableName() string {
	return "sms_home_brand"
}
