package entity

import (
	pb "github.com/ccjshop/go-mall/proto/mall"
	"github.com/shopspring/decimal"
)

// Product 商品信息表
// 商品信息主要包括四部分：商品的基本信息、商品的促销信息、商品的属性信息、商品的关联，商品表是整个商品的基本信息部分。
type Product struct {
	// 基本信息
	ID                uint64          `gorm:"column:id;type:bigint;primary_key;auto_increment;comment:主键"`
	ProductCategoryID uint64          `gorm:"column:product_category_id;type:bigint;unsigned;not null;default:0;comment:商品分类id"` // pms_product_category#id
	Name              string          `gorm:"column:name;type:varchar(200);not null;default:'';comment:商品名称"`
	SubTitle          string          `gorm:"column:sub_title;type:varchar(255);not null;default:'';comment:副标题"`
	BrandID           uint64          `gorm:"column:brand_id;type:bigint;unsigned;not null;default:0;comment:品牌id"` // pms_brand#id
	Description       string          `gorm:"column:description;type:text;not null;comment:商品描述"`
	ProductSN         string          `gorm:"column:product_sn;type:varchar(64);not null;default:'';comment:货号"`
	Unit              string          `gorm:"column:unit;type:varchar(16);not null;default:'';comment:单位"`
	Weight            decimal.Decimal `gorm:"column:weight;type:decimal(10,2);not null;default:0.00;comment:商品重量，默认为克"`
	Sort              uint32          `gorm:"column:sort;type:int(10);unsigned;not null;default:0;comment:排序"`

	// 价格-库存 真实的价格库存取的是pms_sku_stock的
	Price         decimal.Decimal `gorm:"column:price;type:decimal(10,2);not null;default:0.00;comment:价格"`
	OriginalPrice decimal.Decimal `gorm:"column:original_price;type:decimal(10,2);not null;default:0.00;comment:市场价"`
	Stock         uint32          `gorm:"column:stock;type:int(10);unsigned;not null;default:0;comment:库存"`
	LowStock      uint32          `gorm:"column:low_stock;type:int(10);unsigned;not null;default:0;comment:库存预警值"`

	// 促销信息
	GiftPoint          uint32           `gorm:"column:gift_point;type:int(10);unsigned;not null;default:0;comment:赠送的积分"`
	GiftGrowth         uint32           `gorm:"column:gift_growth;type:int(10);unsigned;not null;default:0;comment:赠送的成长值"`
	UsePointLimit      uint32           `gorm:"column:use_point_limit;type:int(10);unsigned;not null;default:0;comment:限制使用的积分数"`
	PreviewStatus      uint8            `gorm:"column:preview_status;type:tinyint(4);unsigned;not null;default:0;comment:是否为预告商品：0->不是；1->是"`
	PublishStatus      uint8            `gorm:"column:publish_status;type:tinyint(4);unsigned;not null;default:0;comment:上架状态：0->下架；1->上架"`
	NewStatus          uint8            `gorm:"column:new_status;type:tinyint(4);unsigned;not null;default:0;comment:新品状态:0->不是新品；1->新品"`
	RecommandStatus    uint8            `gorm:"column:recommand_status;type:tinyint(4);unsigned;not null;default:0;comment:推荐状态；0->不推荐；1->推荐"`
	ServiceIDs         string           `gorm:"column:service_ids;type:varchar(64);not null;default:'';comment:以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮"`
	DetailTitle        string           `gorm:"column:detail_title;type:varchar(255);not null;default:'';comment:详情标题"`
	DetailDesc         string           `gorm:"column:detail_desc;type:text;not null;comment:详情描述"`
	Keywords           string           `gorm:"column:keywords;type:varchar(255);not null;default:'';comment:关键字"`
	Note               string           `gorm:"column:note;type:varchar(255);not null;default:'';comment:备注"`
	PromotionType      pb.PromotionType `gorm:"column:promotion_type;type:tinyint(4);unsigned;not null;default:0;comment:促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购"`
	PromotionPrice     decimal.Decimal  `gorm:"column:promotion_price;type:decimal(10,2);not null;default:0.00;comment:促销价格"`
	PromotionStartTime uint32           `gorm:"column:promotion_start_time;type:int(10);unsigned;not null;default:0;comment:促销开始时间"`
	PromotionEndTime   uint32           `gorm:"column:promotion_end_time;type:int(10);unsigned;not null;default:0;comment:促销结束时间"`

	// 属性信息
	ProductAttributeCategoryID uint64   `gorm:"column:product_attribute_category_id;type:bigint;unsigned;not null;default:0;comment:品牌属性分类id"`
	Pic                        string   `gorm:"column:pic;type:varchar(255);not null;default:'';comment:图片"`
	AlbumPics                  []string `gorm:"column:album_pics;type:varchar(255);not null;default:'';serializer:json;comment:画册图片，连产品图片限制为5张，以逗号分割"`
	DetailHTML                 string   `gorm:"column:detail_html;type:text;not null;comment:电脑端详情"`
	DetailMobileHTML           string   `gorm:"column:detail_mobile_html;type:text;not null;comment:移动端详情"`

	// 状态
	VerifyStatus uint8 `gorm:"column:verify_status;type:tinyint(4);unsigned;not null;default:0;comment:审核状态：0->未审核；1->审核通过"`
	DeleteStatus uint8 `gorm:"column:delete_status;type:tinyint(4);unsigned;not null;default:0;comment:删除状态：0->未删除；1->已删除"`

	// 其他
	FeightTemplateID  uint64 `gorm:"column:feight_template_id;type:bigint;unsigned;not null;default:0;comment:运费模版id"`
	Sale              uint32 `gorm:"column:sale;type:int(10);unsigned;not null;default:0;comment:销量"`
	PromotionPerLimit uint32 `gorm:"column:promotion_per_limit;type:int(10);unsigned;not null;default:0;comment:活动限购数量"`

	// 冗余字段
	BrandName           string `gorm:"column:brand_name;type:varchar(255);comment:品牌名称"`
	ProductCategoryName string `gorm:"column:product_category_name;type:varchar(255);comment:商品分类名称"`

	// 公共字段
	BaseTime
}

func (c Product) TableName() string {
	return "pms_product"
}

// Products 商品信息集合
type Products []*Product

// CategoryIDs 获取商品品牌ID
func (p Products) CategoryIDs() []uint64 {
	ids := make([]uint64, len(p))
	for i, product := range p {
		ids[i] = product.ProductCategoryID
	}
	return ids
}

// BrandIDs 获取商品分类ID
func (p Products) BrandIDs() []uint64 {
	ids := make([]uint64, len(p))
	for i, product := range p {
		ids[i] = product.BrandID
	}
	return ids
}
