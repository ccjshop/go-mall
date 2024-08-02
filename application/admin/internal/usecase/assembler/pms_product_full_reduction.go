package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// ProductFullReductionsToModel entity转pb
func ProductFullReductionsToModel(productFullReductions []*entity.ProductFullReduction) []*pb.ProductFullReduction {
	res := make([]*pb.ProductFullReduction, 0)
	for _, productFullReduction := range productFullReductions {
		res = append(res, ProductFullReductionToModel(productFullReduction))
	}
	return res
}

// ProductFullReductionToModel entity转pb
func ProductFullReductionToModel(productFullReduction *entity.ProductFullReduction) *pb.ProductFullReduction {
	return &pb.ProductFullReduction{
		Id:          productFullReduction.ID,
		ProductId:   productFullReduction.ProductID,
		FullPrice:   productFullReduction.FullPrice.String(),
		ReducePrice: productFullReduction.ReducePrice.String(),
	}
}

// ProductFullReductionsToEntity pb转entity
func ProductFullReductionsToEntity(productFullReductionPbs []*pb.ProductFullReduction) []*entity.ProductFullReduction {
	res := make([]*entity.ProductFullReduction, 0)
	for _, productFullReductionPb := range productFullReductionPbs {
		res = append(res, ProductFullReductionToEntity(productFullReductionPb))
	}
	return res
}

// ProductFullReductionToEntity pb转entity
func ProductFullReductionToEntity(productFullReductionPb *pb.ProductFullReduction) *entity.ProductFullReduction {
	return &entity.ProductFullReduction{
		ID:          productFullReductionPb.Id,
		ProductID:   productFullReductionPb.ProductId,
		FullPrice:   util.DecimalUtils.ToDecimalFixed2(productFullReductionPb.FullPrice),
		ReducePrice: util.DecimalUtils.ToDecimalFixed2(productFullReductionPb.ReducePrice),
	}
}
