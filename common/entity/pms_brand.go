package entity

// Brand 商品品牌表
type Brand struct {
	ID          uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;comment:主键"`
	Name        string `gorm:"column:name;type:varchar(64);not null;default:'';comment:名称"`
	FirstLetter string `gorm:"column:first_letter;type:varchar(8);not null;default:'';comment:首字母"`
	Logo        string `gorm:"column:logo;type:varchar(255);not null;default:'';comment:品牌logo"`
	BigPic      string `gorm:"column:big_pic;type:varchar(255);not null;default:'';comment:专区大图"`
	BrandStory  string `gorm:"column:brand_story;type:text;not null;comment:品牌故事"`
	Sort        uint32 `gorm:"column:sort;type:int(10);unsigned;not null;default:0;comment:排序"`
	// status
	FactoryStatus uint8 `gorm:"column:factory_status;type:tinyint(4);unsigned;not null;default:0;comment:是否为品牌制造商：0->不是；1->是"`
	ShowStatus    uint8 `gorm:"column:show_status;type:tinyint(4);unsigned;not null;default:0;comment:是否显示"`
	// 冗余字段
	ProductCount        uint32 `gorm:"column:product_count;type:int(10);unsigned;not null;default:0;comment:产品数量"`
	ProductCommentCount uint32 `gorm:"column:product_comment_count;type:int(10);unsigned;not null;default:0;comment:产品评论数量"`
	// 公共字段
	BaseTime
}

func (c Brand) TableName() string {
	return "pms_brand"
}

// Brands 商品商品品牌集合
type Brands []*Brand

// NameMap 获取商品品牌名 key=id value=name
func (b Brands) NameMap() map[uint64]string {
	res := make(map[uint64]string)
	for _, brand := range b {
		res[brand.ID] = brand.Name
	}
	return res
}
