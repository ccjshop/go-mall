package usecase

import (
	"context"

	"github.com/ccjshop/go-mall/application/admin/internal/usecase/assembler"
	"github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// PrefrenceAreaUseCase 优选专区管理Service实现类
type PrefrenceAreaUseCase struct {
	prefrenceAreaRepo IPrefrenceAreaRepo // 操作优选专区
}

// NewPrefrenceArea 创建优选专区管理Service实现类
func NewPrefrenceArea(prefrenceAreaRepo IPrefrenceAreaRepo) *PrefrenceAreaUseCase {
	return &PrefrenceAreaUseCase{
		prefrenceAreaRepo: prefrenceAreaRepo,
	}
}

// CreatePrefrenceArea 添加优选专区
func (c PrefrenceAreaUseCase) CreatePrefrenceArea(ctx context.Context, param *pb.AddOrUpdatePrefrenceAreaParam) error {
	// 数据转换
	prefrenceArea := assembler.AddOrUpdatePrefrenceAreaParamToEntity(param)

	// 保存
	if err := c.prefrenceAreaRepo.Create(ctx, prefrenceArea); err != nil {
		return err
	}

	return nil
}

// UpdatePrefrenceArea 修改优选专区
func (c PrefrenceAreaUseCase) UpdatePrefrenceArea(ctx context.Context, param *pb.AddOrUpdatePrefrenceAreaParam) error {
	var (
		oldPrefrenceArea *entity.PrefrenceArea
		newPrefrenceArea *entity.PrefrenceArea
		err              error
	)

	// 老数据
	if oldPrefrenceArea, err = c.prefrenceAreaRepo.GetByID(ctx, param.GetId()); err != nil {
		return err
	}

	// 新数据
	newPrefrenceArea = assembler.AddOrUpdatePrefrenceAreaParamToEntity(param)
	newPrefrenceArea.ID = param.Id
	newPrefrenceArea.CreatedAt = oldPrefrenceArea.CreatedAt

	// 更新优选专区
	return c.prefrenceAreaRepo.Update(ctx, newPrefrenceArea)
}

// GetPrefrenceAreas 分页查询优选专区
func (c PrefrenceAreaUseCase) GetPrefrenceAreas(ctx context.Context, param *pb.GetPrefrenceAreasParam) ([]*pb.PrefrenceArea, uint32, error) {
	opts := make([]db.DBOption, 0)
	prefrenceAreas, pageTotal, err := c.prefrenceAreaRepo.GetByDBOption(ctx, param.GetPageNum(), param.GetPageSize(), opts...)
	if err != nil {
		return nil, 0, err
	}

	results := make([]*pb.PrefrenceArea, 0)
	for _, prefrenceArea := range prefrenceAreas {
		results = append(results, assembler.PrefrenceAreaEntityToModel(prefrenceArea))
	}
	return results, pageTotal, nil
}

// GetPrefrenceArea 根据id获取优选专区
func (c PrefrenceAreaUseCase) GetPrefrenceArea(ctx context.Context, id uint64) (*pb.PrefrenceArea, error) {
	prefrenceArea, err := c.prefrenceAreaRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return assembler.PrefrenceAreaEntityToModel(prefrenceArea), nil
}

// DeletePrefrenceArea 删除优选专区
func (c PrefrenceAreaUseCase) DeletePrefrenceArea(ctx context.Context, id uint64) error {
	return c.prefrenceAreaRepo.DeleteByID(ctx, id)
}
