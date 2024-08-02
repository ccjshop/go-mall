package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// OrderItemEntityToModel entity转pb
func OrderItemEntityToModel(orderItem *entity.OrderItem) *pb.OrderItem {
	return &pb.OrderItem{}
}
