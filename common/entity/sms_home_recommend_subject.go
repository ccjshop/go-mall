package entity

// HomeRecommendSubject 首页专题推荐表
// 用于管理首页显示的专题精选信息。
type HomeRecommendSubject struct {
	ID              uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;"`
	SubjectID       uint64 `gorm:"column:subject_id;type:bigint;unsigned;not null;default:0;comment:专题id"`
	SubjectName     string `gorm:"column:subject_name;type:varchar(64);not null;default:'';comment:专题名称"`
	RecommendStatus uint8  `gorm:"column:recommend_status;type:tinyint(4);unsigned;not null;default:0;comment:推荐状态：0->不推荐;1->推荐"`
	Sort            uint32 `gorm:"column:sort;type:int(10);unsigned;not null;default:0;comment:排序"`
	// 公共字段
	BaseTime
}

func (s HomeRecommendSubject) TableName() string {
	return "sms_home_recommend_subject"
}
