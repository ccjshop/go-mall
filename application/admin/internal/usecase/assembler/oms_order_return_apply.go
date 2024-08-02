package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// OrderReturnApplyEntityToModel entity转pb
func OrderReturnApplyEntityToModel(orderReturnApply *entity.OrderReturnApply) *pb.OrderReturnApply {
	return &pb.OrderReturnApply{
		Id:      orderReturnApply.ID,
		OrderId: orderReturnApply.OrderID,
		// 商品信息
		ProductPic:       util.ImgUtils.GetFullUrl(orderReturnApply.ProductPic),
		ProductName:      orderReturnApply.ProductName,
		ProductBrand:     orderReturnApply.ProductBrand,
		ProductId:        orderReturnApply.ProductID,
		ProductRealPrice: orderReturnApply.ProductRealPrice.String(),
		ProductAttr:      orderReturnApply.ProductAttr,
		ProductCount:     orderReturnApply.ProductCount,
		ProductPrice:     orderReturnApply.ProductPrice.String(),
		//
		Status:         uint32(orderReturnApply.Status),
		OrderSn:        orderReturnApply.OrderSN,
		CreateTime:     orderReturnApply.CreateTime,
		MemberUsername: orderReturnApply.MemberUsername,
		ReturnName:     orderReturnApply.ReturnName,
		ReturnPhone:    orderReturnApply.ReturnPhone,
		Reason:         orderReturnApply.Reason,
		Description:    orderReturnApply.Description,
		ProofPics:      util.ImgUtils.GetFullUrls(orderReturnApply.ProofPics),
		//
		ReturnAmount:     orderReturnApply.ReturnAmount.String(),
		CompanyAddressId: orderReturnApply.CompanyAddressID,
		// 商家-处理人
		HandleMan:  orderReturnApply.HandleMan,
		HandleTime: orderReturnApply.HandleTime,
		HandleNote: orderReturnApply.HandleNote,
		// 商家-收货人
		ReceiveMan:  orderReturnApply.ReceiveMan,
		ReceiveTime: orderReturnApply.ReceiveTime,
		ReceiveNote: orderReturnApply.ReceiveNote,
	}
}
