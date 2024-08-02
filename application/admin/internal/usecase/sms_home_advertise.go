package usecase

import (
	"context"

	"github.com/ccjshop/go-mall/application/admin/internal/usecase/assembler"
	"github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// HomeAdvertiseUseCase 首页轮播广告表管理Service实现类
type HomeAdvertiseUseCase struct {
	homeAdvertiseRepo IHomeAdvertiseRepo // 操作首页轮播广告表
}

// NewHomeAdvertise 创建首页轮播广告表管理Service实现类
func NewHomeAdvertise(homeAdvertiseRepo IHomeAdvertiseRepo) *HomeAdvertiseUseCase {
	return &HomeAdvertiseUseCase{
		homeAdvertiseRepo: homeAdvertiseRepo,
	}
}

// CreateHomeAdvertise 添加首页轮播广告表
func (c HomeAdvertiseUseCase) CreateHomeAdvertise(ctx context.Context, param *pb.AddOrUpdateHomeAdvertiseParam) error {
	// 数据转换
	homeAdvertise := assembler.AddOrUpdateHomeAdvertiseParamToEntity(param)

	// 保存
	if err := c.homeAdvertiseRepo.Create(ctx, homeAdvertise); err != nil {
		return err
	}

	return nil
}

// UpdateHomeAdvertise 修改首页轮播广告表
func (c HomeAdvertiseUseCase) UpdateHomeAdvertise(ctx context.Context, param *pb.AddOrUpdateHomeAdvertiseParam) error {
	var (
		oldHomeAdvertise *entity.HomeAdvertise
		newHomeAdvertise *entity.HomeAdvertise
		err              error
	)

	// 老数据
	if oldHomeAdvertise, err = c.homeAdvertiseRepo.GetByID(ctx, param.GetId()); err != nil {
		return err
	}

	// 新数据
	newHomeAdvertise = assembler.AddOrUpdateHomeAdvertiseParamToEntity(param)
	newHomeAdvertise.ID = param.Id
	newHomeAdvertise.CreatedAt = oldHomeAdvertise.CreatedAt

	// 更新首页轮播广告表
	return c.homeAdvertiseRepo.Update(ctx, newHomeAdvertise)
}

// GetHomeAdvertises 分页查询首页轮播广告表
func (c HomeAdvertiseUseCase) GetHomeAdvertises(ctx context.Context, param *pb.GetHomeAdvertisesParam) ([]*pb.HomeAdvertise, uint32, error) {
	opts := make([]db.DBOption, 0)

	homeAdvertises, pageTotal, err := c.homeAdvertiseRepo.GetByDBOption(ctx, param.GetPageNum(), param.GetPageSize(), opts...)
	if err != nil {
		return nil, 0, err
	}

	results := make([]*pb.HomeAdvertise, 0)
	for _, homeAdvertise := range homeAdvertises {
		results = append(results, assembler.HomeAdvertiseEntityToModel(homeAdvertise))
	}
	return results, pageTotal, nil
}

// GetHomeAdvertise 根据id获取首页轮播广告表
func (c HomeAdvertiseUseCase) GetHomeAdvertise(ctx context.Context, id uint64) (*pb.HomeAdvertise, error) {
	homeAdvertise, err := c.homeAdvertiseRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return assembler.HomeAdvertiseEntityToModel(homeAdvertise), nil
}

// DeleteHomeAdvertise 删除首页轮播广告表
func (c HomeAdvertiseUseCase) DeleteHomeAdvertise(ctx context.Context, id uint64) error {
	return c.homeAdvertiseRepo.DeleteByID(ctx, id)
}
