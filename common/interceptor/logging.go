package interceptor

import (
	"log"
	"net/http"
)

// responseRecorder 是一个包装 http.ResponseWriter 的结构体，
// 它记录了状态码和写入的字节数。
type responseRecorder struct {
	http.ResponseWriter
	status  int
	written int64
}

// WriteHeader 捕获状态码。
func (r *responseRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

// Write 捕获写入的字节数。
func (r *responseRecorder) Write(p []byte) (int, error) {
	if r.status == 0 {
		r.status = http.StatusOK
	}
	n, err := r.ResponseWriter.Write(p)
	r.written += int64(n)
	return n, err
}

// NewLoggingMiddleware 记录每个请求和响应的中间件
func NewLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := &responseRecorder{ResponseWriter: w}
		next.ServeHTTP(rec, r)

		log.Printf("Request: %s %s %s\n", r.Method, r.URL.Path, r.Proto)
		log.Printf("Response: %d %s, %d bytes\n\n\n", rec.status, http.StatusText(rec.status), rec.written)

		// 打印查询参数
		//queryParams := r.URL.Query()
		//log.Printf("Query Params: %v\n", queryParams)
		//
		//// 读取请求体
		//var bodyBytes []byte
		//if r.Body != nil {
		//	bodyBytes, _ = ioutil.ReadAll(r.Body)
		//}
		//
		//// 重要：读取完请求体后，需要将其内容重新放回 r.Body 中，以便后续处理程序可以再次读取
		//r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		//
		//// 打印请求体（假设是字符串）
		//log.Printf("Request Body: %s\n", string(bodyBytes))
		//
		//// 创建一个响应记录器来捕获状态和写入的字节数
		//rec := &responseRecorder{ResponseWriter: w}
		//// 将请求传递给下一个处理程序
		//next.ServeHTTP(rec, r)
		//
		//// 记录请求和响应
		//log.Printf("Request: %s %s %s\n", r.Method, r.URL.Path, r.Proto)
		//log.Printf("Response: %d %s, %d bytes\n\n\n", rec.status, http.StatusText(rec.status), rec.written)
	})
}
