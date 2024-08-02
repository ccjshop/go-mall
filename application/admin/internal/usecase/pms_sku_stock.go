package usecase

import (
	"context"

	"github.com/ccjshop/go-mall/application/admin/internal/usecase/assembler"
	"github.com/ccjshop/go-mall/application/admin/pkg/constant"
	"github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// SkuStockUseCase sku的库存管理Service实现类
type SkuStockUseCase struct {
	skuStockRepo ISkuStockRepo // 操作sku的库存
}

// NewSkuStock 创建sku的库存管理Service实现类
func NewSkuStock(skuStockRepo ISkuStockRepo) *SkuStockUseCase {
	return &SkuStockUseCase{
		skuStockRepo: skuStockRepo,
	}
}

// BatchUpdateSkuStock 批量更新sku的库存
func (c SkuStockUseCase) BatchUpdateSkuStock(ctx context.Context, param *pb.BatchUpdateSkuStockParam) error {
	// 过滤掉错误的数据
	var skuStocks []*entity.SkuStock
	for _, skuStock := range param.SkuStocks {
		if skuStock.GetProductId() == param.ProductId {
			skuStocks = append(skuStocks, assembler.SkuStockToEntity(skuStock))
		}
	}
	// 批量插入或者修改
	return c.skuStockRepo.BatchUpdateOrInsertSkuStock(ctx, skuStocks)
}

// GetSkuStocksByProductID 根据商品id分页查询sku的库存
func (c SkuStockUseCase) GetSkuStocksByProductID(ctx context.Context, param *pb.GetSkuStocksByProductIdParam) ([]*pb.SkuStock, error) {
	opts := make([]db.DBOption, 0)
	opts = append(opts, c.skuStockRepo.WithByProductID(param.GetProductId()))
	if len(param.GetSkuCode()) != 0 {
		opts = append(opts, c.skuStockRepo.WithBySkuCode(param.GetSkuCode()))
	}
	skuStocks, _, err := c.skuStockRepo.GetByDBOption(ctx, 1, constant.FindAllCountLimit, opts...)
	if err != nil {
		return nil, err
	}

	return assembler.SkuStocksToModel(skuStocks), nil
}
