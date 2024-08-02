package entity

import "github.com/shopspring/decimal"

// Order 订单表
// 订单表，需要注意的是订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单。
type Order struct {
	// 基本信息
	ID       uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;comment:订单id"`
	MemberID uint64 `gorm:"column:member_id;type:bigint;unsigned;not null;default:0;comment:会员id"`
	OrderSN  string `gorm:"column:order_sn;type:varchar(64);not null;default:'';comment:订单编号"`
	Note     string `gorm:"column:note;type:varchar(500);not null;default:'';comment:订单备注"`
	// 优惠
	CouponID      uint64 `gorm:"column:coupon_id;type:bigint;unsigned;not null;default:0;comment:优惠券id"`
	PromotionInfo string `gorm:"column:promotion_info;type:varchar(100);not null;default:'';comment:活动信息"`
	// 积分成长值
	UseIntegration uint32 `gorm:"column:use_integration;type:int(10);unsigned;not null;default:0;comment:下单时使用的积分"`
	Integration    uint32 `gorm:"column:integration;type:int(10);unsigned;not null;default:0;comment:可以获得的积分"`
	Growth         uint32 `gorm:"column:growth;type:int(10);unsigned;not null;default:0;comment:可以活动的成长值"`
	// 类型
	PayType    uint8 `gorm:"column:pay_type;type:tinyint(4);unsigned;not null;default:0;comment:支付方式：0->未支付；1->支付宝；2->微信"`
	SourceType uint8 `gorm:"column:source_type;type:tinyint(4);unsigned;not null;default:0;comment:订单来源：0->PC订单；1->app订单"`
	OrderType  uint8 `gorm:"column:order_type;type:tinyint(4);unsigned;not null;default:0;comment:订单类型：0->正常订单；1->秒杀订单"`
	// 物流
	DeliveryCompany string `gorm:"column:delivery_company;type:varchar(64);not null;default:'';comment:物流公司(配送方式)"`
	DeliverySN      string `gorm:"column:delivery_sn;type:varchar(64);not null;default:'';comment:物流单号"`
	ReceiveTime     uint32 `gorm:"column:receive_time;type:int(10);unsigned;not null;default:0;comment:确认收货时间"`
	AutoConfirmDay  uint32 `gorm:"column:auto_confirm_day;type:int(10);unsigned;not null;default:0;comment:自动确认时间（天）"`
	// 收货人信息
	ReceiverName          string `gorm:"column:receiver_name;type:varchar(100);not null;default:'';comment:收货人姓名"`
	ReceiverPhone         string `gorm:"column:receiver_phone;type:varchar(32);not null;default:'';comment:收货人电话"`
	ReceiverPostCode      string `gorm:"column:receiver_post_code;type:varchar(32);not null;default:'';comment:收货人邮编"`
	ReceiverProvince      string `gorm:"column:receiver_province;type:varchar(32);not null;default:'';comment:省份/直辖市"`
	ReceiverCity          string `gorm:"column:receiver_city;type:varchar(32);not null;default:'';comment:城市"`
	ReceiverRegion        string `gorm:"column:receiver_region;type:varchar(32);not null;default:'';comment:区"`
	ReceiverDetailAddress string `gorm:"column:receiver_detail_address;type:varchar(200);not null;default:'';comment:详细地址"`
	// 费用信息 payAmount = totalAmount + freightAmount - promotionAmount - couponAmount - integrationAmount
	TotalAmount       decimal.Decimal `gorm:"column:total_amount;type:decimal(10,2);not null;default:0.00;comment:订单总金额"`
	FreightAmount     decimal.Decimal `gorm:"column:freight_amount;type:decimal(10,2);not null;default:0.00;comment:运费金额"`
	PromotionAmount   decimal.Decimal `gorm:"column:promotion_amount;type:decimal(10,2);not null;default:0.00;comment:促销优化金额（促销价、满减、阶梯价）"`
	CouponAmount      decimal.Decimal `gorm:"column:coupon_amount;type:decimal(10,2);not null;default:0.00;comment:优惠券抵扣金额"`
	IntegrationAmount decimal.Decimal `gorm:"column:integration_amount;type:decimal(10,2);not null;default:0.00;comment:积分抵扣金额"`
	PayAmount         decimal.Decimal `gorm:"column:pay_amount;type:decimal(10,2);not null;default:0.00;comment:应付金额（实际支付金额）"`
	DiscountAmount    decimal.Decimal `gorm:"column:discount_amount;type:decimal(10,2);not null;default:0.00;comment:管理员后台调整订单使用的折扣金额"`
	// 发票信息
	BillType          uint8  `gorm:"column:bill_type;type:tinyint(4);unsigned;not null;default:0;comment:发票类型：0->不开发票；1->电子发票；2->纸质发票"`
	BillHeader        string `gorm:"column:bill_header;type:varchar(200);not null;default:'';comment:发票抬头"`
	BillContent       string `gorm:"column:bill_content;type:varchar(200);not null;default:'';comment:发票内容"`
	BillReceiverPhone string `gorm:"column:bill_receiver_phone;type:varchar(32);not null;default:'';comment:收票人电话"`
	BillReceiverEmail string `gorm:"column:bill_receiver_email;type:varchar(64);not null;default:'';comment:收票人邮箱"`
	// 状态
	Status        uint8 `gorm:"column:status;type:tinyint(4);unsigned;not null;default:0;comment:订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单"`
	ConfirmStatus uint8 `gorm:"column:confirm_status;type:tinyint(4);unsigned;not null;default:0;comment:确认收货状态：0->未确认；1->已确认"`
	DeleteStatus  uint8 `gorm:"column:delete_status;type:tinyint(4);unsigned;not null;default:0;comment:删除状态：0->未删除；1->已删除"`
	// 时间
	PaymentTime  uint32 `gorm:"column:payment_time;type:int(10);unsigned;not null;default:0;comment:支付时间"`
	DeliveryTime uint32 `gorm:"column:delivery_time;type:int(10);unsigned;not null;default:0;comment:发货时间"`
	CommentTime  uint32 `gorm:"column:comment_time;type:int(10);unsigned;not null;default:0;comment:评价时间"`
	CreateTime   uint32 `gorm:"column:create_time;type:int(10);unsigned;not null;default:0;comment:提交时间"`
	ModifyTime   uint32 `gorm:"column:modify_time;type:int(10);unsigned;not null;default:0;comment:修改时间"`
	// 冗余字段
	MemberUsername string `gorm:"column:member_username;type:varchar(64);not null;default:'';comment:用户帐号"`
	// 公共字段
	BaseTime
}

func (o Order) TableName() string {
	return "oms_order"
}
