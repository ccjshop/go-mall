package entity

import "github.com/shopspring/decimal"

// OrderReturnApply 订单退货申请表
type OrderReturnApply struct {
	ID      uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;comment:主键"`
	OrderID uint64 `gorm:"column:order_id;type:bigint;unsigned;not null;default:0;comment:订单id"`
	// 商品信息
	ProductPic       string          `gorm:"column:product_pic;type:varchar(500);not null;default:'';comment:商品图片"`
	ProductName      string          `gorm:"column:product_name;type:varchar(200);not null;default:'';comment:商品名称"`
	ProductBrand     string          `gorm:"column:product_brand;type:varchar(200);not null;default:'';comment:商品品牌"`
	ProductID        uint64          `gorm:"column:product_id;type:bigint;unsigned;not null;default:0;comment:退货商品id"`
	ProductRealPrice decimal.Decimal `gorm:"column:product_real_price;type:decimal(10,2);not null;default:0.00;comment:商品实际支付单价"`
	ProductAttr      string          `gorm:"column:product_attr;type:varchar(500);not null;default:'';comment:商品销售属性：颜色：红色；尺码：xl;"`
	ProductCount     uint32          `gorm:"column:product_count;type:int(10);unsigned;not null;default:0;comment:退货数量"`
	ProductPrice     decimal.Decimal `gorm:"column:product_price;type:decimal(10,2);not null;default:0.00;comment:商品单价"`
	//
	Status         uint8    `gorm:"column:status;type:tinyint(4);unsigned;not null;default:0;comment:申请状态：0->待处理；1->退货中；2->已完成；3->已拒绝"`
	OrderSN        string   `gorm:"column:order_sn;type:varchar(64);not null;default:'';comment:订单编号"`
	CreateTime     uint32   `gorm:"column:create_time;type:int(10);unsigned;not null;default:0;comment:申请时间"`
	MemberUsername string   `gorm:"column:member_username;type:varchar(64);not null;default:'';comment:会员用户名"`
	ReturnName     string   `gorm:"column:return_name;type:varchar(100);not null;default:'';comment:退货人姓名"`
	ReturnPhone    string   `gorm:"column:return_phone;type:varchar(100);not null;default:'';comment:退货人电话"`
	Reason         string   `gorm:"column:reason;type:varchar(200);not null;default:'';comment:原因"`
	Description    string   `gorm:"column:description;type:varchar(500);not null;default:'';comment:描述"`
	ProofPics      []string `gorm:"column:proof_pics;type:varchar(1000);not null;default:'';serializer:json;comment:凭证图片，json字符串数组"`
	//
	ReturnAmount     decimal.Decimal `gorm:"column:return_amount;type:decimal(10,2);not null;default:0.00;comment:退款金额"`
	CompanyAddressID uint64          `gorm:"column:company_address_id;type:bigint;unsigned;not null;default:0;comment:收货地址表id"` // oms_company_address#id
	// 商家-处理人
	HandleMan  string `gorm:"column:handle_man;type:varchar(100);not null;default:'';comment:处理人员"`
	HandleTime uint32 `gorm:"column:handle_time;type:int(10);unsigned;not null;default:0;comment:处理时间"`
	HandleNote string `gorm:"column:handle_note;type:varchar(500);not null;default:'';comment:处理备注"`
	// 商家-收货人
	ReceiveMan  string `gorm:"column:receive_man;type:varchar(100);not null;default:'';comment:收货人"`
	ReceiveTime uint32 `gorm:"column:receive_time;type:int(10);unsigned;not null;default:0;comment:收货时间"`
	ReceiveNote string `gorm:"column:receive_note;type:varchar(500);not null;default:'';comment:收货备注"`
	// 公共字段
	BaseTime
}

func (o OrderReturnApply) TableName() string {
	return "oms_order_return_apply"
}
