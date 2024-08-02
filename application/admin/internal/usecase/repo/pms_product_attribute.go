package repo

import (
	"context"
	"errors"

	"github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	"gorm.io/gorm"
)

// ProductAttributeRepo 商品属性参数表
type ProductAttributeRepo struct {
	*db.GenericDao[entity.ProductAttribute, uint64]
}

// NewProductAttributeRepo 创建
func NewProductAttributeRepo(conn *gorm.DB) *ProductAttributeRepo {
	return &ProductAttributeRepo{
		GenericDao: &db.GenericDao[entity.ProductAttribute, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initProductAttributeField)
}

var (
	// 全字段修改ProductAttribute那些字段不修改
	notUpdateProductAttributeField = []string{
		"created_at",
	}
	updateProductAttributeField []string
)

// initProductAttributeField 全字段修改
func initProductAttributeField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.ProductAttribute{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateProductAttributeField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateProductAttributeField...)
	return nil
}

// Create 创建商品属性参数表
func (r ProductAttributeRepo) Create(ctx context.Context, productAttribute *entity.ProductAttribute) error {
	if productAttribute.ID > 0 {
		return errors.New("illegal argument productAttribute id exist")
	}
	return r.GenericDao.Create(ctx, productAttribute)
}

// DeleteByID 根据主键ID删除商品属性参数表
func (r ProductAttributeRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改商品属性参数表
func (r ProductAttributeRepo) Update(ctx context.Context, productAttribute *entity.ProductAttribute) error {
	if productAttribute.ID == 0 {
		return errors.New("illegal argument productAttribute exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateProductAttributeField).Updates(productAttribute).Error
}

// GetByID 根据主键ID查询商品属性参数表
func (r ProductAttributeRepo) GetByID(ctx context.Context, id uint64) (*entity.ProductAttribute, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByDBOption 根据动态条件查询商品属性参数表
func (r ProductAttributeRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.ProductAttribute, uint32, error) {
	var (
		res       = make([]*entity.ProductAttribute, 0)
		pageTotal = int64(0)
		offset    = (pageNum - 1) * pageSize
	)

	session := r.GenericDao.DB.WithContext(ctx)
	for _, opt := range opts {
		session = opt(session)
	}

	session = session.Offset(int(offset)).Limit(int(pageSize)).Order("id asc").Find(&res).
		Offset(-1).Limit(-1).Count(&pageTotal)

	if err := session.Error; err != nil {
		return nil, 0, err
	}
	return res, uint32(pageTotal), nil
}

func (r ProductAttributeRepo) WithByName(name string) db.DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name like ?", "%"+name+"%")
	}
}

func (r ProductAttributeRepo) WithByProductAttributeCategoryID(productAttributeCategoryID uint32) db.DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("product_attribute_category_id = ?", productAttributeCategoryID)
	}
}
func (r ProductAttributeRepo) WithByType(tpe uint32) db.DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("type = ?", tpe)
	}
}
