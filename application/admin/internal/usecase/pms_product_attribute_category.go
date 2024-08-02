package usecase

import (
	"context"

	"github.com/ccjshop/go-mall/application/admin/internal/usecase/assembler"
	"github.com/ccjshop/go-mall/common/db"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// ProductAttributeCategoryUseCase 产品属性分类表管理Service实现类
type ProductAttributeCategoryUseCase struct {
	productAttributeCategoryRepo IProductAttributeCategoryRepo // 操作产品属性分类表
}

// NewProductAttributeCategory 创建产品属性分类表
func NewProductAttributeCategory(productAttributeCategoryRepo IProductAttributeCategoryRepo) *ProductAttributeCategoryUseCase {
	return &ProductAttributeCategoryUseCase{
		productAttributeCategoryRepo: productAttributeCategoryRepo,
	}
}

// CreateProductAttributeCategory 添加产品属性分类表
func (c ProductAttributeCategoryUseCase) CreateProductAttributeCategory(ctx context.Context, param *pb.AddOrUpdateProductAttributeCategoryParam) error {
	// 数据转换
	productAttributeCategory := assembler.AddOrUpdateProductAttributeCategoryParamToEntity(param)

	// 保存
	if err := c.productAttributeCategoryRepo.Create(ctx, productAttributeCategory); err != nil {
		return err
	}
	return nil
}

// UpdateProductAttributeCategory 修改产品属性分类
func (c ProductAttributeCategoryUseCase) UpdateProductAttributeCategory(ctx context.Context, param *pb.AddOrUpdateProductAttributeCategoryParam) error {
	oldProductAttributeCategory, err := c.productAttributeCategoryRepo.GetByID(ctx, param.GetId())
	if err != nil {
		return err
	}

	// 数据转换
	productAttributeCategory := assembler.AddOrUpdateProductAttributeCategoryParamToEntity(param)
	productAttributeCategory.ID = param.Id
	productAttributeCategory.CreatedAt = oldProductAttributeCategory.CreatedAt

	// 更新产品属性分类
	return c.productAttributeCategoryRepo.Update(ctx, productAttributeCategory)
}

// GetProductAttributeCategories 分页查询产品属性分类
func (c ProductAttributeCategoryUseCase) GetProductAttributeCategories(ctx context.Context, param *pb.GetProductAttributeCategoriesParam) ([]*pb.ProductAttributeCategory, uint32, error) {
	opts := make([]db.DBOption, 0)
	if len(param.GetName()) != 0 {
		opts = append(opts, c.productAttributeCategoryRepo.WithByName(param.GetName()))
	}
	productAttributeCategories, pageTotal, err := c.productAttributeCategoryRepo.GetByDBOption(ctx, param.GetPageNum(), param.GetPageSize(), opts...)
	if err != nil {
		return nil, 0, err
	}

	results := make([]*pb.ProductAttributeCategory, 0)
	for _, productAttributeCategory := range productAttributeCategories {
		results = append(results, assembler.ProductAttributeCategoryEntityToModel(productAttributeCategory))
	}
	return results, pageTotal, nil
}

// GetProductAttributeCategory 根据id获取产品属性分类
func (c ProductAttributeCategoryUseCase) GetProductAttributeCategory(ctx context.Context, id uint64) (*pb.ProductAttributeCategory, error) {
	productAttributeCategory, err := c.productAttributeCategoryRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return assembler.ProductAttributeCategoryEntityToModel(productAttributeCategory), nil
}

// DeleteProductAttributeCategory 删除产品属性分类
func (c ProductAttributeCategoryUseCase) DeleteProductAttributeCategory(ctx context.Context, id uint64) error {
	return c.productAttributeCategoryRepo.DeleteByID(ctx, id)
}
