package ctx

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/ccjshop/go-mall/common/retcode"
	"google.golang.org/grpc/metadata"
)

// CtxUtils 上下文工具类
type CtxUtils struct {
}

// 定义一个类型，用作上下文中的键
type contextKey string

// 定义上下文中使用的键
const (
	UsernameKey     contextKey = "username"     // 用户名
	UserIDKey       contextKey = "userId"       // 用户名
	RequestIDHeader contextKey = "X-Request-ID" // 请求ID
)

// SetUserID 将用户id设置到上下文中
func (c CtxUtils) SetUserID(ctx context.Context, userID uint64) context.Context {
	return context.WithValue(ctx, UserIDKey, fmt.Sprintf("%d", userID))
}

// GetUserID 从上下文中获取用户ID
func (c CtxUtils) GetUserID(ctx context.Context) (uint64, error) {
	if userID, ok := c.getFromHttpContext(ctx, UserIDKey); ok {
		return userID, nil
	}
	if userID, ok := c.getFromGrpcContext(ctx, UserIDKey); ok {
		return userID, nil
	}
	return 0, retcode.NewError(retcode.NeedLogin)
}

func (c CtxUtils) getFromHttpContext(ctx context.Context, key contextKey) (uint64, bool) {
	value, ok := ctx.Value(key).(string)
	if !ok {
		return 0, false
	}
	userID := strToUint64(value, 0)
	return userID, true
}

func (c CtxUtils) getFromGrpcContext(ctx context.Context, key contextKey) (uint64, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return 0, false
	}
	slice, ok := md[strings.ToLower(string(key))]
	if !ok || len(slice) == 0 {
		return 0, false
	}
	userID := strToUint64(slice[0], 0)
	return userID, true
}

func (c CtxUtils) GetUserIDKey() string {
	return string(UserIDKey)
}

func strToUint64(strNum string, defaultNum ...uint64) uint64 {
	num, err := strconv.ParseUint(strNum, 10, 64)
	if err != nil {
		if len(defaultNum) > 0 {
			return defaultNum[0]
		} else {
			return 0
		}
	}
	return num
}
