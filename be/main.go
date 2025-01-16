package main

import (
	"fmt"
	"net/http"
	"spider/handlers"
)

func main() {
	// 注册身份验证中间件
	http.Handle("/api/changePwd", handlers.AuthMiddleware(http.HandlerFunc(handlers.ChangePwdHandler)))
	http.Handle("/api/createBlogs", handlers.AuthMiddleware(http.HandlerFunc(handlers.CreateBlogsHandler)))

	// 其他公共路由
	http.HandleFunc("/api/login", handlers.LoginHandler)
	http.HandleFunc("/api/register", handlers.RegisterHandler)
	http.HandleFunc("/api/logout", handlers.LogoutHandler)
	http.HandleFunc("/api/checkAccount", handlers.CheckAccountHandler)

	// 启动服务
	fmt.Println("Server is starting on http://localhost:6616...")

	// 这里使用默认的路由器
	http.ListenAndServe(":6616", nil)
}
