package entity

// CommentReplay 商品评论回复表
type CommentReplay struct {
	ID             uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;comment:主键"`
	CommentID      uint64 `gorm:"column:comment_id;type:bigint;not null;default:0;comment:评论id"`
	MemberNickName string `gorm:"column:member_nick_name;type:varchar(255);not null;default:'';comment:会员昵称"`
	MemberIcon     string `gorm:"column:member_icon;type:varchar(255);not null;default:'';comment:会员头像"`
	Content        string `gorm:"column:content;type:varchar(1000);not null;default:'';comment:内容"`
	CreateTime     uint32 `gorm:"column:create_time;type:int(10);unsigned;not null;default:0;comment:创建时间"`
	Type           uint32 `gorm:"column:type;type:tinyint(4);unsigned;not null;default:0;comment:评论人员类型；0->会员；1->管理员"`
	// 公共字段
	BaseTime
}

func (c CommentReplay) TableName() string {
	return "pms_comment_replay"
}
