package repo

import (
	"context"
	"errors"

	db2 "github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	"gorm.io/gorm"
)

// ProductFullReductionRepo 产品满减表(只针对同商品)
type ProductFullReductionRepo struct {
	*db2.GenericDao[entity.ProductFullReduction, uint64]
}

// NewProductFullReductionRepo 创建
func NewProductFullReductionRepo(conn *gorm.DB) *ProductFullReductionRepo {
	return &ProductFullReductionRepo{
		GenericDao: &db2.GenericDao[entity.ProductFullReduction, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initProductFullReductionField)
}

var (
	// 全字段修改ProductFullReduction那些字段不修改
	notUpdateProductFullReductionField = []string{
		"created_at",
	}
	updateProductFullReductionField []string
)

// InitProductFullReductionField 全字段修改
func initProductFullReductionField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.ProductFullReduction{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateProductFullReductionField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateProductFullReductionField...)
	return nil
}

// Create 创建产品满减
func (r ProductFullReductionRepo) Create(ctx context.Context, productFullReduction *entity.ProductFullReduction) error {
	if productFullReduction.ID > 0 {
		return errors.New("illegal argument productFullReduction id exist")
	}
	return r.GenericDao.Create(ctx, productFullReduction)
}

// DeleteByID 根据主键ID删除产品满减
func (r ProductFullReductionRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改产品满减
func (r ProductFullReductionRepo) Update(ctx context.Context, productFullReduction *entity.ProductFullReduction) error {
	if productFullReduction.ID == 0 {
		return errors.New("illegal argument productFullReduction exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateProductFullReductionField).Updates(productFullReduction).Error
}

// GetByID 根据主键ID查询产品满减
func (r ProductFullReductionRepo) GetByID(ctx context.Context, id uint64) (*entity.ProductFullReduction, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByDBOption 根据动态条件查询产品满减
func (r ProductFullReductionRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db2.DBOption) ([]*entity.ProductFullReduction, uint32, error) {
	var (
		res       = make([]*entity.ProductFullReduction, 0)
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

// BatchCreateWithTX 创建产品满减
func (r ProductFullReductionRepo) BatchCreateWithTX(ctx context.Context, productID uint64, productFullReductions []*entity.ProductFullReduction) error {
	for _, productFullReduction := range productFullReductions {
		productFullReduction.ProductID = productID
	}
	tdb, err := db2.GetTransactionDB(ctx)
	if err != nil {
		return err
	}
	return tdb.WithContext(ctx).Create(productFullReductions).Error
}

// DeleteByProductIDWithTX 根据商品ID删除记录
func (r ProductFullReductionRepo) DeleteByProductIDWithTX(ctx context.Context, productID uint64) error {
	tdb, err := db2.GetTransactionDB(ctx)
	if err != nil {
		return err
	}
	return tdb.WithContext(ctx).Where("product_id = ?", productID).Delete(&entity.ProductFullReduction{}).Error
}

// GetByProductID 根据产品ID查询产品满减
func (r ProductFullReductionRepo) GetByProductID(ctx context.Context, productID uint64) ([]*entity.ProductFullReduction, error) {
	res := make([]*entity.ProductFullReduction, 0)
	if err := r.GenericDao.DB.WithContext(ctx).
		Where("product_id = ?", productID).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
