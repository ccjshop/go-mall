package entity

// Member 会员表
// 前台会员表，定义了前台会员的一些基本信息。
type Member struct {
	// 账号信息
	ID       uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;comment:用户id"`
	Username string `gorm:"column:username;type:varchar(64);not null;default:'';comment:用户名"`
	Password string `gorm:"column:password;type:varchar(64);not null;default:'';comment:密码"`
	// 基本信息
	Nickname              string `gorm:"column:nickname;type:varchar(64);not null;default:'';comment:昵称"`
	Icon                  string `gorm:"column:icon;type:varchar(500);not null;default:'';comment:头像"`
	Gender                uint8  `gorm:"column:gender;type:tinyint(4);unsigned;not null;default:0;comment:性别：0->未知；1->男；2->女"`
	Birthday              uint32 `gorm:"column:birthday;type:int(10);unsigned;not null;default:0;comment:生日"` // 整数，YYYYMMDD
	PersonalizedSignature string `gorm:"column:personalized_signature;type:varchar(200);not null;default:'';comment:个性签名"`
	Phone                 string `gorm:"column:phone;type:varchar(64);not null;default:'';comment:手机号码"`
	City                  string `gorm:"column:city;type:varchar(64);not null;default:'';comment:所在城市"`
	Job                   string `gorm:"column:job;type:varchar(100);not null;default:'';comment:职业"`
	//
	MemberLevelID uint64 `gorm:"column:member_level_id;type:bigint;unsigned;not null;default:0;comment:会员等级id"`
	SourceType    uint8  `gorm:"column:source_type;type:tinyint(4);unsigned;not null;default:0;comment:用户来源"`
	// 业务字段
	Integration        uint32 `gorm:"column:integration;type:int(10);unsigned;not null;default:0;comment:积分"`
	Growth             uint32 `gorm:"column:growth;type:int(10);unsigned;not null;default:0;comment:成长值"`
	LuckeyCount        uint32 `gorm:"column:luckey_count;type:int(10);unsigned;not null;default:0;comment:剩余抽奖次数"`
	HistoryIntegration uint32 `gorm:"column:history_integration;type:int(10);unsigned;not null;default:0;comment:历史积分数量"`
	//
	Status     uint8  `gorm:"column:status;type:tinyint(4);unsigned;not null;default:0;comment:帐号启用状态:0->禁用；1->启用"`
	CreateTime uint32 `gorm:"column:create_time;type:int(10);unsigned;not null;default:0;comment:注册时间"`
	// 公共字段
	BaseTime
}

func (m Member) TableName() string {
	return "ums_member"
}

type Members []*Member

// GetMap 获取集合 key=id value=*Member
func (m Members) GetMap() map[uint64]*Member {
	res := make(map[uint64]*Member)
	for _, item := range m {
		res[item.ID] = item
	}
	return res
}
