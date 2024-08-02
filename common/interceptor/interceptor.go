package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

// ChainUnaryServerInterceptors 设置多个UnaryInterceptor
func ChainUnaryServerInterceptors(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		for i := len(interceptors) - 1; i >= 0; i-- {
			handler = createHandler(interceptors[i], handler)
		}
		return handler(ctx, req)
	}
}

func createHandler(interceptor grpc.UnaryServerInterceptor, handler grpc.UnaryHandler) grpc.UnaryHandler {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return interceptor(ctx, req, nil, handler)
	}
}
