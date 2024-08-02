package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// HomeAdvertisesEntityToDetail entity转pb
func HomeAdvertisesEntityToDetail(homeAdvertises []*entity.HomeAdvertise) []*pb.HomeContentRsp_HomeAdvertise {
	res := make([]*pb.HomeContentRsp_HomeAdvertise, 0)
	for _, homeAdvertise := range homeAdvertises {
		res = append(res, HomeAdvertiseEntityToDetail(homeAdvertise))
	}
	return res
}

// HomeAdvertiseEntityToDetail entity转pb
func HomeAdvertiseEntityToDetail(homeAdvertise *entity.HomeAdvertise) *pb.HomeContentRsp_HomeAdvertise {
	return &pb.HomeContentRsp_HomeAdvertise{
		Pic: util.ImgUtils.GetFullUrl(homeAdvertise.Pic),
		Url: homeAdvertise.URL,
	}
}
