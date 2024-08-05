// Package app configures and runs application.
package app

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ccjshop/go-mall/application/portal/config"
	"github.com/ccjshop/go-mall/application/portal/internal/controller/grpcsrv"
	"github.com/ccjshop/go-mall/application/portal/internal/usecase"
	"github.com/ccjshop/go-mall/application/portal/internal/usecase/repo"
	"github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/interceptor"
	"github.com/ccjshop/go-mall/common/logger"
	"github.com/ccjshop/go-mall/common/pkg/crypto"
	"github.com/ccjshop/go-mall/common/pkg/jwt"
	"github.com/ccjshop/go-mall/common/util"
	pb "github.com/ccjshop/go-mall/proto/mall"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	customLog := logger.New(cfg.Log.Level)

	// 初始化数据库
	conn, err := db.GetConn(cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Timeout, cfg.DB.DbName)
	if err != nil {
		customLog.Fatal(fmt.Errorf("app - Run - db.GetConn: %w", err))
	}

	// gorm事务封装
	db.InitTransaction(conn)

	// oss url 前缀
	util.ImgUtils.InitBaseUrl(cfg.Oss.BaseUrl)

	// 全字段更新，初始化那些字段不更新，那些字段需要更新
	if err := repo.InitField(conn); err != nil {
		customLog.Fatal(fmt.Errorf("app - Run - repo.InitField: %w", err))
	}

	var (
		passwordEncoder = &crypto.BcryptPasswordEncoder{}
		jwtTokenUtil    = jwt.NewJWT(jwt.JwtConfig{
			TimeOut: time.Second * time.Duration(cfg.Jwt.TimeOut),
			Issuer:  cfg.Jwt.Issuer,
			SignKey: cfg.Jwt.SignKey,
		})
	)

	var (
		productCategoryRepo               = repo.NewProductCategoryRepo(conn)
		productRepo                       = repo.NewProductRepo(conn)
		brandRepo                         = repo.NewBrandRepo(conn)
		productAttributeRepo              = repo.NewProductAttributeRepo(conn)
		productAttributeValueRepo         = repo.NewProductAttributeValueRepo(conn)
		skuStockRepo                      = repo.NewSkuStockRepo(conn)
		memberRepo                        = repo.NewMemberRepo(conn)
		cartItemRepo                      = repo.NewCartItemRepo(conn)
		homeAdvertiseRepo                 = repo.NewHomeAdvertiseRepo(conn)
		orderRepo                         = repo.NewOrderRepo(conn)
		orderItemRepo                     = repo.NewOrderItemRepo(conn)
		memberReceiveAddressRepo          = repo.NewMemberReceiveAddressRepo(conn)
		jsonDynamicConfigRepo             = repo.NewJsonDynamicConfigRepo(conn)
		productLadderRepo                 = repo.NewProductLadderRepo(conn)
		productFullReductionRepo          = repo.NewProductFullReductionRepo(conn)
		couponRepo                        = repo.NewCouponRepo(conn)
		couponHistoryRepo                 = repo.NewCouponHistoryRepo(conn)
		couponProductCategoryRelationRepo = repo.NewCouponProductCategoryRelationRepo(conn)
		couponProductRelationRepo         = repo.NewCouponProductRelationRepo(conn)
	)
	homeUseCase := usecase.NewHome(productCategoryRepo, homeAdvertiseRepo, brandRepo)
	productUseCase := usecase.NewProduct(
		productRepo,
		brandRepo,
		productAttributeRepo,
		productAttributeValueRepo,
		skuStockRepo,
	)

	memberUseCase := usecase.NewMember(cfg, passwordEncoder, jwtTokenUtil, memberRepo)

	promotionUseCase := usecase.NewPromotion(productRepo, skuStockRepo, productLadderRepo, productFullReductionRepo)

	cartItemUseCase := usecase.NewCartItem(cartItemRepo, memberRepo, productRepo, brandRepo, promotionUseCase)

	couponUseCase := usecase.NewCoupon(couponRepo, couponHistoryRepo, couponProductCategoryRelationRepo, couponProductRelationRepo)

	orderUseCase := usecase.NewOrder(
		orderRepo,
		orderItemRepo,
		cartItemRepo,
		memberRepo,
		memberReceiveAddressRepo,
		jsonDynamicConfigRepo,
		skuStockRepo,
		couponHistoryRepo,
		cartItemUseCase,
		couponUseCase,
	)

	// grpc服务
	grpcSrvImpl := grpcsrv.New(
		homeUseCase,
		productUseCase,
		memberUseCase,
		cartItemUseCase,
		orderUseCase,
	)
	grpcServer, err := configGrpc(customLog, grpcSrvImpl, cfg, jwtTokenUtil, cfg.HTTP.IP, cfg.HTTP.Port)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - configGrpc: %w", err))
	}

	// 打印当前进程的 ID
	customLog.Info("project started with pid %d", os.Getpid())

	// 监听关闭信号
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		customLog.Info("app - Run - signal: " + s.String())
	}

	// grpc优雅关闭，5s内必须完成关闭
	gracefulStopWithTimeout(customLog, grpcServer, 5*time.Second)
}

// gracefulStopWithTimeout 尝试在给定的超时时间内优雅地关闭 gRPC 服务器。
// 如果服务器在超时时间内成功关闭，那么函数会打印一条信息并返回。
// 如果超时时间到了但服务器还没有关闭，那么函数会强制关闭服务器并打印一条错误信息。
func gracefulStopWithTimeout(customLog *logger.Logger, grpcServer *grpc.Server, timeout time.Duration) {
	// 创建一个带有超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// 创建一个通道，用于接收服务器关闭的信号
	ch := make(chan struct{})
	go func() {
		// 在一个新的 goroutine 中优雅地关闭服务器
		grpcServer.GracefulStop()
		// 当服务器关闭时，关闭通道
		close(ch)
	}()

	select {
	case <-ch:
		// 如果从通道接收到了信号，说明服务器已经成功关闭
		customLog.Info("gRPC server stopped gracefully")
	case <-ctx.Done():
		// 如果上下文超时，说明服务器没有在给定的时间内关闭，此时强制关闭服务器
		customLog.Error("gRPC server stop timeout, force stopping")
		grpcServer.Stop()
	}
}

func configGrpc(customLog *logger.Logger, grpcSrvImpl grpcsrv.PortalApi, cfg *config.Config, jwtTokenUtil *jwt.JWT, ip string, port uint32) (*grpc.Server, error) {
	var (
		addr = fmt.Sprintf("%s:%d", ip, port)
	)
	// 创建一个gRPC server对象
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.ChainUnaryServerInterceptors(
			//interceptor.JWTAuthInterceptor(jwtTokenUtil, cfg.Jwt.Whitelist, cfg.Jwt.TokenHeader, cfg.Jwt.TokenHead),
			interceptor.ValidationInterceptor,
			interceptor.PanicRecoveryInterceptor,
		),
		),
	)

	// 注册grpc服务
	pb.RegisterPortalHomeApiServer(grpcServer, grpcSrvImpl)
	pb.RegisterPortalProductApiServer(grpcServer, grpcSrvImpl)
	pb.RegisterPortalMemberApiServer(grpcServer, grpcSrvImpl)
	pb.RegisterPortalCartItemApiServer(grpcServer, grpcSrvImpl)
	pb.RegisterPortalOrderApiServer(grpcServer, grpcSrvImpl)
	pb.RegisterPortalCouponApiServer(grpcServer, grpcSrvImpl)

	// gRPC-Gateway mux
	gwmux := runtime.NewServeMux(
		runtime.WithMetadata(interceptor.CustomAnnotator), // 注册自定义 Annotator
		runtime.WithErrorHandler(interceptor.CustomHTTPError),
	)
	dops := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterPortalHomeApiHandlerFromEndpoint(context.Background(), gwmux, addr, dops); err != nil {
		return nil, err
	}
	if err := pb.RegisterPortalProductApiHandlerFromEndpoint(context.Background(), gwmux, addr, dops); err != nil {
		return nil, err
	}
	if err := pb.RegisterPortalMemberApiHandlerFromEndpoint(context.Background(), gwmux, addr, dops); err != nil {
		return nil, err
	}
	if err := pb.RegisterPortalCartItemApiHandlerFromEndpoint(context.Background(), gwmux, addr, dops); err != nil {
		return nil, err
	}
	if err := pb.RegisterPortalOrderApiHandlerFromEndpoint(context.Background(), gwmux, addr, dops); err != nil {
		return nil, err
	}
	if err := pb.RegisterPortalCouponApiHandlerFromEndpoint(context.Background(), gwmux, addr, dops); err != nil {
		return nil, err
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	// 创建一个自定义的 CORS 中间件配置
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"*"}, // 允许的源列表
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           86400, // 预检请求的结果可以被缓存的最大秒数
	}

	// 统一返回值处理
	responseWrapper := interceptor.WrapResponseMiddleware(mux)
	// jwt拦截
	jwtWrapper := interceptor.NewJWTAuthMiddleware(responseWrapper, jwtTokenUtil, cfg.Jwt.Blacklist, cfg.Jwt.Whitelist, cfg.Jwt.TokenHeader, cfg.Jwt.TokenHead)
	// 请求响应日志记录
	logWrapper := interceptor.NewLoggingMiddleware(jwtWrapper)
	// cors处理器
	corsWrapper := cors.New(corsOptions).Handler(logWrapper)

	// 定义HTTP server配置
	gwServer := &http.Server{
		Addr:    addr,
		Handler: grpcHandlerFunc(grpcServer, corsWrapper), // 请求的统一入口
	}

	// tpc监听
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	go func() {
		customLog.Fatal(gwServer.Serve(lis)) // 启动HTTP服务
	}()

	return grpcServer, nil
}

// grpcHandlerFunc 将gRPC请求和HTTP请求分别调用不同的handler处理
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}
