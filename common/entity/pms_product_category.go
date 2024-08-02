package entity

// ProductCategory 商品分类表
type ProductCategory struct {
	ID          uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;comment:分类编号"`
	ParentID    uint64 `gorm:"column:parent_id;type:bigint;unsigned;not null;default:0;comment:上级分类的编号：0表示一级分类"`
	Name        string `gorm:"column:name;type:varchar(64);not null;default:'';comment:分类名称"`
	Icon        string `gorm:"column:icon;type:varchar(255);not null;default:'';comment:图标"`
	ProductUnit string `gorm:"column:product_unit;type:varchar(64);not null;default:'';comment:商品单位"`
	Sort        uint32 `gorm:"column:sort;type:int(10);unsigned;not null;default:0;comment:排序"`
	//
	Description string `gorm:"column:description;type:text;not null;comment:描述"`
	Keywords    string `gorm:"column:keywords;type:varchar(255);not null;default:'';comment:关键字"`
	// 状态
	NavStatus  uint8 `gorm:"column:nav_status;type:tinyint(4);unsigned;not null;default:0;comment:是否显示在导航栏：0->不显示；1->显示"`         //
	ShowStatus uint8 `gorm:"column:show_status;type:tinyint(4);unsigned;not null;default:0;comment:显示状态，控制在移动端是否展示：0->不显示；1->显示"` //
	// 计算得出
	Level uint8 `gorm:"column:level;type:tinyint(4);not null;default:0;comment:分类级别：0->1级；1->2级"`
	// 冗余字段
	ProductCount uint32 `gorm:"column:product_count;type:int(10);unsigned;not null;default:0;comment:商品数量"`
	// 公共字段
	BaseTime
}

func (c ProductCategory) TableName() string {
	return "pms_product_category"
}

// ProductCategories 商品分类集合
type ProductCategories []*ProductCategory

// NameMap 获取分类名 key=id value=name
func (c ProductCategories) NameMap() map[uint64]string {
	res := make(map[uint64]string)
	for _, category := range c {
		res[category.ID] = category.Name
	}
	return res
}
