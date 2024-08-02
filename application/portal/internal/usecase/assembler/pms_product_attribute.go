package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

func ProductAttributesToDetail(productAttributes entity.ProductAttributes) []*pb.ProductDetailRsp_ProductAttribute {
	res := make([]*pb.ProductDetailRsp_ProductAttribute, 0)
	for _, productAttribute := range productAttributes {
		res = append(res, &pb.ProductDetailRsp_ProductAttribute{
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
		})
	}
	return res
}
