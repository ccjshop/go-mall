package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// JsonDynamicConfigEntityToModel entity转pb
func JsonDynamicConfigEntityToModel(jsonDynamicConfig *entity.JsonDynamicConfig) *pb.JsonDynamicConfig {
	return &pb.JsonDynamicConfig{
		Id:         jsonDynamicConfig.ID,
		BizType:    string(jsonDynamicConfig.BizType),
		BizDesc:    jsonDynamicConfig.BizDesc,
		Content:    jsonDynamicConfig.Content,
		JsonSchema: jsonDynamicConfig.JsonSchema,
	}
}

// AddOrUpdateJsonDynamicConfigParamToEntity pb转entity
func AddOrUpdateJsonDynamicConfigParamToEntity(param *pb.AddOrUpdateJsonDynamicConfigParam) *entity.JsonDynamicConfig {
	return &entity.JsonDynamicConfig{
		BizType:    entity.BizType(param.BizType),
		BizDesc:    param.BizDesc,
		Content:    param.Content,
		JsonSchema: param.JsonSchema,
	}
}
