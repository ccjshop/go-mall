package repo

import (
	"context"
	"errors"

	db2 "github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	"gorm.io/gorm"
)

// SkuStockRepo sku的库存
type SkuStockRepo struct {
	*db2.GenericDao[entity.SkuStock, uint64]
}

// NewSkuStockRepo 创建
func NewSkuStockRepo(conn *gorm.DB) *SkuStockRepo {
	return &SkuStockRepo{
		GenericDao: &db2.GenericDao[entity.SkuStock, uint64]{
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

func (r SkuStockRepo) WithByProductID(productID uint64) db2.DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("product_id = ?", productID)
	}
}
func (r SkuStockRepo) WithBySkuCode(skuCode string) db2.DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("sku_code like ?", "%"+skuCode+"%")
	}
}

// Create 创建sku的库存
func (r SkuStockRepo) Create(ctx context.Context, skuStock *entity.SkuStock) error {
	if skuStock.ID > 0 {
		return errors.New("illegal argument skuStock id exist")
	}
	return r.GenericDao.Create(ctx, skuStock)
}

// DeleteByID 根据主键ID删除sku的库存
func (r SkuStockRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改sku的库存
func (r SkuStockRepo) Update(ctx context.Context, skuStock *entity.SkuStock) error {
	if skuStock.ID == 0 {
		return errors.New("illegal argument skuStock exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateSkuStockField).Updates(skuStock).Error
}

// GetByID 根据主键ID查询sku的库存
func (r SkuStockRepo) GetByID(ctx context.Context, id uint64) (*entity.SkuStock, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByDBOption 根据动态条件查询sku的库存
func (r SkuStockRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db2.DBOption) ([]*entity.SkuStock, uint32, error) {
	var (
		res       = make([]*entity.SkuStock, 0)
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

// GetByProductID 根据商品ID查询sku的库存
func (r SkuStockRepo) GetByProductID(ctx context.Context, productID uint64) ([]*entity.SkuStock, error) {
	var (
		res = make([]*entity.SkuStock, 0)
	)
	if err := r.GenericDao.DB.WithContext(ctx).
		Where("product_id = ?", productID).
		Order("id asc").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// BatchCreateWithTX 创建sku库存
func (r SkuStockRepo) BatchCreateWithTX(ctx context.Context, productID uint64, skuStocks []*entity.SkuStock) error {
	for _, skuStock := range skuStocks {
		skuStock.ProductID = productID
	}
	tdb, err := db2.GetTransactionDB(ctx)
	if err != nil {
		return err
	}
	return tdb.WithContext(ctx).Create(skuStocks).Error
}

// BatchUpdateOrInsertSkuStock 批量插入或者更新
func (r SkuStockRepo) BatchUpdateOrInsertSkuStock(ctx context.Context, stocks []*entity.SkuStock) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, stock := range stocks {
			findStock := &entity.SkuStock{}
			findErr := tx.Where("id = ?", stock.ID).First(&findStock).Error
			if findErr == nil {
				// 没有错误直接更新
				if err := tx.Where("id = ?", stock.ID).Select(updateSkuStockField).Updates(stock).Error; err != nil {
					return err
				}
			} else {
				if errors.Is(findErr, gorm.ErrRecordNotFound) {
					// 记录没找到插入
					if err := tx.Create(stock).Error; err != nil {
						return err
					}
				} else {
					// 返回错误
					return findErr
				}
			}
		}
		return nil
	})
}

// DeleteByProductIDWithTX 根据商品ID删除记录
func (r SkuStockRepo) DeleteByProductIDWithTX(ctx context.Context, productID uint64) error {
	tdb, err := db2.GetTransactionDB(ctx)
	if err != nil {
		return err
	}
	return tdb.WithContext(ctx).Where("product_id = ?", productID).Delete(&entity.SkuStock{}).Error
}

// BatchDeleteByIDWithTX 根据ID删除记录
func (r SkuStockRepo) BatchDeleteByIDWithTX(ctx context.Context, ids []uint64) error {
	tdb, err := db2.GetTransactionDB(ctx)
	if err != nil {
		return err
	}
	return tdb.WithContext(ctx).Where("id in ?", ids).Delete(&entity.SkuStock{}).Error
}

// BatchUpDateByIDWithTX 根据ID修改记录
func (r SkuStockRepo) BatchUpDateByIDWithTX(ctx context.Context, skuStocks []*entity.SkuStock) error {
	tdb, err := db2.GetTransactionDB(ctx)
	if err != nil {
		return err
	}
	for _, skuStock := range skuStocks {
		// 没有错误直接更新
		if err := tdb.Where("id = ?", skuStock.ID).Select(updateSkuStockField).Updates(skuStock).Error; err != nil {
			return err
		}
	}
	return nil
}
