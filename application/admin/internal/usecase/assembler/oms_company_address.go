package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// CompanyAddressEntityToModel entity转pb
func CompanyAddressEntityToModel(companyAddress *entity.CompanyAddress) *pb.CompanyAddress {
	return &pb.CompanyAddress{
		Id:            companyAddress.ID,
		AddressName:   companyAddress.AddressName,
		SendStatus:    uint32(companyAddress.SendStatus),
		ReceiveStatus: uint32(companyAddress.ReceiveStatus),
		Name:          companyAddress.Name,
		Phone:         companyAddress.Phone,
		Province:      companyAddress.Province,
		City:          companyAddress.City,
		Region:        companyAddress.Region,
		DetailAddress: companyAddress.DetailAddress,
	}
}

// AddOrUpdateCompanyAddressParamToEntity pb转entity
func AddOrUpdateCompanyAddressParamToEntity(param *pb.AddOrUpdateCompanyAddressParam) *entity.CompanyAddress {
	return &entity.CompanyAddress{}
}
