package handlers

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"spider/db"

	"github.com/gorilla/sessions"
)

// 用于在客户端和服务器之间存储和管理会话数据
var store = sessions.NewCookieStore([]byte("alongstory"))

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

	// 获取会话
	session, _ := store.Get(req, "session-spider")

	// 如果用户名为空，尝试从会话中读取
	if n == "" {
		n, _ = session.Values["name"].(string)
	}

	// 如果密码为空，尝试从会话中读取
	if c == "" {
		c, _ = session.Values["pwd"].(string)
	}

	// 更新会话中的用户名和密码
	if n != "" {
		session.Values["name"] = n
	}

	if c != "" {
		session.Values["pwd"] = c
	}

	// 保存会话数据
	err = session.Save(req, w)
	if err != nil {
		http.Error(w, fmt.Sprintf("保存会话失败, %s", err.Error()), http.StatusInternalServerError)
		return
	}

	// 执行 SQL 查询验证用户
	var name, cert string
	err = conn.QueryRow(context.Background(), "SELECT account, user_token FROM t_user WHERE account=$1 AND user_token=crypt($2, user_token)", n, c).Scan(&name, &cert)
	if err != nil {
		// 返回错误信息
		http.Error(w, fmt.Sprintf("验证失败, %s", url.QueryEscape(err.Error())), http.StatusUnauthorized)
		return
	}

	// 返回成功响应
	fmt.Fprintf(w, `{"status":0,"msg":"登录成功"}`)
}
