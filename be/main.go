package main

import (
	"fmt"
	"net/http"
	"spider/handlers"
)

func main() {
	// 注册路由
	http.HandleFunc("/api/login", handlers.LoginHandler)
	http.HandleFunc("/api/register", handlers.RegisterHandler)

	// 启动服务
	fmt.Println("Server is starting on http://localhost:6616...")
	http.ListenAndServe(":6616", nil)
}
