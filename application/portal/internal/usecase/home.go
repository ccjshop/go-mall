package usecase

import (
	"context"

	"github.com/ccjshop/go-mall/application/portal/internal/usecase/assembler"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// HomeUseCase 首页
type HomeUseCase struct {
	productCategoryRepo IProductCategoryRepo // 操作商品分类
	homeAdvertiseRepo   IHomeAdvertiseRepo   // 操作首页轮播广告
	brandRepo           IBrandRepo           // 操作商品品牌
}

// NewHome 创建首页Service实现类
func NewHome(
	productCategoryRepo IProductCategoryRepo,
	homeAdvertiseRepo IHomeAdvertiseRepo,
	brandRepo IBrandRepo,
) *HomeUseCase {
	return &HomeUseCase{
		productCategoryRepo: productCategoryRepo,
		homeAdvertiseRepo:   homeAdvertiseRepo,
		brandRepo:           brandRepo,
	}
}

// HomeContent 获取首页内容
func (h HomeUseCase) HomeContent(ctx context.Context, req *pb.HomeContentReq) (*pb.HomeContentRsp, error) {
	res := &pb.HomeContentRsp{}
	// 获取首页广告
	advertises, err := h.homeAdvertiseRepo.GetHomeAdvertises(ctx)
	if err != nil {
		return nil, err
	}
	// 获取推荐品牌
	//h.brandRepo.GetRecommendBrandList(ctx, 0, 6)

	// 获取秒杀信息
	// 获取新品推荐
	// 获取人气推荐
	// 获取推荐专题

	res.Advertises = assembler.HomeAdvertisesEntityToDetail(advertises)
	return res, nil
}

// ProductCategoryList 分页查询订单
func (h HomeUseCase) ProductCategoryList(ctx context.Context, req *pb.ProductCategoryListReq) ([]*pb.ProductCategoryItem, error) {
	var (
		res = make([]*pb.ProductCategoryItem, 0)
	)
	categories, err := h.productCategoryRepo.GetShowProductCategory(ctx, req.GetParentId())
	if err != nil {
		return nil, err
	}
	for _, category := range categories {
		res = append(res, assembler.ProductCategoryEntityToModel(category))
	}
	return res, nil
}
