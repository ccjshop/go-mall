package repo

import (
	"context"

	db2 "github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"gorm.io/gorm"
)

// ProductAttributeValueRepo 产品参数信息
type ProductAttributeValueRepo struct {
	*db2.GenericDao[entity.ProductAttributeValue, uint64]
}

// NewProductAttributeValueRepo 创建
func NewProductAttributeValueRepo(conn *gorm.DB) *ProductAttributeValueRepo {
	return &ProductAttributeValueRepo{
		GenericDao: &db2.GenericDao[entity.ProductAttributeValue, uint64]{
			DB: conn,
		},
	}
}

// GetByProductAttributeID 根据productAttributeID查询产品参数信息
func (r ProductAttributeValueRepo) GetByProductAttributeID(ctx context.Context, productID uint64, productAttributeIDs []uint64) (entity.ProductAttributeValues, error) {
	res := make([]*entity.ProductAttributeValue, 0)
	if err := r.GenericDao.DB.WithContext(ctx).
		Where("product_id = ?", productID).
		Where("product_attribute_id in ?", productAttributeIDs).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
