package usecase

import (
	"context"

	"github.com/ccjshop/go-mall/application/admin/internal/usecase/assembler"
	"github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// CompanyAddressUseCase 公司收发货地址管理Service实现类
type CompanyAddressUseCase struct {
	companyAddressRepo ICompanyAddressRepo // 操作公司收发货地址
}

// NewCompanyAddress 创建公司收发货地址管理Service实现类
func NewCompanyAddress(companyAddressRepo ICompanyAddressRepo) *CompanyAddressUseCase {
	return &CompanyAddressUseCase{
		companyAddressRepo: companyAddressRepo,
	}
}

// CreateCompanyAddress 添加公司收发货地址
func (c CompanyAddressUseCase) CreateCompanyAddress(ctx context.Context, param *pb.AddOrUpdateCompanyAddressParam) error {
	// 数据转换
	companyAddress := assembler.AddOrUpdateCompanyAddressParamToEntity(param)

	// 保存
	if err := c.companyAddressRepo.Create(ctx, companyAddress); err != nil {
		return err
	}

	return nil
}

// UpdateCompanyAddress 修改公司收发货地址
func (c CompanyAddressUseCase) UpdateCompanyAddress(ctx context.Context, param *pb.AddOrUpdateCompanyAddressParam) error {
	var (
		oldCompanyAddress *entity.CompanyAddress
		newCompanyAddress *entity.CompanyAddress
		err               error
	)

	// 老数据
	if oldCompanyAddress, err = c.companyAddressRepo.GetByID(ctx, param.GetId()); err != nil {
		return err
	}

	// 新数据
	newCompanyAddress = assembler.AddOrUpdateCompanyAddressParamToEntity(param)
	newCompanyAddress.ID = param.Id
	newCompanyAddress.CreatedAt = oldCompanyAddress.CreatedAt

	// 更新公司收发货地址
	return c.companyAddressRepo.Update(ctx, newCompanyAddress)
}

// GetCompanyAddresses 分页查询公司收发货地址
func (c CompanyAddressUseCase) GetCompanyAddresses(ctx context.Context, param *pb.GetCompanyAddressesParam) ([]*pb.CompanyAddress, uint32, error) {
	opts := make([]db.DBOption, 0)
	companyAddresses, pageTotal, err := c.companyAddressRepo.GetByDBOption(ctx, param.GetPageNum(), param.GetPageSize(), opts...)
	if err != nil {
		return nil, 0, err
	}

	results := make([]*pb.CompanyAddress, 0)
	for _, companyAddress := range companyAddresses {
		results = append(results, assembler.CompanyAddressEntityToModel(companyAddress))
	}
	return results, pageTotal, nil
}

// GetCompanyAddress 根据id获取公司收发货地址
func (c CompanyAddressUseCase) GetCompanyAddress(ctx context.Context, id uint64) (*pb.CompanyAddress, error) {
	companyAddress, err := c.companyAddressRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return assembler.CompanyAddressEntityToModel(companyAddress), nil
}

// DeleteCompanyAddress 删除公司收发货地址
func (c CompanyAddressUseCase) DeleteCompanyAddress(ctx context.Context, id uint64) error {
	return c.companyAddressRepo.DeleteByID(ctx, id)
}
