package repo

import (
	"context"
	"errors"

	"github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	"gorm.io/gorm"
)

// HomeAdvertiseRepo 首页轮播广告表
type HomeAdvertiseRepo struct {
	*db.GenericDao[entity.HomeAdvertise, uint64]
}

// NewHomeAdvertiseRepo 创建
func NewHomeAdvertiseRepo(conn *gorm.DB) *HomeAdvertiseRepo {
	return &HomeAdvertiseRepo{
		GenericDao: &db.GenericDao[entity.HomeAdvertise, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initHomeAdvertiseField)
}

var (
	// 全字段修改HomeAdvertise那些字段不修改
	notUpdateHomeAdvertiseField = []string{
		"created_at",
		"click_count",
		"order_count",
	}
	updateHomeAdvertiseField []string
)

// InitHomeAdvertiseField 全字段修改
func initHomeAdvertiseField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.HomeAdvertise{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateHomeAdvertiseField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateHomeAdvertiseField...)
	return nil
}

// Create 创建首页轮播广告表
func (r HomeAdvertiseRepo) Create(ctx context.Context, homeAdvertise *entity.HomeAdvertise) error {
	if homeAdvertise.ID > 0 {
		return errors.New("illegal argument homeAdvertise id exist")
	}
	return r.GenericDao.Create(ctx, homeAdvertise)
}

// DeleteByID 根据主键ID删除首页轮播广告表
func (r HomeAdvertiseRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改首页轮播广告表
func (r HomeAdvertiseRepo) Update(ctx context.Context, homeAdvertise *entity.HomeAdvertise) error {
	if homeAdvertise.ID == 0 {
		return errors.New("illegal argument homeAdvertise exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateHomeAdvertiseField).Updates(homeAdvertise).Error
}

// GetByID 根据主键ID查询首页轮播广告表
func (r HomeAdvertiseRepo) GetByID(ctx context.Context, id uint64) (*entity.HomeAdvertise, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByDBOption 根据动态条件查询首页轮播广告表
func (r HomeAdvertiseRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.HomeAdvertise, uint32, error) {
	var (
		res       = make([]*entity.HomeAdvertise, 0)
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
