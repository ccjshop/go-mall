package usecase

import (
	"context"

	"github.com/ccjshop/go-mall/application/admin/internal/usecase/assembler"
	"github.com/ccjshop/go-mall/common/db"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// ProductAttributeUseCase 商品属性参数管理Service实现类
type ProductAttributeUseCase struct {
	productAttributeRepo IProductAttributeRepo // 操作商品属性参数表
}

// NewProductAttribute 创建商品属性参数管理Service实现类
func NewProductAttribute(productAttributeRepo IProductAttributeRepo) *ProductAttributeUseCase {
	return &ProductAttributeUseCase{
		productAttributeRepo: productAttributeRepo,
	}
}

// CreateProductAttribute 添加商品属性参数
func (c ProductAttributeUseCase) CreateProductAttribute(ctx context.Context, param *pb.AddOrUpdateProductAttributeParam) error {
	// 数据转换
	productAttribute := assembler.AddOrUpdateProductAttributeParamToEntity(param)

	// 保存
	if err := c.productAttributeRepo.Create(ctx, productAttribute); err != nil {
		return err
	}

	return nil
}

// UpdateProductAttribute 修改商品属性参数
func (c ProductAttributeUseCase) UpdateProductAttribute(ctx context.Context, param *pb.AddOrUpdateProductAttributeParam) error {
	oldProductAttribute, err := c.productAttributeRepo.GetByID(ctx, param.GetId())
	if err != nil {
		return err
	}

	// 数据转换
	productAttribute := assembler.AddOrUpdateProductAttributeParamToEntity(param)
	productAttribute.ID = param.Id
	productAttribute.CreatedAt = oldProductAttribute.CreatedAt

	// 更新商品属性参数
	return c.productAttributeRepo.Update(ctx, productAttribute)
}

// GetProductAttributes 分页查询商品属性参数
func (c ProductAttributeUseCase) GetProductAttributes(ctx context.Context, param *pb.GetProductAttributesParam) ([]*pb.ProductAttribute, uint32, error) {
	opts := make([]db.DBOption, 0)
	if len(param.GetName()) != 0 {
		opts = append(opts, c.productAttributeRepo.WithByName(param.GetName()))
	}
	if param.GetProductAttributeCategoryId() != nil {
		opts = append(opts, c.productAttributeRepo.WithByProductAttributeCategoryID(param.GetProductAttributeCategoryId().GetValue()))
	}
	if param.GetType() != nil {
		opts = append(opts, c.productAttributeRepo.WithByType(param.GetType().GetValue()))
	}

	productAttributes, pageTotal, err := c.productAttributeRepo.GetByDBOption(ctx, param.GetPageNum(), param.GetPageSize(), opts...)
	if err != nil {
		return nil, 0, err
	}

	results := make([]*pb.ProductAttribute, 0)
	for _, productAttribute := range productAttributes {
		results = append(results, assembler.ProductAttributeEntityToModel(productAttribute))
	}
	return results, pageTotal, nil
}

// GetProductAttribute 根据id获取商品属性参数
func (c ProductAttributeUseCase) GetProductAttribute(ctx context.Context, id uint64) (*pb.ProductAttribute, error) {
	productAttribute, err := c.productAttributeRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return assembler.ProductAttributeEntityToModel(productAttribute), nil
}

// DeleteProductAttribute 删除商品属性参数
func (c ProductAttributeUseCase) DeleteProductAttribute(ctx context.Context, id uint64) error {
	return c.productAttributeRepo.DeleteByID(ctx, id)
}
