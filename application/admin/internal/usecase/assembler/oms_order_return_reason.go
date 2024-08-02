package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// OrderReturnReasonEntityToModel entity转pb
func OrderReturnReasonEntityToModel(orderReturnReason *entity.OrderReturnReason) *pb.OrderReturnReason {
	return &pb.OrderReturnReason{
		Id:         orderReturnReason.ID,
		Name:       orderReturnReason.Name,
		Sort:       orderReturnReason.Sort,
		Status:     uint32(orderReturnReason.Status),
		CreateTime: orderReturnReason.CreateTime,
	}
}

// AddOrUpdateOrderReturnReasonParamToEntity pb转entity
func AddOrUpdateOrderReturnReasonParamToEntity(param *pb.AddOrUpdateOrderReturnReasonParam) *entity.OrderReturnReason {
	return &entity.OrderReturnReason{
		Name:   param.Name,
		Sort:   param.Sort,
		Status: uint8(param.Status),
	}
}
