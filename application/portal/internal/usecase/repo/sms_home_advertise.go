package repo

import (
	"context"

	"github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
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

// GetHomeAdvertises 获取首页广告
func (r HomeAdvertiseRepo) GetHomeAdvertises(ctx context.Context) ([]*entity.HomeAdvertise, error) {
	res := make([]*entity.HomeAdvertise, 0)
	if err := r.GenericDao.DB.WithContext(ctx).
		Where("type = 1").
		Where("status = 1").
		Order("sort desc").
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
