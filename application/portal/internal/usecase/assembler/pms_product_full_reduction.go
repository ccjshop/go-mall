package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// ProductFullReductionEntityToModel entity转pb
func ProductFullReductionEntityToModel(productFullReduction *entity.ProductFullReduction) *pb.ProductFullReduction {
	return &pb.ProductFullReduction{}
}
