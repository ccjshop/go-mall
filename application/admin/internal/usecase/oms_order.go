package usecase

import (
	"context"

	"github.com/ccjshop/go-mall/application/admin/internal/usecase/assembler"
	"github.com/ccjshop/go-mall/common/db"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// OrderUseCase 订单管理Service实现类
type OrderUseCase struct {
	orderRepo               IOrderRepo               // 操作订单
	orderItemRepo           IOrderItemRepo           // 订单商品信息
	orderOperateHistoryRepo IOrderOperateHistoryRepo // 订单操作历史记录
}

// NewOrder 创建订单管理Service实现类
func NewOrder(orderRepo IOrderRepo, orderItemRepo IOrderItemRepo, orderOperateHistoryRepo IOrderOperateHistoryRepo) *OrderUseCase {
	return &OrderUseCase{
		orderRepo:               orderRepo,
		orderItemRepo:           orderItemRepo,
		orderOperateHistoryRepo: orderOperateHistoryRepo,
	}
}

// GetOrders 分页查询订单
func (c OrderUseCase) GetOrders(ctx context.Context, param *pb.GetOrdersParam) ([]*pb.Order, uint32, error) {
	opts := make([]db.DBOption, 0)
	if param.GetId() != nil {
		opts = append(opts, c.orderRepo.WithByID(param.GetId().GetValue()))
	}

	orders, pageTotal, err := c.orderRepo.GetByDBOption(ctx, param.GetPageNum(), param.GetPageSize(), opts...)
	if err != nil {
		return nil, 0, err
	}

	results := make([]*pb.Order, 0)
	for _, order := range orders {
		results = append(results, assembler.OrderEntityToModel(order))
	}
	return results, pageTotal, nil
}

// GetOrder 根据id获取订单
func (c OrderUseCase) GetOrder(ctx context.Context, orderID uint64) (*pb.Order, error) {
	order, err := c.orderRepo.GetByID(ctx, orderID)
	if err != nil {
		return nil, err
	}
	// 补充订单商品列表
	orderItems, err := c.orderItemRepo.GetByOrderID(ctx, orderID)
	if err != nil {
		return nil, err
	}
	// 补充订单操作记录列表
	orderOperateHistories, err := c.orderOperateHistoryRepo.GetByOrderID(ctx, orderID)
	if err != nil {
		return nil, err
	}

	res := assembler.OrderEntityToModel(order)
	res.OrderOperateHistories = assembler.OrderOperateHistoriesToModel(orderOperateHistories)
	res.OrderItems = assembler.OrderItemsToModel(orderItems)

	return res, nil
}

// DeleteOrder 删除订单
func (c OrderUseCase) DeleteOrder(ctx context.Context, id uint64) error {
	return c.orderRepo.DeleteByID(ctx, id)
}
