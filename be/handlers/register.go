package handlers

import (
	"context"
	"fmt"
	"net/http"
	"spider/db"
)

func RegisterHandler(w http.ResponseWriter, req *http.Request) {
	// 数据库连接
	conn, err := db.ConnectDB()
	if err != nil {
		http.Error(w, fmt.Sprintf("无法连接数据库, %s", err.Error()), http.StatusInternalServerError)
		return
	}
	defer conn.Close(context.Background())

	a := req.URL.Query().Get("account")
	p := req.URL.Query().Get("pwd")
	n := req.URL.Query().Get("name")
	m := req.URL.Query().Get("mobilePhone")

	// 数据库操作等
	// 执行 SQL 插入操作，有一个表 t_user 来存储用户信息
	_, err = conn.Exec(context.Background(), `
		INSERT INTO t_user (account, user_token, name, mobile_phone, status)
        VALUES ($1, crypt($2, gen_salt('bf')), $3, $4, '00')
    `, a, p, n, m)
	if err != nil {
		http.Error(w, fmt.Sprintf("注册失败: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	//获取会话
	session, _ := store.Get(req, "session-spider")

	//更新会话
	session.Values["account"] = a
	// 密码不应该保存在session
	// session.Values["pwd"] = p
	// 如果是频繁使用，就保存
	// session.Values["name"] = n
	// session.Values["mobilePhone"] = m

	// 将会话数据保存到响应中，同时把会话 ID 存储到浏览器的 cookie 中
	// 会话 ID 会通过 Set-Cookie HTTP 头部返回给浏览器，浏览器会自动存储它。每次后续的请求中，浏览器会自动发送这个 cookie，从而让服务器知道是哪一个用户发起的请求
	err = session.Save(req, w)
	if err != nil {
		http.Error(w, fmt.Sprintf("保存会话失败, %s", err.Error()), http.StatusInternalServerError)
		return
	}

	// 返回成功响应
	fmt.Fprintf(w, `{"status":0,"msg":"注册成功"}`)
}
