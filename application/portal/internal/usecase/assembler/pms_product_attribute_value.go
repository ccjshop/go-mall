package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

func ProductAttributeValuesToDetail(attributeValues entity.ProductAttributeValues) []*pb.ProductDetailRsp_ProductAttributeValue {
	res := make([]*pb.ProductDetailRsp_ProductAttributeValue, 0)
	for _, attributeValue := range attributeValues {
		res = append(res, &pb.ProductDetailRsp_ProductAttributeValue{
			Id:                 attributeValue.ID,
			ProductId:          attributeValue.ProductID,
			ProductAttributeId: attributeValue.ProductAttributeID,
			Value:              attributeValue.Value,
		})
	}
	return res

}
