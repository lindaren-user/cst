package handlers

import (
	"net/http"
)

// 验证用户是否登录的中间件
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 获取当前会话
		session, err := store.Get(r, "session-spider")
		if err != nil || session.Values["account"] == nil {
			// 如果会话不存在或没有存储账户信息，返回 401 未授权
			http.Error(w, "未登录或会话失效", http.StatusUnauthorized)
			return
		}

		// 用户已经登录，继续执行请求处理
		next.ServeHTTP(w, r)
	})
}
