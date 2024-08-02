package repo

import (
	"context"
	"errors"

	db2 "github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	"gorm.io/gorm"
)

// PrefrenceAreaRepo 优选专区
type PrefrenceAreaRepo struct {
	*db2.GenericDao[entity.PrefrenceArea, uint64]
}

// NewPrefrenceAreaRepo 创建
func NewPrefrenceAreaRepo(conn *gorm.DB) *PrefrenceAreaRepo {
	return &PrefrenceAreaRepo{
		GenericDao: &db2.GenericDao[entity.PrefrenceArea, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initPrefrenceAreaField)
}

var (
	// 全字段修改PrefrenceArea那些字段不修改
	notUpdatePrefrenceAreaField = []string{
		"created_at",
	}
	updatePrefrenceAreaField []string
)

// InitPrefrenceAreaField 全字段修改
func initPrefrenceAreaField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.PrefrenceArea{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updatePrefrenceAreaField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdatePrefrenceAreaField...)
	return nil
}

// Create 创建优选专区
func (r PrefrenceAreaRepo) Create(ctx context.Context, prefrenceArea *entity.PrefrenceArea) error {
	if prefrenceArea.ID > 0 {
		return errors.New("illegal argument prefrenceArea id exist")
	}
	return r.GenericDao.Create(ctx, prefrenceArea)
}

// DeleteByID 根据主键ID删除优选专区
func (r PrefrenceAreaRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改优选专区
func (r PrefrenceAreaRepo) Update(ctx context.Context, prefrenceArea *entity.PrefrenceArea) error {
	if prefrenceArea.ID == 0 {
		return errors.New("illegal argument prefrenceArea exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updatePrefrenceAreaField).Updates(prefrenceArea).Error
}

// GetByID 根据主键ID查询优选专区
func (r PrefrenceAreaRepo) GetByID(ctx context.Context, id uint64) (*entity.PrefrenceArea, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByDBOption 根据动态条件查询优选专区
func (r PrefrenceAreaRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db2.DBOption) ([]*entity.PrefrenceArea, uint32, error) {
	var (
		res       = make([]*entity.PrefrenceArea, 0)
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
