package interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// CorsInterceptor 跨域
func CorsInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			// 检查请求元数据中的"Origin"，并响应相应的CORS头部
			if origin := md.Get("Origin"); len(origin) > 0 {
				// 设置CORS头部
				metadata.AppendToOutgoingContext(ctx, "Access-Control-Allow-Origin", "*")
				metadata.AppendToOutgoingContext(ctx, "Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
				metadata.AppendToOutgoingContext(ctx, "Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			}
		}

		// 继续处理请求
		return handler(ctx, req)
	}
}
