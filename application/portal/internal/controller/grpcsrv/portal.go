package grpcsrv

import (
	"github.com/ccjshop/go-mall/application/portal/internal/usecase"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

var _ pb.PortalHomeApiServer = &PortalApiImpl{}
var _ pb.PortalProductApiServer = &PortalApiImpl{}
var _ pb.PortalMemberApiServer = &PortalApiImpl{}
var _ pb.PortalCartItemApiServer = &PortalApiImpl{}
var _ pb.PortalOrderApiServer = &PortalApiImpl{}
var _ pb.PortalCouponApiServer = &PortalApiImpl{}

type PortalApi interface {
	pb.PortalHomeApiServer
	pb.PortalProductApiServer
	pb.PortalMemberApiServer
	pb.PortalCartItemApiServer
	pb.PortalOrderApiServer
	pb.PortalCouponApiServer
}

type PortalApiImpl struct {
	pb.UnimplementedPortalHomeApiServer
	pb.UnimplementedPortalProductApiServer
	pb.UnimplementedPortalMemberApiServer
	pb.UnimplementedPortalCartItemApiServer
	pb.UnimplementedPortalOrderApiServer
	pb.UnimplementedPortalCouponApiServer

	home            usecase.IHomeUseCase
	product         usecase.IProductUseCase
	memberUseCase   usecase.IMemberUseCase
	cartItemUseCase usecase.ICartItemUseCase
	orderUseCase    usecase.IOrderUseCase
}

func New(
	home usecase.IHomeUseCase,
	product usecase.IProductUseCase,
	memberUseCase usecase.IMemberUseCase,
	cartItemUseCase usecase.ICartItemUseCase,
	orderUseCase usecase.IOrderUseCase,
) PortalApi {
	return &PortalApiImpl{
		home:            home,
		product:         product,
		memberUseCase:   memberUseCase,
		cartItemUseCase: cartItemUseCase,
		orderUseCase:    orderUseCase,
	}
}
