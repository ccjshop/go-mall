package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// MemberReceiveAddressEntityToModel entity转pb
func MemberReceiveAddressEntityToModel(memberReceiveAddress *entity.MemberReceiveAddress) *pb.MemberReceiveAddress {
	return &pb.MemberReceiveAddress{}
}

// MemberReceiveAddressEntityToDetail entity转pb
func MemberReceiveAddressEntityToDetail(memberReceiveAddresses []*entity.MemberReceiveAddress) []*pb.GenerateConfirmOrderRsp_MemberReceiveAddress {
	res := make([]*pb.GenerateConfirmOrderRsp_MemberReceiveAddress, 0)
	for _, memberReceiveAddress := range memberReceiveAddresses {
		res = append(res, &pb.GenerateConfirmOrderRsp_MemberReceiveAddress{
			Id:            memberReceiveAddress.ID,
			MemberId:      memberReceiveAddress.MemberID,
			Name:          memberReceiveAddress.Name,
			PhoneNumber:   memberReceiveAddress.PhoneNumber,
			DefaultStatus: uint32(memberReceiveAddress.DefaultStatus),
			PostCode:      memberReceiveAddress.PostCode,
			Province:      memberReceiveAddress.Province,
			City:          memberReceiveAddress.City,
			Region:        memberReceiveAddress.Region,
			DetailAddress: memberReceiveAddress.DetailAddress,
		})
	}
	return res
}
