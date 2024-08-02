package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// PrefrenceAreaEntityToModel entity转pb
func PrefrenceAreaEntityToModel(prefrenceArea *entity.PrefrenceArea) *pb.PrefrenceArea {
	return &pb.PrefrenceArea{
		Id:         prefrenceArea.ID,
		Name:       prefrenceArea.Name,
		SubTitle:   prefrenceArea.SubTitle,
		Pic:        prefrenceArea.Pic,
		Sort:       prefrenceArea.Sort,
		ShowStatus: uint32(prefrenceArea.ShowStatus),
	}
}

// AddOrUpdatePrefrenceAreaParamToEntity pb转entity
func AddOrUpdatePrefrenceAreaParamToEntity(param *pb.AddOrUpdatePrefrenceAreaParam) *entity.PrefrenceArea {
	return &entity.PrefrenceArea{}
}
