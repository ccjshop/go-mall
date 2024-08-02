package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// HomeAdvertiseEntityToModel entity转pb
func HomeAdvertiseEntityToModel(homeAdvertise *entity.HomeAdvertise) *pb.HomeAdvertise {
	return &pb.HomeAdvertise{
		Id:   homeAdvertise.ID,
		Name: homeAdvertise.Name,
		Pic:  util.ImgUtils.GetFullUrl(homeAdvertise.Pic),
		Url:  homeAdvertise.URL,
		Sort: homeAdvertise.Sort,
		Note: homeAdvertise.Note,
		// 类型
		Type: uint32(homeAdvertise.Type),
		// 时间
		StartTime: homeAdvertise.StartTime,
		EndTime:   homeAdvertise.EndTime,
		// 状态
		Status: uint32(homeAdvertise.Status),
		// 统计
		ClickCount: homeAdvertise.ClickCount,
		OrderCount: homeAdvertise.OrderCount,
	}
}

// AddOrUpdateHomeAdvertiseParamToEntity pb转entity
func AddOrUpdateHomeAdvertiseParamToEntity(param *pb.AddOrUpdateHomeAdvertiseParam) *entity.HomeAdvertise {
	return &entity.HomeAdvertise{
		Name: param.Name,
		Pic:  util.ImgUtils.GetRelativeUrl(param.Pic),
		URL:  param.Url,
		Sort: param.Sort,
		Note: param.Note,
		// 类型
		Type: uint8(param.Type),
		// 时间
		StartTime: param.StartTime,
		EndTime:   param.EndTime,
		// 状态
		Status: uint8(param.Status),
		// 统计
		ClickCount: param.ClickCount,
		OrderCount: param.OrderCount,
	}
}
