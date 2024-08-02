package repo

import (
	"context"
	"errors"

	db "github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	"gorm.io/gorm"
)

// SkuStockRepo sku的库存
type SkuStockRepo struct {
	*db.GenericDao[entity.SkuStock, uint64]
}

// NewSkuStockRepo 创建
func NewSkuStockRepo(conn *gorm.DB) *SkuStockRepo {
	return &SkuStockRepo{
		GenericDao: &db.GenericDao[entity.SkuStock, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initSkuStockField)
}

var (
	// 全字段修改SkuStock那些字段不修改
	notUpdateSkuStockField = []string{
		"created_at",
	}
	updateSkuStockField []string
)

// InitSkuStockField 全字段修改
func initSkuStockField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.SkuStock{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateSkuStockField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateSkuStockField...)
	return nil
}

// GetByID 根据主键ID查询sku的库存
func (r SkuStockRepo) GetByID(ctx context.Context, id uint64) (*entity.SkuStock, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// Update 修改sku的库存
func (r SkuStockRepo) Update(ctx context.Context, skuStock *entity.SkuStock) error {
	if skuStock.ID == 0 {
		return errors.New("illegal argument skuStock exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateSkuStockField).Updates(skuStock).Error
}

// GetByProductID 根据商品ID查询sku的库存
func (r SkuStockRepo) GetByProductID(ctx context.Context, productID uint64) (entity.SkuStocks, error) {
	var (
		res = make([]*entity.SkuStock, 0)
	)
	if err := r.GenericDao.DB.WithContext(ctx).
		Where("product_id = ?", productID).
		Order("id desc").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// GetByProductIDs 根据商品ID查询sku的库存
func (r SkuStockRepo) GetByProductIDs(ctx context.Context, productIDs []uint64) (entity.SkuStocks, error) {
	var (
		res = make([]*entity.SkuStock, 0)
	)
	if err := r.GenericDao.DB.WithContext(ctx).
		Where("product_id in ?", productIDs).
		Order("id desc").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
