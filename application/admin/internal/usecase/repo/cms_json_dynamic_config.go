package repo

import (
	"context"
	"errors"

	"github.com/ccjshop/go-mall/common/db"
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

// Create 创建JSON动态配置
func (r JsonDynamicConfigRepo) Create(ctx context.Context, jsonDynamicConfig *entity.JsonDynamicConfig) error {
	if jsonDynamicConfig.ID > 0 {
		return errors.New("illegal argument jsonDynamicConfig id exist")
	}
	return r.GenericDao.Create(ctx, jsonDynamicConfig)
}

// DeleteByID 根据主键ID删除JSON动态配置
func (r JsonDynamicConfigRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改JSON动态配置
func (r JsonDynamicConfigRepo) Update(ctx context.Context, jsonDynamicConfig *entity.JsonDynamicConfig) error {
	if jsonDynamicConfig.ID == 0 {
		return errors.New("illegal argument jsonDynamicConfig exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateJsonDynamicConfigField).Updates(jsonDynamicConfig).Error
}

// GetByID 根据主键ID查询JSON动态配置
func (r JsonDynamicConfigRepo) GetByID(ctx context.Context, id uint64) (*entity.JsonDynamicConfig, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByDBOption 根据动态条件查询JSON动态配置
func (r JsonDynamicConfigRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.JsonDynamicConfig, uint32, error) {
	var (
		res       = make([]*entity.JsonDynamicConfig, 0)
		pageTotal = int64(0)
		offset    = (pageNum - 1) * pageSize
	)

	session := r.GenericDao.DB.WithContext(ctx)
	for _, opt := range opts {
		session = opt(session)
	}

	session = session.Offset(int(offset)).Limit(int(pageSize)).Order("id desc").Find(&res).
		Offset(-1).Limit(-1).Count(&pageTotal)

	if err := session.Error; err != nil {
		return nil, 0, err
	}
	return res, uint32(pageTotal), nil
}
