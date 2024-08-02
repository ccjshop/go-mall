package entity

// HomeRecommendProduct 人气推荐商品表
// 用于管理首页显示的人气推荐信息。
type HomeRecommendProduct struct {
	ID              uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;"`
	ProductID       uint64 `gorm:"column:product_id;type:bigint;unsigned;not null;default:0"`
	ProductName     string `gorm:"column:product_name;type:varchar(500);not null;default:''"`
	RecommendStatus uint8  `gorm:"column:recommend_status;type:tinyint(4);unsigned;not null;default:0"`
	Sort            uint32 `gorm:"column:sort;type:int(10);unsigned;not null;default:0"`
	// 公共字段
	BaseTime
}

func (p HomeRecommendProduct) TableName() string {
	return "sms_home_recommend_product"
}
