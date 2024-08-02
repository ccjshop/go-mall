package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// ProductLaddersToModel entity转pb
func ProductLaddersToModel(productLadders []*entity.ProductLadder) []*pb.ProductLadder {
	res := make([]*pb.ProductLadder, 0)
	for _, productLadder := range productLadders {
		res = append(res, ProductLadderToModel(productLadder))
	}
	return res
}

// ProductLadderToModel entity转pb
func ProductLadderToModel(productLadder *entity.ProductLadder) *pb.ProductLadder {
	return &pb.ProductLadder{
		Id:        productLadder.ID,
		ProductId: productLadder.ProductID,
		Count:     productLadder.Count,
		Discount:  productLadder.Discount.String(),
		Price:     productLadder.Price.String(),
	}
}

// ProductLaddersToEntity pb转entity
func ProductLaddersToEntity(productLadderPbs []*pb.ProductLadder) []*entity.ProductLadder {
	res := make([]*entity.ProductLadder, 0)
	for _, productLadderPb := range productLadderPbs {
		res = append(res, ProductLadderToEntity(productLadderPb))
	}
	return res
}

// ProductLadderToEntity pb转entity
func ProductLadderToEntity(productLadderPb *pb.ProductLadder) *entity.ProductLadder {
	return &entity.ProductLadder{
		ID:        productLadderPb.Id,
		ProductID: productLadderPb.ProductId,
		Count:     productLadderPb.Count,
		Discount:  util.DecimalUtils.ToDecimalFixed2(productLadderPb.Discount),
		Price:     util.DecimalUtils.ToDecimalFixed2(productLadderPb.Price),
	}
}
