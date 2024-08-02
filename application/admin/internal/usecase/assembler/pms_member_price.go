package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// MemberPricesToModel entity转pb
func MemberPricesToModel(memberPrices []*entity.MemberPrice) []*pb.MemberPrice {
	res := make([]*pb.MemberPrice, 0)
	for _, memberPrice := range memberPrices {
		res = append(res, MemberPriceToModel(memberPrice))
	}
	return res
}

// MemberPriceToModel entity转pb
func MemberPriceToModel(memberPrice *entity.MemberPrice) *pb.MemberPrice {
	return &pb.MemberPrice{
		Id:              memberPrice.ID,
		ProductId:       memberPrice.ProductID,
		MemberLevelId:   memberPrice.MemberLevelID,
		MemberPrice:     memberPrice.MemberPrice.String(),
		MemberLevelName: memberPrice.MemberLevelName,
	}
}

// MemberPricesToEntity pb转entity
func MemberPricesToEntity(memberPricePbs []*pb.MemberPrice) []*entity.MemberPrice {
	res := make([]*entity.MemberPrice, 0)
	for _, memberPricePb := range memberPricePbs {
		res = append(res, MemberPriceToEntity(memberPricePb))
	}
	return res
}

// MemberPriceToEntity pb转entity
func MemberPriceToEntity(memberPricePb *pb.MemberPrice) *entity.MemberPrice {
	return &entity.MemberPrice{
		ID:              memberPricePb.Id,
		ProductID:       memberPricePb.ProductId,
		MemberLevelID:   memberPricePb.MemberLevelId,
		MemberPrice:     util.DecimalUtils.ToDecimalFixed2(memberPricePb.MemberPrice),
		MemberLevelName: memberPricePb.MemberLevelName,
	}
}
