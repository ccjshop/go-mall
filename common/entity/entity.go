package entity

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// BaseTime 基础的时间字段，每个dao结构体都应该嵌套这个字段
type BaseTime struct {
	CreatedAt uint32 `gorm:"type:int(10);unsigned;not null;default:0;comment:创建时间"` // 使用时间戳秒数填充创建时间
	UpdatedAt uint32 `gorm:"type:int(10);unsigned;not null;default:0;comment:修改时间"` // 使用时间戳秒数填充修改时间
}

// BeforeCreate gorm创建之前回调
func (bt *BaseTime) BeforeCreate(tx *gorm.DB) (err error) {
	timestamp := uint32(time.Now().Unix())
	bt.CreatedAt = timestamp
	bt.UpdatedAt = timestamp
	return
}

// BeforeUpdate gorm修改之前回调
func (bt *BaseTime) BeforeUpdate(tx *gorm.DB) (err error) {
	bt.UpdatedAt = uint32(time.Now().Unix())
	return
}

// BigDecimal 是一个包装了 shopspring/decimal.Decimal 的自定义类型
type BigDecimal struct {
	decimal.Decimal
}

// Scan 实现了 sql.Scanner 接口
func (bd *BigDecimal) Scan(value interface{}) error {
	return bd.Decimal.Scan(value)
}

// Value 实现了 driver.Valuer 接口
func (bd BigDecimal) Value() (driver.Value, error) {
	return bd.Decimal.Value()
}

// Add returns d + d2.
func (bd BigDecimal) Add(d2 BigDecimal) BigDecimal {
	return BigDecimal{Decimal: bd.Decimal.Add(d2.Decimal)}
}

// Init 初始化
func Init(db *gorm.DB) error {
	pmsSchemas := []tableSchema{
		{
			TableName: "商品分类表",
			StructPtr: &ProductCategory{},
		},
		{
			TableName: "商品分类和属性的关系表",
			StructPtr: &ProductCategoryAttributeRelation{},
		},
		{
			TableName: "商品品牌表",
			StructPtr: &Brand{},
		},
		// 商品属性维护
		{
			TableName: "商品属性分类表",
			StructPtr: &ProductAttributeCategory{},
		},
		{
			TableName: "商品属性表",
			StructPtr: &ProductAttribute{},
		},
		{
			TableName: "商品属性值表",
			StructPtr: &ProductAttributeValue{},
		},
		//
		{
			TableName: "商品表",
			StructPtr: &Product{},
		},
		{
			TableName: "商品SKU表",
			StructPtr: &SkuStock{},
		},
		// 优惠
		{
			TableName: "商品阶梯价格表",
			StructPtr: &ProductLadder{},
		},
		{
			TableName: "商品满减表",
			StructPtr: &ProductFullReduction{},
		},
		{
			TableName: "商品会员价格表",
			StructPtr: &MemberPrice{},
		},
		// 评价回复
		{
			TableName: "商品评价表",
			StructPtr: &Comment{},
		},
		{
			TableName: "商品评价回复表",
			StructPtr: &CommentReplay{},
		},
		// 审计
		{
			TableName: "商品审核记录表",
			StructPtr: &ProductVertifyRecord{},
		},
		{
			TableName: "商品操作记录表",
			StructPtr: &ProductOperateLog{},
		},
		{
			TableName: "JSON动态配置",
			StructPtr: &JsonDynamicConfig{},
		},
	}

	cmsSchemas := []tableSchema{
		{
			TableName: "优选专区",
			StructPtr: &PrefrenceArea{},
		},
		{
			TableName: "优选专区和产品关系表",
			StructPtr: &PrefrenceAreaProductRelation{},
		},
		{
			TableName: "专题表",
			StructPtr: &Subject{},
		},
		{
			TableName: "专题商品关系表",
			StructPtr: &SubjectProductRelation{},
		},

		{
			TableName: "相册表",
			StructPtr: &Album{},
		},
		{
			TableName: "画册图片表",
			StructPtr: &AlbumPic{},
		},
		{
			TableName: "运费模版",
			StructPtr: &FeightTemplate{},
		},
	}

	omsSchemas := []tableSchema{
		{
			TableName: "订单表",
			StructPtr: &Order{},
		},
		{
			TableName: "订单商品信息表",
			StructPtr: &OrderItem{},
		},
		{
			TableName: "订单操作历史记录表",
			StructPtr: &OrderOperateHistory{},
		},
		{
			TableName: "退货原因表",
			StructPtr: &OrderReturnReason{},
		},
		{
			TableName: "订单退货申请表",
			StructPtr: &OrderReturnApply{},
		},
		{
			TableName: "公司收发货地址表",
			StructPtr: &CompanyAddress{},
		},
		{
			TableName: "购物车表",
			StructPtr: &CartItem{},
		},
	}

	smsSchemas := []tableSchema{
		{
			TableName: "限时购表",
			StructPtr: &FlashPromotion{},
		},
		//{
		//	TableName: "限时购场次表",
		//	StructPtr: &FlashPromotionSession{},
		//},
		//{
		//	TableName: "限时购与商品关系表",
		//	StructPtr: &FlashPromotionProductRelation{},
		//},
		//{
		//	TableName: "限时购通知记录表",
		//	StructPtr: &FlashPromotionLog{},
		//},
		// 优惠券
		{
			TableName: "优惠券表",
			StructPtr: &Coupon{},
		},
		{
			TableName: "优惠券和商品的关系表",
			StructPtr: &CouponProductRelation{},
		},
		{
			TableName: "优惠券和商品分类关系表",
			StructPtr: &CouponProductCategoryRelation{},
		},
		{
			TableName: "优惠券使用、领取历史表",
			StructPtr: &CouponHistory{},
		},
		// 首页
		{
			TableName: "首页品牌推荐表",
			StructPtr: &HomeBrand{},
		},
		{
			TableName: "新品推荐商品表",
			StructPtr: &HomeNewProduct{},
		},
		{
			TableName: "人气推荐商品表",
			StructPtr: &HomeRecommendProduct{},
		},
		{
			TableName: "首页专题推荐表",
			StructPtr: &HomeRecommendSubject{},
		},
		{
			TableName: "首页轮播广告表",
			StructPtr: &HomeAdvertise{},
		},
		// 用户信息
		{
			TableName: "会员表",
			StructPtr: &Member{},
		},
		{
			TableName: "会员收货地址表",
			StructPtr: &MemberReceiveAddress{},
		},
	}

	schemas := make([]tableSchema, 0)
	schemas = append(schemas, pmsSchemas...)
	schemas = append(schemas, cmsSchemas...)
	schemas = append(schemas, omsSchemas...)
	schemas = append(schemas, smsSchemas...)
	return autoMigrate(db, schemas)
}

// tableSchema 自动建表描述信息
type tableSchema struct {
	TableName string      // 表名
	StructPtr interface{} // 结构体指针
}

func autoMigrate(db *gorm.DB, schemas []tableSchema) error {
	for _, schema := range schemas {
		if err := db.
			Set("gorm:table_options", fmt.Sprintf("CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT='%s'", schema.TableName)).
			AutoMigrate(schema.StructPtr); err != nil {
			return err
		}
	}
	return nil
}
