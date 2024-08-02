package entity

// MemberReceiveAddress 会员收货地址表
// 会员收货地址表，定义了前台会员的收货地址信息。
type MemberReceiveAddress struct {
	ID            uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;comment:地址id"`
	MemberID      uint64 `gorm:"column:member_id;type:bigint;unsigned;not null;default:0;comment:会员id"`
	Name          string `gorm:"column:name;type:varchar(100);not null;default:'';comment:收货人名称"`
	PhoneNumber   string `gorm:"column:phone_number;type:varchar(64);not null;default:'';comment:电话号码"`
	DefaultStatus uint8  `gorm:"column:default_status;type:tinyint(4);unsigned;not null;default:0;comment:是否为默认"`
	// 地址
	PostCode      string `gorm:"column:post_code;type:varchar(100);not null;default:'';comment:邮政编码"`
	Province      string `gorm:"column:province;type:varchar(100);not null;default:'';comment:省份/直辖市"`
	City          string `gorm:"column:city;type:varchar(100);not null;default:'';comment:城市"`
	Region        string `gorm:"column:region;type:varchar(100);not null;default:'';comment:区"`
	DetailAddress string `gorm:"column:detail_address;type:varchar(128);not null;default:'';comment:详细地址(街道)"`
	// 公共字段
	BaseTime
}

func (a MemberReceiveAddress) TableName() string {
	return "ums_member_receive_address"
}

type MemberReceiveAddresses []*MemberReceiveAddress
