package grpcsrv

import (
	"context"

	"github.com/ccjshop/go-mall/common/retcode"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// HomeContent 首页内容信息展示
func (s PortalApiImpl) HomeContent(ctx context.Context, req *pb.HomeContentReq) (*pb.HomeContentRsp, error) {
	return s.home.HomeContent(ctx, req)
}

// ProductCategoryList 获取首页商品分类
func (s PortalApiImpl) ProductCategoryList(ctx context.Context, req *pb.ProductCategoryListReq) (*pb.ProductCategoryListRsp, error) {
	var (
		res = &pb.ProductCategoryListRsp{}
	)
	categories, err := s.home.ProductCategoryList(ctx, req)
	if err != nil {
		return nil, err
	}
	res.Data = categories

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}
