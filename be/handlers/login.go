package handlers

import (
	"context"
	"fmt"
	"net/http"
	"spider/db"
)

func LoginHandler(w http.ResponseWriter, req *http.Request) {
	// 获取数据库连接
	conn, err := db.ConnectDB()
	if err != nil {
		http.Error(w, fmt.Sprintf("无法连接数据库, %s", err.Error()), http.StatusInternalServerError)
		return
	}
	defer conn.Close(context.Background())

	// 从请求 URL 获取用户名和密码
	a := req.URL.Query().Get("name")
	c := req.URL.Query().Get("pwd")

	// 执行 SQL 查询验证用户
	var account, cert string
	err = conn.QueryRow(context.Background(), "SELECT account, user_token FROM t_user WHERE account=$1 AND user_token=crypt($2, user_token)", a, c).Scan(&account, &cert)
	if err != nil {
		// 返回错误信息
		fmt.Fprintf(w, `{"status":1,"msg":"验证失败: %s"}`, err.Error())
		return
	}

	// 如果会话不存在，会自动生成一个新的会话 ID
	session, err := store.Get(req, "session-spider")
	if err != nil {
		http.Error(w, "无法获取会话", http.StatusInternalServerError)
		return
	}

	session.Values["account"] = a

	// 使用了 gorilla/sessions，不需要手动设置 http.Cookie
	// Set-Cookie: session-spider=MTczNjk5Njg1NXx...; Path=/; HttpOnly; Secure
	// 使用 session.Save(req, w) 保存会话数据，这时 gorilla/sessions 会自动在响应中设置 session-spider Cookie
	session.Save(req, w)

	// 返回成功响应
	fmt.Fprintf(w, `{"status":0,"msg":"登录成功"}`)
}
