package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

func ProductEntityToProductListItem(product *entity.Product) *pb.SearchProductRsp_Product {
	return &pb.SearchProductRsp_Product{
		Id:       product.ID,
		Pic:      util.ImgUtils.GetFullUrl(product.Pic),
		Name:     product.Name,
		SubTitle: product.SubTitle,
		Price:    product.Price.String(),
		Sale:     product.Sale,
	}
}

func ProductEntityToDetail(product *entity.Product) *pb.ProductDetailRsp_Product {
	return &pb.ProductDetailRsp_Product{
		Id:            product.ID,
		AlbumPics:     util.ImgUtils.GetFullUrls(product.AlbumPics),
		Pic:           util.ImgUtils.GetFullUrl(product.Pic),
		Name:          product.Name,
		SubTitle:      product.SubTitle,
		Price:         product.Price.String(),
		OriginalPrice: product.OriginalPrice.String(),
		Sale:          product.Sale,
		Stock:         product.Stock,
		ServiceIds:    product.ServiceIDs,
	}
}
