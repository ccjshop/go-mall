package grpcsrv

import (
	"github.com/ccjshop/go-mall/application/admin/internal/usecase"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

var _ pb.CmsAdminApiServer = &AdminApiImpl{}
var _ pb.OmsAdminApiServer = &AdminApiImpl{}
var _ pb.SmsAdminApiServer = &AdminApiImpl{}

type AdminApi interface {
	pb.CmsAdminApiServer
	pb.OmsAdminApiServer
	pb.SmsAdminApiServer
}

type AdminApiImpl struct {
	pb.UnimplementedCmsAdminApiServer
	pb.UnimplementedOmsAdminApiServer
	pb.UnimplementedSmsAdminApiServer

	category                 usecase.IProductCategoryUseCase
	brand                    usecase.IBrandUseCase
	productAttributeCategory usecase.IProductAttributeCategoryUseCase
	productAttribute         usecase.IProductAttributeUseCase
	product                  usecase.IProductUseCase
	skuStock                 usecase.ISkuStockUseCase
	subject                  usecase.ISubjectUseCase
	prefrenceArea            usecase.IPrefrenceAreaUseCase
	jsonDynamicConfig        usecase.IJsonDynamicConfigUseCase
	//
	orderReturnReason usecase.IOrderReturnReasonUseCase
	order             usecase.IOrderUseCase
	orderReturnApply  usecase.IOrderReturnApplyUseCase
	companyAddress    usecase.ICompanyAddressUseCase
	//
	homeAdvertise usecase.IHomeAdvertiseUseCase
}

func New(category usecase.IProductCategoryUseCase,
	brand usecase.IBrandUseCase,
	productAttributeCategory usecase.IProductAttributeCategoryUseCase,
	productAttribute usecase.IProductAttributeUseCase,
	product usecase.IProductUseCase,
	skuStock usecase.ISkuStockUseCase,
	subject usecase.ISubjectUseCase,
	prefrenceArea usecase.IPrefrenceAreaUseCase,
	jsonDynamicConfig usecase.IJsonDynamicConfigUseCase,
	//
	orderReturnReason usecase.IOrderReturnReasonUseCase,
	order usecase.IOrderUseCase,
	orderReturnApply usecase.IOrderReturnApplyUseCase,
	companyAddress usecase.ICompanyAddressUseCase,
	//
	homeAdvertise usecase.IHomeAdvertiseUseCase,
) AdminApi {
	return &AdminApiImpl{
		category:                 category,
		brand:                    brand,
		productAttributeCategory: productAttributeCategory,
		productAttribute:         productAttribute,
		product:                  product,
		skuStock:                 skuStock,
		subject:                  subject,
		prefrenceArea:            prefrenceArea,
		jsonDynamicConfig:        jsonDynamicConfig,
		//
		orderReturnReason: orderReturnReason,
		order:             order,
		orderReturnApply:  orderReturnApply,
		companyAddress:    companyAddress,
		homeAdvertise:     homeAdvertise,
	}

}
