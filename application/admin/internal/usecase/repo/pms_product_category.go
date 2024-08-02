package repo

import (
	"context"
	"errors"

	db2 "github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	"gorm.io/gorm"
)

// ProductCategoryRepo 商品分类表
type ProductCategoryRepo struct {
	*db2.GenericDao[entity.ProductCategory, uint64]
}

// NewProductCategoryRepo 创建
func NewProductCategoryRepo(conn *gorm.DB) *ProductCategoryRepo {
	return &ProductCategoryRepo{
		GenericDao: &db2.GenericDao[entity.ProductCategory, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initProductCategoryField)
}

var (
	// 全字段修改PmsProductCategory那些字段不修改
	notUpdateProductCategoryField = []string{
		"product_count",
		"level",
		"created_at",
	}
	updateProductCategoryField []string
)

func initProductCategoryField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.ProductCategory{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateProductCategoryField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateProductCategoryField...)
	return nil
}

func (p ProductCategoryRepo) WithByParentID(parentID uint64) db2.DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("parent_id = ?", parentID)
	}
}

func (p ProductCategoryRepo) WithByID(id uint64) db2.DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	}
}
func (p ProductCategoryRepo) WithByName(name string) db2.DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name like ?", "%"+name+"%")
	}
}

// Create 创建商品分类
func (p ProductCategoryRepo) Create(ctx context.Context, productCategory *entity.ProductCategory) error {
	if productCategory.ID > 0 {
		return errors.New("productCategory id exist")
	}
	return p.GenericDao.Create(ctx, productCategory)
}

// DeleteByID 根据主键ID删除商品分类
func (p ProductCategoryRepo) DeleteByID(ctx context.Context, categoryID uint64) error {
	return p.GenericDao.DeleteByID(ctx, categoryID)
}

// Update 修改商品分类
func (p ProductCategoryRepo) Update(ctx context.Context, productCategory *entity.ProductCategory) error {
	if productCategory.ID == 0 {
		return errors.New("productCategory not exist")
	}
	return p.GenericDao.DB.WithContext(ctx).Select(updateProductCategoryField).Updates(productCategory).Error
}

// GetByID 根据主键ID查询商品分类
func (p ProductCategoryRepo) GetByID(ctx context.Context, id uint64) (*entity.ProductCategory, error) {
	return p.GenericDao.GetByID(ctx, id)
}

// GetByIDs 根据主键ID查询商品分类
func (p ProductCategoryRepo) GetByIDs(ctx context.Context, ids []uint64) (entity.ProductCategories, error) {
	res := make([]*entity.ProductCategory, 0)
	if err := p.GenericDao.DB.WithContext(ctx).Where("id in ?", ids).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateWithTX 创建商品分类
func (p ProductCategoryRepo) CreateWithTX(ctx context.Context, productCategory *entity.ProductCategory) error {
	tdb, err := db2.GetTransactionDB(ctx)
	if err != nil {
		return err
	}
	return tdb.WithContext(ctx).Create(productCategory).Error
}

// UpdateWithTX 修改商品分类
func (p ProductCategoryRepo) UpdateWithTX(ctx context.Context, productCategory *entity.ProductCategory) error {
	if productCategory.ID == 0 {
		return errors.New("productCategory not exist")
	}
	tdb, err := db2.GetTransactionDB(ctx)
	if err != nil {
		return err
	}
	return tdb.WithContext(ctx).Select(updateProductCategoryField).Updates(productCategory).Error
}

// GetByDBOption 根据动态条件查询商品分类
func (p ProductCategoryRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db2.DBOption) ([]*entity.ProductCategory, uint32, error) {
	var (
		res       = make([]*entity.ProductCategory, 0)
		pageTotal = int64(0)
		offset    = (pageNum - 1) * pageSize
	)

	db := p.GenericDao.DB.WithContext(ctx)
	for _, opt := range opts {
		db = opt(db)
	}

	db = db.Offset(int(offset)).Limit(int(pageSize)).Order("id desc").Find(&res).
		Offset(-1).Limit(-1).Count(&pageTotal)

	if err := db.Error; err != nil {
		return nil, 0, err
	}
	return res, uint32(pageTotal), nil
}

// UpdateFieldByID 根据ID修改
func (p ProductCategoryRepo) UpdateFieldByID(ctx context.Context, category *entity.ProductCategory, fields ...string) error {
	if category.ID == 0 || len(fields) == 0 {
		return errors.New("illegal argument")
	}
	return p.DB.WithContext(ctx).Select(fields).Updates(category).Error
}

// UpdateProductCategoryNavStatus 修改导航栏显示状态
func (p ProductCategoryRepo) UpdateProductCategoryNavStatus(ctx context.Context, categoryIDs []uint64, navStatus uint32) error {
	return p.DB.WithContext(ctx).
		Model(&entity.ProductCategory{}).
		Where("id in ?", categoryIDs).
		Update("nav_status", navStatus).Error
}

// UpdateProductCategoryShowStatus 修改显示状态
func (p ProductCategoryRepo) UpdateProductCategoryShowStatus(ctx context.Context, categoryIDs []uint64, showStatus uint32) error {
	return p.DB.WithContext(ctx).
		Model(&entity.ProductCategory{}).
		Where("id in ?", categoryIDs).
		Update("show_status", showStatus).Error
}
