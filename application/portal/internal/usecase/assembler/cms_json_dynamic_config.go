package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// JsonDynamicConfigEntityToModel entity转pb
func JsonDynamicConfigEntityToModel(jsonDynamicConfig *entity.JsonDynamicConfig) *pb.JsonDynamicConfig {
	return &pb.JsonDynamicConfig{}
}

func IntegrationConsumeSettingEntityToDetail(integrationConsumeSetting entity.UmsIntegrationConsumeSetting) *pb.GenerateConfirmOrderRsp_IntegrationConsumeSetting {
	return &pb.GenerateConfirmOrderRsp_IntegrationConsumeSetting{
		DeductionPerAmount: integrationConsumeSetting.DeductionPerAmount,
		MaxPercentPerOrder: integrationConsumeSetting.MaxPercentPerOrder,
		UseUnit:            integrationConsumeSetting.UseUnit,
		CouponStatus:       uint32(integrationConsumeSetting.CouponStatus),
	}
}
