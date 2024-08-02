package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// ProductAttributeEntityToModel entity转pb
func ProductAttributeEntityToModel(productAttribute *entity.ProductAttribute) *pb.ProductAttribute {
	return &pb.ProductAttribute{
		Id:                         productAttribute.ID,
		Type:                       uint32(productAttribute.Type),
		ProductAttributeCategoryId: productAttribute.ProductAttributeCategoryID,
		Name:                       productAttribute.Name,
		Sort:                       productAttribute.Sort,
		//
		SelectType: uint32(productAttribute.SelectType),
		InputType:  uint32(productAttribute.InputType),
		InputList:  productAttribute.InputList,
		//
		FilterType: uint32(productAttribute.FilterType),
		SearchType: uint32(productAttribute.SearchType),
		//
		RelatedStatus: uint32(productAttribute.RelatedStatus),
		HandAddStatus: uint32(productAttribute.HandAddStatus),
	}

}

// AddOrUpdateProductAttributeParamToEntity pb转entity
func AddOrUpdateProductAttributeParamToEntity(param *pb.AddOrUpdateProductAttributeParam) *entity.ProductAttribute {
	return &entity.ProductAttribute{
		Type:                       uint8(param.GetType()),
		ProductAttributeCategoryID: param.GetProductAttributeCategoryId(),
		Name:                       param.GetName(),
		SelectType:                 uint8(param.GetSelectType()),
		InputType:                  uint8(param.GetInputType()),
		InputList:                  param.GetInputList(),
		Sort:                       param.GetSort(),
		FilterType:                 uint8(param.GetFilterType()),
		SearchType:                 uint8(param.GetSearchType()),
		RelatedStatus:              uint8(param.GetRelatedStatus()),
		HandAddStatus:              uint8(param.GetHandAddStatus()),
	}
}
