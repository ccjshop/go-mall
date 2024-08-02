package usecase

import (
	"context"

	"github.com/ccjshop/go-mall/application/admin/internal/usecase/assembler"
	"github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// JsonDynamicConfigUseCase JSON动态配置管理Service实现类
type JsonDynamicConfigUseCase struct {
	jsonDynamicConfigRepo IJsonDynamicConfigRepo // 操作JSON动态配置
}

// NewJsonDynamicConfig 创建JSON动态配置管理Service实现类
func NewJsonDynamicConfig(jsonDynamicConfigRepo IJsonDynamicConfigRepo) *JsonDynamicConfigUseCase {
	return &JsonDynamicConfigUseCase{
		jsonDynamicConfigRepo: jsonDynamicConfigRepo,
	}
}

// CreateJsonDynamicConfig 添加JSON动态配置
func (c JsonDynamicConfigUseCase) CreateJsonDynamicConfig(ctx context.Context, param *pb.AddOrUpdateJsonDynamicConfigParam) error {
	// 数据转换
	jsonDynamicConfig := assembler.AddOrUpdateJsonDynamicConfigParamToEntity(param)

	// 保存
	if err := c.jsonDynamicConfigRepo.Create(ctx, jsonDynamicConfig); err != nil {
		return err
	}

	return nil
}

// UpdateJsonDynamicConfig 修改JSON动态配置
func (c JsonDynamicConfigUseCase) UpdateJsonDynamicConfig(ctx context.Context, param *pb.AddOrUpdateJsonDynamicConfigParam) error {
	var (
		oldJsonDynamicConfig *entity.JsonDynamicConfig
		newJsonDynamicConfig *entity.JsonDynamicConfig
		err                  error
	)

	// 老数据
	if oldJsonDynamicConfig, err = c.jsonDynamicConfigRepo.GetByID(ctx, param.GetId()); err != nil {
		return err
	}

	// 新数据
	newJsonDynamicConfig = assembler.AddOrUpdateJsonDynamicConfigParamToEntity(param)
	newJsonDynamicConfig.ID = param.Id
	newJsonDynamicConfig.CreatedAt = oldJsonDynamicConfig.CreatedAt

	// 更新JSON动态配置
	return c.jsonDynamicConfigRepo.Update(ctx, newJsonDynamicConfig)
}

// GetJsonDynamicConfigs 分页查询JSON动态配置
func (c JsonDynamicConfigUseCase) GetJsonDynamicConfigs(ctx context.Context, param *pb.GetJsonDynamicConfigsParam) ([]*pb.JsonDynamicConfig, uint32, error) {
	opts := make([]db.DBOption, 0)
	jsonDynamicConfigs, pageTotal, err := c.jsonDynamicConfigRepo.GetByDBOption(ctx, param.GetPageNum(), param.GetPageSize(), opts...)
	if err != nil {
		return nil, 0, err
	}

	results := make([]*pb.JsonDynamicConfig, 0)
	for _, jsonDynamicConfig := range jsonDynamicConfigs {
		results = append(results, assembler.JsonDynamicConfigEntityToModel(jsonDynamicConfig))
	}
	return results, pageTotal, nil
}

// GetJsonDynamicConfig 根据id获取JSON动态配置
func (c JsonDynamicConfigUseCase) GetJsonDynamicConfig(ctx context.Context, id uint64) (*pb.JsonDynamicConfig, error) {
	jsonDynamicConfig, err := c.jsonDynamicConfigRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return assembler.JsonDynamicConfigEntityToModel(jsonDynamicConfig), nil
}

// DeleteJsonDynamicConfig 删除JSON动态配置
func (c JsonDynamicConfigUseCase) DeleteJsonDynamicConfig(ctx context.Context, id uint64) error {
	return c.jsonDynamicConfigRepo.DeleteByID(ctx, id)
}
