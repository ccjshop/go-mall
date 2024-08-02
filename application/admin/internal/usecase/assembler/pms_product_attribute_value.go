package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// ProductAttributeValuesToModel entity转pb
func ProductAttributeValuesToModel(productAttributeValues []*entity.ProductAttributeValue) []*pb.ProductAttributeValue {
	res := make([]*pb.ProductAttributeValue, 0)
	for _, productAttributeValue := range productAttributeValues {
		res = append(res, ProductAttributeValueToModel(productAttributeValue))
	}
	return res
}

// ProductAttributeValueToModel entity转pb
func ProductAttributeValueToModel(productAttributeValue *entity.ProductAttributeValue) *pb.ProductAttributeValue {
	return &pb.ProductAttributeValue{
		Id:                 productAttributeValue.ID,
		ProductId:          productAttributeValue.ProductID,
		ProductAttributeId: productAttributeValue.ProductAttributeID,
		Value:              productAttributeValue.Value,
	}
}

// ProductAttributeValuesToEntity pb转entity
func ProductAttributeValuesToEntity(productAttributeValuePbs []*pb.ProductAttributeValue) []*entity.ProductAttributeValue {
	res := make([]*entity.ProductAttributeValue, 0)
	for _, productAttributeValuePb := range productAttributeValuePbs {
		res = append(res, ProductAttributeValueToEntity(productAttributeValuePb))
	}
	return res
}

// ProductAttributeValueToEntity pb转entity
func ProductAttributeValueToEntity(productAttributeValuePb *pb.ProductAttributeValue) *entity.ProductAttributeValue {
	return &entity.ProductAttributeValue{
		ID:                 productAttributeValuePb.Id,
		ProductID:          productAttributeValuePb.ProductId,
		ProductAttributeID: productAttributeValuePb.ProductAttributeId,
		Value:              productAttributeValuePb.Value,
	}
}
