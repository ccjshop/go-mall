package repo

import (
	"context"

	db "github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	"gorm.io/gorm"
)

// JsonDynamicConfigRepo JSON动态配置
type JsonDynamicConfigRepo struct {
	*db.GenericDao[entity.JsonDynamicConfig, uint64]
}

// NewJsonDynamicConfigRepo 创建
func NewJsonDynamicConfigRepo(conn *gorm.DB) *JsonDynamicConfigRepo {
	return &JsonDynamicConfigRepo{
		GenericDao: &db.GenericDao[entity.JsonDynamicConfig, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initJsonDynamicConfigField)
}

var (
	// 全字段修改JsonDynamicConfig那些字段不修改
	notUpdateJsonDynamicConfigField = []string{
		"created_at",
	}
	updateJsonDynamicConfigField []string
)

// InitJsonDynamicConfigField 全字段修改
func initJsonDynamicConfigField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.JsonDynamicConfig{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateJsonDynamicConfigField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateJsonDynamicConfigField...)
	return nil
}

// GetByBizType 根据业务类型查询JSON动态配置
func (r JsonDynamicConfigRepo) GetByBizType(ctx context.Context, bizType entity.BizType) (string, error) {
	res := &entity.JsonDynamicConfig{}
	if err := r.GenericDao.DB.WithContext(ctx).Where("biz_type = ?", bizType).Find(&res).Error; err != nil {
		return "", err
	}
	return res.Content, nil
}
