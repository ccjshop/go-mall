package entity

import (
	pb "github.com/ccjshop/go-mall/proto/mall"
	"github.com/shopspring/decimal"
)

// Coupon 优惠券表
// 用于存储优惠券信息，需要注意的是优惠券的使用类型：0->全场通用；1->指定分类；2->指定商品，不同使用类型的优惠券使用范围不一样。
type Coupon struct {
	// 基本信息
	ID     uint64          `gorm:"column:id;type:bigint;primary_key;auto_increment;"`
	Name   string          `gorm:"column:name;type:varchar(100);not null;default:'';comment:名称"`
	Amount decimal.Decimal `gorm:"column:amount;type:decimal(10,2);not null;default:0.00;comment:金额"`
	Note   string          `gorm:"column:note;type:varchar(200);not null;default:'';comment:备注"`
	Code   string          `gorm:"column:code;type:varchar(64);not null;default:'';comment:优惠码"`
	// 数量
	Count        uint32 `gorm:"column:count;type:int(10);unsigned;not null;default:0;comment:数量"`
	PublishCount uint32 `gorm:"column:publish_count;type:int(10);unsigned;not null;default:0;comment:发行数量"`
	UseCount     uint32 `gorm:"column:use_count;type:int(10);unsigned;not null;default:0;comment:已使用数量"`
	ReceiveCount uint32 `gorm:"column:receive_count;type:int(10);unsigned;not null;default:0;comment:领取数量"`
	// 类型
	Type    uint8            `gorm:"column:type;type:tinyint(4);unsigned;not null;default:0;comment:优惠卷类型；0->全场赠券；1->会员赠券；2->购物赠券；3->注册赠券"`
	UseType pb.CouponUseType `gorm:"column:use_type;type:tinyint(4);unsigned;not null;default:0;comment:使用类型：0->全场通用；1->指定分类；2->指定商品"`
	// 领取限制
	PerLimit    uint32 `gorm:"column:per_limit;type:int(10);unsigned;not null;default:0;comment:每人限领张数"`
	EnableTime  uint32 `gorm:"column:enable_time;type:int(10);unsigned;not null;default:0;comment:可以领取的日期"`
	MemberLevel uint8  `gorm:"column:member_level;type:tinyint(4);unsigned;not null;default:0;comment:可领取的会员类型：0->无限制"`
	// 使用限制
	MinPoint  decimal.Decimal `gorm:"column:min_point;type:decimal(10,2);not null;default:0.00;comment:使用门槛；0表示无门槛"`
	Platform  uint8           `gorm:"column:platform;type:tinyint(4);unsigned;not null;default:0;comment:使用平台：0->全部；1->移动；2->PC"`
	StartTime uint32          `gorm:"column:start_time;type:int(10);unsigned;not null;default:0;comment:开始使用时间"`
	EndTime   uint32          `gorm:"column:end_time;type:int(10);unsigned;not null;default:0;comment:结束使用时间"`
	// 公共字段
	BaseTime
}

func (c Coupon) TableName() string {
	return "sms_coupon"
}

type Coupons []*Coupon
