package grpcsrv

import (
	"context"

	"github.com/ccjshop/go-mall/common/retcode"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// BatchUpdateSkuStock 批量添加sku的库存
func (s *AdminApiImpl) BatchUpdateSkuStock(ctx context.Context, param *pb.BatchUpdateSkuStockParam) (*pb.CommonRsp, error) {
	var (
		res = &pb.CommonRsp{}
	)
	if err := s.skuStock.BatchUpdateSkuStock(ctx, param); err != nil {
		return nil, err
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// GetSkuStocksByProductId 根据商品id分页查询sku的库存
func (s *AdminApiImpl) GetSkuStocksByProductId(ctx context.Context, param *pb.GetSkuStocksByProductIdParam) (*pb.GetSkuStocksByProductIdRsp, error) {
	var (
		res = &pb.GetSkuStocksByProductIdRsp{}
	)

	skuStocks, err := s.skuStock.GetSkuStocksByProductID(ctx, param)
	if err != nil {
		return nil, err
	}
	res.Data = &pb.SkuStocksData{
		Data: skuStocks,
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}
