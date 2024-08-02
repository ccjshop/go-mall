package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

func BrandEntityToDetail(brand *entity.Brand) *pb.ProductDetailRsp_Brand {
	return &pb.ProductDetailRsp_Brand{
		Name:        brand.Name,
		FirstLetter: brand.FirstLetter,
		Logo:        util.ImgUtils.GetFullUrl(brand.Logo),
	}
}
