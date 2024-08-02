package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

func ProductCategoryEntityToModel(category *entity.ProductCategory) *pb.ProductCategoryItem {
	return &pb.ProductCategoryItem{
		Id:       category.ID,
		ParentId: category.ParentID,
		Name:     category.Name,
		Icon:     util.ImgUtils.GetFullUrl(category.Icon),
	}
}
