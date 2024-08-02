package usecase

import (
	"context"

	"github.com/ccjshop/go-mall/application/admin/internal/usecase/assembler"
	"github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// OrderReturnReasonUseCase 退货原因管理Service实现类
type OrderReturnReasonUseCase struct {
	orderReturnReasonRepo IOrderReturnReasonRepo // 操作退货原因
}

// NewOrderReturnReason 创建退货原因管理Service实现类
func NewOrderReturnReason(orderReturnReasonRepo IOrderReturnReasonRepo) *OrderReturnReasonUseCase {
	return &OrderReturnReasonUseCase{
		orderReturnReasonRepo: orderReturnReasonRepo,
	}
}

// CreateOrderReturnReason 添加退货原因
func (c OrderReturnReasonUseCase) CreateOrderReturnReason(ctx context.Context, param *pb.AddOrUpdateOrderReturnReasonParam) error {
	// 数据转换
	orderReturnReason := assembler.AddOrUpdateOrderReturnReasonParamToEntity(param)

	// 保存
	if err := c.orderReturnReasonRepo.Create(ctx, orderReturnReason); err != nil {
		return err
	}

	return nil
}

// UpdateOrderReturnReason 修改退货原因
func (c OrderReturnReasonUseCase) UpdateOrderReturnReason(ctx context.Context, param *pb.AddOrUpdateOrderReturnReasonParam) error {
	var (
		oldOrderReturnReason *entity.OrderReturnReason
		newOrderReturnReason *entity.OrderReturnReason
		err                  error
	)

	// 老数据
	if oldOrderReturnReason, err = c.orderReturnReasonRepo.GetByID(ctx, param.GetId()); err != nil {
		return err
	}

	// 新数据
	newOrderReturnReason = assembler.AddOrUpdateOrderReturnReasonParamToEntity(param)
	newOrderReturnReason.ID = param.Id
	newOrderReturnReason.CreateTime = oldOrderReturnReason.CreateTime
	newOrderReturnReason.CreatedAt = oldOrderReturnReason.CreatedAt

	// 更新退货原因
	return c.orderReturnReasonRepo.Update(ctx, newOrderReturnReason)
}

// GetOrderReturnReasons 分页查询退货原因
func (c OrderReturnReasonUseCase) GetOrderReturnReasons(ctx context.Context, param *pb.GetOrderReturnReasonsParam) ([]*pb.OrderReturnReason, uint32, error) {
	opts := make([]db.DBOption, 0)
	if param.GetId() != nil {
		opts = append(opts, c.orderReturnReasonRepo.WithByID(param.GetId().GetValue()))
	}
	orderReturnReasons, pageTotal, err := c.orderReturnReasonRepo.GetByDBOption(ctx, param.GetPageNum(), param.GetPageSize(), opts...)
	if err != nil {
		return nil, 0, err
	}

	results := make([]*pb.OrderReturnReason, 0)
	for _, orderReturnReason := range orderReturnReasons {
		results = append(results, assembler.OrderReturnReasonEntityToModel(orderReturnReason))
	}
	return results, pageTotal, nil
}

// GetOrderReturnReason 根据id获取退货原因
func (c OrderReturnReasonUseCase) GetOrderReturnReason(ctx context.Context, id uint64) (*pb.OrderReturnReason, error) {
	orderReturnReason, err := c.orderReturnReasonRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return assembler.OrderReturnReasonEntityToModel(orderReturnReason), nil
}

// DeleteOrderReturnReason 删除退货原因
func (c OrderReturnReasonUseCase) DeleteOrderReturnReason(ctx context.Context, id uint64) error {
	return c.orderReturnReasonRepo.DeleteByID(ctx, id)
}
