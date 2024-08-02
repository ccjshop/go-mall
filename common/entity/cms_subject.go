package entity

// Subject 专题表
type Subject struct {
	ID              uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;comment:主键"`
	CategoryID      uint64 `gorm:"column:category_id;type:bigint;unsigned;not null;default:0;comment:分类id"` //
	Title           string `gorm:"column:title;type:varchar(100);not null;default:'';comment:标题"`
	Pic             string `gorm:"column:pic;type:varchar(500);not null;default:'';comment:专题主图"`
	AlbumPics       string `gorm:"column:album_pics;type:varchar(1000);not null;default:'';comment:画册图片，用逗号分割"`
	Description     string `gorm:"column:description;type:varchar(1000);not null;default:'';comment:描述"`
	Content         string `gorm:"column:content;type:text;not null;comment:内容"`
	ShowStatus      uint8  `gorm:"column:show_status;type:tinyint(4);unsigned;not null;default:0;comment:显示状态：0->不显示；1->显示"`
	RecommendStatus uint8  `gorm:"column:recommend_status;type:tinyint(4);unsigned;not null;default:0;comment:推荐状态"`
	CreateTime      uint32 `gorm:"column:create_time;type:int(10);unsigned;not null;default:0;comment:创建时间"`
	// 冗余字段
	CategoryName string `gorm:"column:category_name;type:varchar(200);not null;default:'';comment:专题分类名称"`
	ForwardCount uint32 `gorm:"column:forward_count;type:int(10);unsigned;not null;default:0;comment:转发数"`
	CollectCount uint32 `gorm:"column:collect_count;type:int(10);unsigned;not null;default:0;comment:收藏数量"`
	ReadCount    uint32 `gorm:"column:read_count;type:int(10);unsigned;not null;default:0;comment:阅读数量"`
	CommentCount uint32 `gorm:"column:comment_count;type:int(10);unsigned;not null;default:0;comment:评论数量"`
	ProductCount uint32 `gorm:"column:product_count;type:int(10);unsigned;not null;default:0;comment:关联产品数量"`
	// 公共字段
	BaseTime
}

func (c Subject) TableName() string {
	return "cms_subject"
}
