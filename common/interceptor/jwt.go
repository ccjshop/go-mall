package interceptor

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/ccjshop/go-mall/common/pkg/jwt"
	pkg_jwt "github.com/ccjshop/go-mall/common/pkg/jwt"
	"github.com/ccjshop/go-mall/common/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// NewJWTAuthMiddleware 是一个用于验证 JWT 令牌的 HTTP 中间件
func NewJWTAuthMiddleware(next http.Handler, jwtTokenUtil *pkg_jwt.JWT, blacklist []string, whitelist []string, tokenHeader string, tokenHead string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 尝试解析 JWT
		_, ctx, parseErr := parseJWT(r, jwtTokenUtil, tokenHeader, tokenHead)
		if len(blacklist) != 0 {
			blackLogic(ctx, next, w, r, blacklist, parseErr)
		} else {
			whiteLogic(ctx, next, w, r, whitelist, parseErr)
		}
	})
}

func whiteLogic(ctx context.Context, next http.Handler, w http.ResponseWriter, r *http.Request, whitelist []string, parseErr error) {
	// 检查白名单，白名单直接放行
	for _, path := range whitelist {
		if strings.HasPrefix(r.URL.Path, path) {
			// 如果在白名单中，直接传递给下一个处理器
			next.ServeHTTP(w, r)
			return
		}
	}

	// 如果请求不在白名单中，且JWT解析失败
	if parseErr != nil {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}

	// 如果JWT解析成功，且请求不在黑名单中，也不在白名单中，传递给下一个处理器
	next.ServeHTTP(w, r.WithContext(ctx))
}

func blackLogic(ctx context.Context, next http.Handler, w http.ResponseWriter, r *http.Request, blacklist []string, parseErr error) {
	// 检查黑名单，黑名单url必须限制登录
	for _, path := range blacklist {
		if strings.HasPrefix(r.URL.Path, path) {
			if parseErr != nil {
				// 如果在黑名单中且JWT解析失败，返回错误
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}
			// 如果在黑名单中且JWT解析成功，将上下文传递到下一个处理器
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
	}
	// 未命中黑名单，直接放行
	next.ServeHTTP(w, r.WithContext(ctx))
}

func parseJWT(r *http.Request, jwtTokenUtil *pkg_jwt.JWT, tokenHeader string, tokenHead string) (*jwt.CustomClaims, context.Context, error) {
	authHeader := r.Header.Get(tokenHeader)
	if authHeader == "" {
		return nil, r.Context(), errors.New("authorization header is required")
	}

	bearerToken := strings.TrimPrefix(authHeader, tokenHead)
	token, err := jwtTokenUtil.ParseToken(bearerToken)
	if err != nil {
		return nil, r.Context(), err
	}

	// 用户信息放入上下文
	ctx := util.CtxUtils.SetUserID(r.Context(), token.UserInfo.UserID)
	return token, ctx, nil
}

// CustomAnnotator 将 HTTP 上下文中的用户名添加到 gRPC 上下文的元数据中
func CustomAnnotator(ctx context.Context, req *http.Request) metadata.MD {
	// 从HTTP context中获取用户名
	userID, err := util.CtxUtils.GetUserID(ctx)
	if err != nil {
		return nil
	}

	// 创建gRPC metadata并添加用户名
	md := metadata.Pairs(string(util.CtxUtils.GetUserIDKey()), fmt.Sprintf("%d", userID))
	return md
}

// JWTAuthInterceptor 创建一个新的 Unary 拦截器，用于验证 JWT 令牌。
func JWTAuthInterceptor(jwtTokenUtil *pkg_jwt.JWT, whitelist []string, tokenHeader string, tokenHead string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// 检查当前请求的方法是否在白名单中
		//for _, path := range whitelist {
		//	if strings.HasPrefix(info.FullMethod, path) {
		//		// 如果在白名单中，直接调用 RPC 方法
		//		return handler(ctx, req)
		//	}
		//}

		// 从 gRPC 元数据中获取 Authorization 头
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "Authorization metadata is required")
		}

		authHeader, ok := md[strings.ToLower(tokenHeader)]
		if !ok || len(authHeader) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "Authorization header is required")
		}

		bearerToken := strings.TrimPrefix(authHeader[0], tokenHead)

		token, err := jwtTokenUtil.ParseToken(bearerToken)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "Invalid token: %v", err)
		}

		// 将用户名设置到 gRPC 上下文中
		newCtx := util.CtxUtils.SetUserID(ctx, token.UserInfo.UserID)

		// 调用 RPC 方法
		return handler(newCtx, req)
	}
}
