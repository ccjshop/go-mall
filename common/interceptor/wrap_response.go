package interceptor

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/ccjshop/go-mall/common/retcode"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// standardResponse 统一的响应结构
type standardResponse struct {
	Code    uint32      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

var (
	marshalError = `{"code": 500, "message": "failed to marshal error message"}`
)

// CustomHTTPError 自定义错误处理函数
func CustomHTTPError(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, req *http.Request, err error) {
	var (
		bizCode  = retcode.RetInternalError
		httpCode = uint32(http.StatusInternalServerError)
		message  = ""
	)
	// 检查错误类型
	st, ok := status.FromError(err)
	if ok {
		bizCode = st.Code()
		message = st.Message()
	} else {
		bizCode = codes.Internal
		message = err.Error()
	}

	// 查找自定义错误映射，如果没有找到，则使用默认的映射
	httpCode, message = retcode.GetHttpCodeAndMsg(bizCode)

	// 设置响应的Content-Type和状态码
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(httpCode))

	// 创建包含code和message的JSON响应体
	responseBody := standardResponse{
		Code:    uint32(bizCode),
		Message: message,
	}

	// 发送JSON响应
	if err := marshaler.NewEncoder(w).Encode(responseBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(marshalError))
	}
}

// WrapResponseMiddleware 中间件函数，用于包装响应
func WrapResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 拦截响应
		wrappedWriter := &responseWriterInterceptor{ResponseWriter: w, body: &bytes.Buffer{}}
		next.ServeHTTP(wrappedWriter, r)

		// 如果响应已经被写入，则不进行包装
		if wrappedWriter.written {
			return
		}

		// 检查响应是否已经是标准格式，返回的内容如果是非结构体（结构体数组、map），需要自己定义code+message+data
		var standardResp standardResponse
		if err := json.Unmarshal(wrappedWriter.body.Bytes(), &standardResp); err == nil && standardResp.Code != 0 {
			// 响应已经是标准格式，直接写入原始的响应Writer
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write(wrappedWriter.body.Bytes())
			return
		}

		// 解析原始响应体
		var data interface{}
		if wrappedWriter.body.Len() > 0 {
			if err := json.Unmarshal(wrappedWriter.body.Bytes(), &data); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		}

		// 创建统一的响应结构
		response := standardResponse{
			Code:    http.StatusOK, // 或者其他适当的成功状态码
			Message: "OK",
			Data:    data,
		}

		// 将统一的响应结构写入原始的响应Writer
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(marshalError))
		}
	})
}

// responseWriterInterceptor 是一个拦截器，用于捕获响应体
type responseWriterInterceptor struct {
	http.ResponseWriter
	body    *bytes.Buffer
	written bool
}

func (w *responseWriterInterceptor) WriteHeader(statusCode int) {
	if w.written {
		return
	}
	w.ResponseWriter.WriteHeader(statusCode)
	w.written = true
}

func (w *responseWriterInterceptor) Write(b []byte) (int, error) {
	if !w.written {
		// 在写入响应之前，先将响应体捕获到缓冲区中
		return w.body.Write(b)
	}
	return w.ResponseWriter.Write(b)
}
