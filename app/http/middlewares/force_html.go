package middlewares

import "net/http"

func ForceHTMLMiddle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 设置头部信息
		w.Header().Set("Content-type", "text/html; charset=utf8")

		// 2. 继续处理
		next.ServeHTTP(w, r)
	})
}
