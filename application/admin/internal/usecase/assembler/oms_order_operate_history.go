package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// OrderOperateHistoriesToModel entity转pb
func OrderOperateHistoriesToModel(orderOperateHistories []*entity.OrderOperateHistory) []*pb.OrderOperateHistory {
	res := make([]*pb.OrderOperateHistory, 0)
	for _, orderOperateHistory := range orderOperateHistories {
		res = append(res, OrderOperateHistoryToModel(orderOperateHistory))
	}
	return res
}

// OrderOperateHistoryToModel entity转pb
func OrderOperateHistoryToModel(orderOperateHistory *entity.OrderOperateHistory) *pb.OrderOperateHistory {
	return &pb.OrderOperateHistory{
		Id:          orderOperateHistory.ID,
		OrderId:     orderOperateHistory.OrderID,
		OperateMan:  orderOperateHistory.OperateMan,
		OrderStatus: uint32(orderOperateHistory.OrderStatus),
		Note:        orderOperateHistory.Note,
		CreateTime:  orderOperateHistory.CreateTime,
	}
}
