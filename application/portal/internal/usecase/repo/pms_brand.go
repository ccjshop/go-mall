package repo

import (
	"context"

	db "github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"gorm.io/gorm"
)

// BrandRepo 商品品牌表
type BrandRepo struct {
	*db.GenericDao[entity.Brand, uint64]
}

// NewBrandRepo 创建
func NewBrandRepo(conn *gorm.DB) *BrandRepo {
	return &BrandRepo{
		GenericDao: &db.GenericDao[entity.Brand, uint64]{
			DB: conn,
		},
	}
}

// GetByID 根据主键ID查询商品品牌表
func (r BrandRepo) GetByID(ctx context.Context, id uint64) (*entity.Brand, error) {
	return r.GenericDao.GetByID(ctx, id)
}
