package repo

import (
	"context"
	"errors"

	db2 "github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	"gorm.io/gorm"
)

// ProductAttributeCategoryRepo 产品属性分类表
type ProductAttributeCategoryRepo struct {
	*db2.GenericDao[entity.ProductAttributeCategory, uint64]
}

// NewProductAttributeCategoryRepo 创建
func NewProductAttributeCategoryRepo(conn *gorm.DB) *ProductAttributeCategoryRepo {
	return &ProductAttributeCategoryRepo{
		GenericDao: &db2.GenericDao[entity.ProductAttributeCategory, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initProductAttributeCategoryField)
}

var (
	// 全字段修改ProductAttributeCategory那些字段不修改
	notUpdateProductAttributeCategoryField = []string{
		"created_at",
	}
	updateProductAttributeCategoryField []string
)

// initProductAttributeCategoryField 全字段修改
func initProductAttributeCategoryField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.ProductAttributeCategory{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateProductAttributeCategoryField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateProductAttributeCategoryField...)
	return nil
}

// Create 创建产品属性分类表
func (r ProductAttributeCategoryRepo) Create(ctx context.Context, productAttributeCategory *entity.ProductAttributeCategory) error {
	if productAttributeCategory.ID > 0 {
		return errors.New("illegal argument productAttributeCategory id exist")
	}
	return r.GenericDao.Create(ctx, productAttributeCategory)
}

// DeleteByID 根据主键ID删除产品属性分类表
func (r ProductAttributeCategoryRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改产品属性分类表
func (r ProductAttributeCategoryRepo) Update(ctx context.Context, productAttributeCategory *entity.ProductAttributeCategory) error {
	if productAttributeCategory.ID == 0 {
		return errors.New("illegal argument productAttributeCategory exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateProductAttributeCategoryField).Updates(productAttributeCategory).Error
}

// GetByID 根据主键ID查询产品属性分类表
func (r ProductAttributeCategoryRepo) GetByID(ctx context.Context, id uint64) (*entity.ProductAttributeCategory, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByDBOption 根据动态条件查询产品属性分类表
func (r ProductAttributeCategoryRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db2.DBOption) ([]*entity.ProductAttributeCategory, uint32, error) {
	var (
		res       = make([]*entity.ProductAttributeCategory, 0)
		pageTotal = int64(0)
		offset    = (pageNum - 1) * pageSize
	)

	db := r.GenericDao.DB.WithContext(ctx)
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

func (r ProductAttributeCategoryRepo) WithByName(name string) db2.DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name like ?", "%"+name+"%")
	}
}
