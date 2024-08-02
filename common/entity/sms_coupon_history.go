package entity

// CouponHistory 优惠券使用历史表
// 用于存储会员领取及使用优惠券的记录，当会员领取到优惠券时，会产生一条优惠券的记录，需要注意的是它的使用状态：0->未使用；1->已使用；2->已过期。
type CouponHistory struct {
	ID         uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;"`
	MemberID   uint64 `gorm:"column:member_id;type:bigint;unsigned;not null;default:0;comment:会员id"`
	CouponID   uint64 `gorm:"column:coupon_id;type:bigint;unsigned;not null;default:0;comment:优惠券id"`
	CouponCode string `gorm:"column:coupon_code;type:varchar(64);not null;default:'';comment:优惠券码"`
	OrderID    uint64 `gorm:"column:order_id;type:bigint;unsigned;not null;default:0;comment:订单id"`
	OrderSN    string `gorm:"column:order_sn;type:varchar(100);not null;default:'';comment:订单号码"`
	GetType    uint8  `gorm:"column:get_type;type:tinyint(4);unsigned;not null;default:0;comment:获取类型：0->后台赠送；1->主动获取"`
	// 状态
	UseStatus uint8 `gorm:"column:use_status;type:tinyint(4);unsigned;not null;default:0;comment:使用状态：0->未使用；1->已使用；2->已过期"`
	// 时间
	CreateTime uint32 `gorm:"column:create_time;type:int(10);unsigned;not null;default:0;comment:创建时间"`
	UseTime    uint32 `gorm:"column:use_time;type:int(10);unsigned;not null;default:0;comment:使用时间"`
	// 冗余字段
	MemberNickname string `gorm:"column:member_nickname;type:varchar(64);not null;default:'';comment:领取人昵称"`
	// 公共字段
	BaseTime
}

func (h CouponHistory) TableName() string {
	return "sms_coupon_history"
}

type CouponHistories []*CouponHistory

// GetCouponIDs 收集所有的 coupon_id
func (h CouponHistories) GetCouponIDs() []uint64 {
	res := make([]uint64, 0)
	for _, item := range h {
		res = append(res, item.CouponID)
	}
	return res
}
