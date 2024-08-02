package repo

import (
	"context"

	"github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
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

// GetProductAttributeByCategoryID 根据产品属性分类表ID获取商品属性参数表
func (r ProductAttributeRepo) GetProductAttributeByCategoryID(ctx context.Context, productAttributeCategoryID uint64) (entity.ProductAttributes, error) {
	res := make([]*entity.ProductAttribute, 0)
	if err := r.GenericDao.DB.WithContext(ctx).
		Where("product_attribute_category_id = ?", productAttributeCategoryID).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
