package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

func ProductCategoryEntityToModel(category *entity.ProductCategory) *pb.ProductCategory {
	return &pb.ProductCategory{
		Id:          category.ID,
		ParentId:    category.ParentID,
		Name:        category.Name,
		Icon:        util.ImgUtils.GetFullUrl(category.Icon),
		ProductUnit: category.ProductUnit,
		Sort:        category.Sort,
		//
		Description: category.Description,
		Keywords:    category.Keywords,
		// 状态
		NavStatus:  uint32(category.NavStatus),
		ShowStatus: uint32(category.ShowStatus),
		// 计算得出
		Level: uint32(category.Level),
		// 冗余字段
		ProductCount: category.ProductCount,
	}
}

func AddOrUpdateProductCategoryParamToEntity(param *pb.AddOrUpdateProductCategoryParam) *entity.ProductCategory {
	return &entity.ProductCategory{
		ParentID:    param.ParentId,
		Name:        param.Name,
		Icon:        util.ImgUtils.GetRelativeUrl(param.Icon),
		ProductUnit: param.ProductUnit,
		Sort:        param.Sort,
		// 状态
		NavStatus:  uint8(param.NavStatus),
		ShowStatus: uint8(param.ShowStatus),
		//
		Keywords:    param.Keywords,
		Description: param.Description,
	}
}
