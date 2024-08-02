package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// OrderEntityToModel entity转pb
func OrderEntityToModel(order *entity.Order) *pb.Order {
	return &pb.Order{}
}
