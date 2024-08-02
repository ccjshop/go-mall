package entity

// CompanyAddress 公司收发货地址表
type CompanyAddress struct {
	ID            uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;comment:主键"`
	AddressName   string `gorm:"column:address_name;type:varchar(200);not null;default:'';comment:地址名称"`
	SendStatus    uint8  `gorm:"column:send_status;type:tinyint(4);unsigned;not null;default:0;comment:默认发货地址：0->否；1->是"`
	ReceiveStatus uint8  `gorm:"column:receive_status;type:tinyint(4);unsigned;not null;default:0;comment:是否默认收货地址：0->否；1->是"`
	Name          string `gorm:"column:name;type:varchar(64);not null;default:'';comment:收发货人姓名"`
	Phone         string `gorm:"column:phone;type:varchar(64);not null;default:'';comment:收货人电话"`
	Province      string `gorm:"column:province;type:varchar(64);not null;default:'';comment:省/直辖市"`
	City          string `gorm:"column:city;type:varchar(64);not null;default:'';comment:市"`
	Region        string `gorm:"column:region;type:varchar(64);not null;default:'';comment:区"`
	DetailAddress string `gorm:"column:detail_address;type:varchar(200);not null;default:'';comment:详细地址"`
	// 公共字段
	BaseTime
}

func (o CompanyAddress) TableName() string {
	return "oms_company_address"
}

type CompanyAddresses []*CompanyAddress
