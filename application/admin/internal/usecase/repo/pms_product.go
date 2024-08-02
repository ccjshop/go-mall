package repo

import (
	"context"
	"errors"

	db2 "github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	"gorm.io/gorm"
)

// ProductRepo 商品信息表
type ProductRepo struct {
	*db2.GenericDao[entity.Product, uint64]
}

// NewProductRepo 创建
func NewProductRepo(conn *gorm.DB) *ProductRepo {
	return &ProductRepo{
		GenericDao: &db2.GenericDao[entity.Product, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initProductField)
}

var (
	// 全字段修改Product那些字段不修改
	notUpdateProductField = []string{
		"created_at",
	}
	updateProductField []string
)

// InitProductField 全字段修改
func initProductField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.Product{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateProductField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateProductField...)
	return nil
}

func (r ProductRepo) WithByID(id uint64) db2.DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	}
}

func (r ProductRepo) WithByName(name string) db2.DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name like ?", "%"+name+"%")
	}
}

func (r ProductRepo) WithByProductSN(productSN string) db2.DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("product_sn = ?", productSN)
	}
}
func (r ProductRepo) WithByBrandID(brandID uint64) db2.DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("brand_id = ?", brandID)
	}
}
func (r ProductRepo) WithByPublishStatus(publishStatus uint32) db2.DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("publish_status = ?", publishStatus)
	}
}
func (r ProductRepo) WithByVerifyStatus(verifyStatus uint32) db2.DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("verify_status = ?", verifyStatus)
	}
}
func (r ProductRepo) WithByProductCategoryID(productCategoryID uint64) db2.DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("product_category_id = ?", productCategoryID)
	}
}

// Create 创建商品
func (r ProductRepo) Create(ctx context.Context, product *entity.Product) error {
	if product.ID > 0 {
		return errors.New("illegal argument product id exist")
	}
	return r.GenericDao.Create(ctx, product)
}

// DeleteByID 根据主键ID删除商品
func (r ProductRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改商品
func (r ProductRepo) Update(ctx context.Context, product *entity.Product) error {
	if product.ID == 0 {
		return errors.New("illegal argument product exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateProductField).Updates(product).Error
}

// GetByID 根据主键ID查询商品
func (r ProductRepo) GetByID(ctx context.Context, id uint64) (*entity.Product, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByDBOption 根据动态条件查询商品
func (r ProductRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db2.DBOption) (entity.Products, uint32, error) {
	var (
		res       = make([]*entity.Product, 0)
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

// CreateWithTX 创建商品
func (r ProductRepo) CreateWithTX(ctx context.Context, product *entity.Product) error {
	tdb, err := db2.GetTransactionDB(ctx)
	if err != nil {
		return err
	}
	return tdb.WithContext(ctx).Create(product).Error
}

// UpdateWithTX 修改商品
func (r ProductRepo) UpdateWithTX(ctx context.Context, product *entity.Product) error {
	tdb, err := db2.GetTransactionDB(ctx)
	if err != nil {
		return err
	}
	return tdb.WithContext(ctx).Select(updateProductField).Updates(product).Error
}
