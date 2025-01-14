package handlers

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
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
	n := req.URL.Query().Get("name")
	c := req.URL.Query().Get("pwd")

	// 执行 SQL 查询验证用户
	var name, cert string
	err = conn.QueryRow(context.Background(), "SELECT account, user_token FROM t_user WHERE account=$1 AND user_token=crypt($2, user_token)", n, c).Scan(&name, &cert)
	if err != nil {
		// 返回错误信息
		http.Error(w, fmt.Sprintf("验证失败, %s", url.QueryEscape(err.Error())), http.StatusUnauthorized)
		return
	}

	// 获取会话
	session, _ := store.Get(req, "session-spider")

	// 更新会话中的用户名和密码
	session.Values["name"] = n
	session.Values["pwd"] = c

	// 将会话数据保存到响应中，同时把会话 ID 存储到浏览器的 cookie 中
	err = session.Save(req, w)
	if err != nil {
		http.Error(w, fmt.Sprintf("保存会话失败, %s", err.Error()), http.StatusInternalServerError)
		return
	}

	// 返回成功响应
	fmt.Fprintf(w, `{"status":0,"msg":"登录成功"}`)
}
