package usecase

import (
	"context"

	"github.com/ccjshop/go-mall/application/admin/internal/usecase/assembler"
	"github.com/ccjshop/go-mall/common/db"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// OrderReturnApplyUseCase 订单退货申请管理Service实现类
type OrderReturnApplyUseCase struct {
	orderReturnApplyRepo IOrderReturnApplyRepo // 操作订单退货申请
	companyAddressRepo   ICompanyAddressRepo   // 公司收发货地址
}

// NewOrderReturnApply 创建订单退货申请管理Service实现类
func NewOrderReturnApply(orderReturnApplyRepo IOrderReturnApplyRepo, companyAddressRepo ICompanyAddressRepo) *OrderReturnApplyUseCase {
	return &OrderReturnApplyUseCase{
		orderReturnApplyRepo: orderReturnApplyRepo,
		companyAddressRepo:   companyAddressRepo,
	}
}

// GetOrderReturnApplies 分页查询订单退货申请
func (c OrderReturnApplyUseCase) GetOrderReturnApplies(ctx context.Context, param *pb.GetOrderReturnAppliesParam) ([]*pb.OrderReturnApply, uint32, error) {
	opts := make([]db.DBOption, 0)
	if param.GetId() != nil {
		opts = append(opts, c.orderReturnApplyRepo.WithByID(param.GetId().GetValue()))
	}
	if param.GetStatus() != nil {
		opts = append(opts, c.orderReturnApplyRepo.WithByStatus(uint8(param.GetStatus().GetValue())))
	}

	orderReturnApplies, pageTotal, err := c.orderReturnApplyRepo.GetByDBOption(ctx, param.GetPageNum(), param.GetPageSize(), opts...)
	if err != nil {
		return nil, 0, err
	}

	results := make([]*pb.OrderReturnApply, 0)
	for _, orderReturnApply := range orderReturnApplies {
		results = append(results, assembler.OrderReturnApplyEntityToModel(orderReturnApply))
	}
	return results, pageTotal, nil
}

// GetOrderReturnApply 根据id获取订单退货申请
func (c OrderReturnApplyUseCase) GetOrderReturnApply(ctx context.Context, id uint64) (*pb.OrderReturnApply, error) {
	orderReturnApply, err := c.orderReturnApplyRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// 补充公司收发货地址
	var companyAddressPb *pb.CompanyAddress
	if orderReturnApply.CompanyAddressID != 0 {
		companyAddress, err := c.companyAddressRepo.GetByID(ctx, orderReturnApply.CompanyAddressID)
		if err != nil {
			return nil, err
		}
		companyAddressPb = assembler.CompanyAddressEntityToModel(companyAddress)
	}

	res := assembler.OrderReturnApplyEntityToModel(orderReturnApply)
	res.CompanyAddress = companyAddressPb
	return res, nil
}

// DeleteOrderReturnApply 删除订单退货申请
func (c OrderReturnApplyUseCase) DeleteOrderReturnApply(ctx context.Context, id uint64) error {
	return c.orderReturnApplyRepo.DeleteByID(ctx, id)
}
