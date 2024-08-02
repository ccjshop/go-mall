package usecase

import (
	"context"

	"github.com/ccjshop/go-mall/application/admin/internal/usecase/assembler"
	"github.com/ccjshop/go-mall/common/db"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// BrandUseCase 商品品牌表管理Service实现类
type BrandUseCase struct {
	brandRepo IBrandRepo // 操作商品品牌表
}

// NewBrandUseCase 创建商品品牌
func NewBrandUseCase(brandRepo IBrandRepo) *BrandUseCase {
	return &BrandUseCase{
		brandRepo: brandRepo,
	}
}

// CreateBrand 添加商品品牌
func (c BrandUseCase) CreateBrand(ctx context.Context, param *pb.AddOrUpdateBrandParam) error {
	// 数据转换
	brand := assembler.AddOrUpdateBrandParamToEntity(param)

	// 保存
	if err := c.brandRepo.Create(ctx, brand); err != nil {
		return err
	}

	return nil
}

// UpdateBrand 修改商品品牌
func (c BrandUseCase) UpdateBrand(ctx context.Context, param *pb.AddOrUpdateBrandParam) error {
	oldBrand, err := c.brandRepo.GetByID(ctx, param.GetId())
	if err != nil {
		return err
	}

	// 数据转换
	newBrand := assembler.AddOrUpdateBrandParamToEntity(param)
	newBrand.ID = param.Id
	newBrand.CreatedAt = oldBrand.CreatedAt

	// 更新商品品牌
	return c.brandRepo.Update(ctx, newBrand)
}

// GetBrands 分页查询商品品牌
func (c BrandUseCase) GetBrands(ctx context.Context, param *pb.GetBrandsParam) ([]*pb.Brand, uint32, error) {
	opts := make([]db.DBOption, 0)
	if len(param.GetName()) != 0 {
		opts = append(opts, c.brandRepo.WithByName(param.GetName()))
	}
	if param.GetShowStatus() != nil {
		opts = append(opts, c.brandRepo.WithByShowStatus(uint8(param.GetShowStatus().GetValue())))
	}

	brands, pageTotal, err := c.brandRepo.GetByDBOption(ctx, param.GetPageNum(), param.GetPageSize(), opts...)
	if err != nil {
		return nil, 0, err
	}

	results := make([]*pb.Brand, 0)
	for _, brand := range brands {
		results = append(results, assembler.BrandEntityToModel(brand))
	}
	return results, pageTotal, nil
}

// GetBrand 根据id获取商品品牌
func (c BrandUseCase) GetBrand(ctx context.Context, id uint64) (*pb.Brand, error) {
	brand, err := c.brandRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return assembler.BrandEntityToModel(brand), nil
}

// DeleteBrand 删除商品品牌表
func (c BrandUseCase) DeleteBrand(ctx context.Context, id uint64) error {
	return c.brandRepo.DeleteByID(ctx, id)
}
