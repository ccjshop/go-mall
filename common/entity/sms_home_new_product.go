package entity

// HomeNewProduct 新品推荐商品表
// 用于管理首页显示的新鲜好物信息。
type HomeNewProduct struct {
	ID              uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;"`
	ProductID       uint64 `gorm:"column:product_id;type:bigint;unsigned;not null;default:0;comment:商品id"`
	ProductName     string `gorm:"column:product_name;type:varchar(500);not null;default:'';comment:商品名称"`
	RecommendStatus uint8  `gorm:"column:recommend_status;type:tinyint(4);unsigned;not null;default:0;comment:推荐状态：0->不推荐;1->推荐"`
	Sort            uint32 `gorm:"column:sort;type:int(10);unsigned;not null;default:0;comment:排序"`
	// 公共字段
	BaseTime
}

func (p HomeNewProduct) TableName() string {
	return "sms_home_new_product"
}
