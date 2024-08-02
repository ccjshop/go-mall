package entity

import (
	"github.com/shopspring/decimal"
)

// CartItem 购物车表
type CartItem struct {
	ID uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;comment:主键"`
	// 用户信息
	MemberID uint64 `gorm:"column:member_id;type:bigint;unsigned;not null;default:0;comment:会员id"` // ums_member#id
	// 商品信息
	ProductID         uint64 `gorm:"column:product_id;type:bigint;unsigned;not null;default:0;comment:商品id"`           // pms_product#id
	ProductName       string `gorm:"column:product_name;type:varchar(500);not null;default:'';comment:商品名称"`           // pms_product#name
	ProductPic        string `gorm:"column:product_pic;type:varchar(1000);not null;default:'';comment:商品主图"`           // pms_product#pic
	ProductSubTitle   string `gorm:"column:product_sub_title;type:varchar(500);not null;default:'';comment:商品副标题（卖点）"` // pms_product#sub_title
	ProductSN         string `gorm:"column:product_sn;type:varchar(200);not null;default:'';comment:商品货号"`             // pms_product#product_sn
	ProductBrand      string `gorm:"column:product_brand;type:varchar(200);not null;default:'';comment:商品品牌"`          // pms_brand#id
	ProductCategoryID uint64 `gorm:"column:product_category_id;type:bigint;unsigned;not null;default:0;comment:商品分类"`  // pms_product_category#id
	// 商品sku
	ProductSkuID   uint64 `gorm:"column:product_sku_id;type:bigint;unsigned;not null;default:0;comment:商品sku id"` // pms_sku_stock#id
	ProductSkuCode string `gorm:"column:product_sku_code;type:varchar(200);not null;default:'';comment:商品sku条码"`  // pms_sku_stock#sku_code
	ProductAttr    string `gorm:"column:product_attr;type:varchar(500);not null;default:'';comment:商品销售属性"`       // [{'key':'颜色','value':'颜色'},{'key':'容量','value':'4G'}]
	// 价格数量
	Price    decimal.Decimal `gorm:"column:price;type:decimal(10,2);not null;default:0.00;comment:添加到购物车的价格"`
	Quantity uint32          `gorm:"column:quantity;type:int(10);unsigned;not null;default:0;comment:购买数量"`
	// 状态
	CreateDate   uint32 `gorm:"column:create_date;type:int(10);unsigned;not null;default:0;comment:创建时间"`
	ModifyDate   uint32 `gorm:"column:modify_date;type:int(10);unsigned;not null;default:0;comment:修改时间"`
	DeleteStatus uint8  `gorm:"column:delete_status;type:tinyint(4);unsigned;not null;default:0;comment:是否删除"`
	// 冗余字段
	MemberNickname string `gorm:"column:member_nickname;type:varchar(500);not null;default:'';comment:会员昵称"`
	// 公共字段
	BaseTime
}

func (c CartItem) TableName() string {
	return "oms_cart_item"
}

type CartItems []*CartItem

// GetMemberIDs 获取会员ID集合
func (c CartItems) GetMemberIDs() []uint64 {
	res := make([]uint64, 0)
	for _, item := range c {
		res = append(res, item.MemberID)
	}
	return res
}

// GetProductIDs 获取商品ID集合
func (c CartItems) GetProductIDs() []uint64 {
	res := make([]uint64, 0)
	for _, item := range c {
		res = append(res, item.ProductID)
	}
	return res
}

// GroupCartItemBySpu 以spu为单位对购物车中商品进行分组 key=商品id value=购物车集合
func (c CartItems) GroupCartItemBySpu() map[uint64]CartItems {
	productCartMap := make(map[uint64]CartItems)
	for _, cartItem := range c {
		productID := cartItem.ProductID
		productCartMap[productID] = append(productCartMap[productID], cartItem)
	}
	return productCartMap
}

// GetCartItemCount 获取购物车中指定商品的数量
func (c CartItems) GetCartItemCount() uint32 {
	count := uint32(0)
	for _, item := range c {
		count += item.Quantity
	}
	return count
}
