package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// ProductLadderEntityToModel entity转pb
func ProductLadderEntityToModel(productLadder *entity.ProductLadder) *pb.ProductLadder {
	return &pb.ProductLadder{}
}
