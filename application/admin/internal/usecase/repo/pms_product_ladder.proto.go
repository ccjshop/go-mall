package repo

import (
	"context"
	"errors"

	"github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	"gorm.io/gorm"
)

// ProductLadderRepo 产品阶梯价格表(只针对同商品)
type ProductLadderRepo struct {
	*db.GenericDao[entity.ProductLadder, uint64]
}

// NewProductLadderRepo 创建
func NewProductLadderRepo(conn *gorm.DB) *ProductLadderRepo {
	return &ProductLadderRepo{
		GenericDao: &db.GenericDao[entity.ProductLadder, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initProductLadderField)
}

var (
	// 全字段修改ProductLadder那些字段不修改
	notUpdateProductLadderField = []string{
		"created_at",
	}
	updateProductLadderField []string
)

// InitProductLadderField 全字段修改
func initProductLadderField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.ProductLadder{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateProductLadderField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateProductLadderField...)
	return nil
}

// Create 创建产品阶梯价格
func (r ProductLadderRepo) Create(ctx context.Context, productLadder *entity.ProductLadder) error {
	if productLadder.ID > 0 {
		return errors.New("illegal argument productLadder id exist")
	}
	return r.GenericDao.Create(ctx, productLadder)
}

// DeleteByID 根据主键ID删除产品阶梯价格
func (r ProductLadderRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改产品阶梯价格
func (r ProductLadderRepo) Update(ctx context.Context, productLadder *entity.ProductLadder) error {
	if productLadder.ID == 0 {
		return errors.New("illegal argument productLadder exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateProductLadderField).Updates(productLadder).Error
}

// GetByID 根据主键ID查询产品阶梯价格
func (r ProductLadderRepo) GetByID(ctx context.Context, id uint64) (*entity.ProductLadder, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByDBOption 根据动态条件查询产品阶梯价格
func (r ProductLadderRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.ProductLadder, uint32, error) {
	var (
		res       = make([]*entity.ProductLadder, 0)
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

// BatchCreateWithTX 创建阶梯价格
func (r ProductLadderRepo) BatchCreateWithTX(ctx context.Context, productID uint64, productLadders []*entity.ProductLadder) error {
	for _, productLadder := range productLadders {
		productLadder.ProductID = productID
	}
	tdb, err := db.GetTransactionDB(ctx)
	if err != nil {
		return err
	}
	return tdb.WithContext(ctx).Create(productLadders).Error
}

// DeleteByProductIDWithTX 根据商品ID删除记录
func (r ProductLadderRepo) DeleteByProductIDWithTX(ctx context.Context, productID uint64) error {
	tdb, err := db.GetTransactionDB(ctx)
	if err != nil {
		return err
	}
	return tdb.WithContext(ctx).Where("product_id = ?", productID).Delete(&entity.ProductLadder{}).Error
}

// GetByProductID 根据商品ID查询商品阶梯价格
func (r ProductLadderRepo) GetByProductID(ctx context.Context, productID uint64) ([]*entity.ProductLadder, error) {
	res := make([]*entity.ProductLadder, 0)
	if err := r.GenericDao.DB.WithContext(ctx).
		Where("product_id = ?", productID).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
