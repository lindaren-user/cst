package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"spider/db"
)

type NewUser struct {
	Account      string `json:"account"`
	Pwd          string `json:"pwd"`
	Name         string `json:"name"`
	Mobile_Phone string `json:"mobilePhone"`
	Role         string `json:"role"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// 数据库连接
	conn, err := db.ConnectDB()
	if err != nil {
		http.Error(w, fmt.Sprintf("无法连接数据库, %s", err.Error()), http.StatusInternalServerError)
		return
	}
	defer conn.Close(context.Background())

	var newUser NewUser
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newUser)
	if err != nil {
		http.Error(w, fmt.Sprintf("错误的json:, %s", err.Error()), http.StatusBadRequest)
		return
	}

	// 获取会话
	session, _ := store.Get(r, "session-spider")
	if session.Values["role"] != "admin" && newUser.Role == "admin" {
		fmt.Fprintf(w, `{"status":1,"msg":"没有权限"}`)
		return
	}

	// 执行 SQL 插入操作，使用 RETURNING 语句获取新插入行的 ID
	var newID int
	err = conn.QueryRow(context.Background(), `
		INSERT INTO t_user (account, user_token, name, mobile_phone, category, status)
        VALUES ($1, crypt($2, gen_salt('bf')), $3, $4, $5,'00')
        RETURNING id  
    `, newUser.Account, newUser.Pwd, newUser.Name, newUser.Mobile_Phone, newUser.Role).Scan(&newID)

	if err != nil {
		http.Error(w, fmt.Sprintf("数据库操作失败: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	// 更新会话
	// session.Values["account"] = newUser.Account
	// 密码不应该保存在session
	// session.Values["pwd"] = p

	// 将会话数据保存到响应中，同时把会话 ID 存储到浏览器的 cookie 中
	// 会话 ID 会通过 Set-Cookie HTTP 头部返回给浏览器，浏览器会自动存储它。每次后续的请求中，浏览器会自动发送这个 cookie，从而让服务器知道是哪一个用户发起的请求
	// err = session.Save(r, w)
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("保存会话失败, %s", err.Error()), http.StatusInternalServerError)
	// 	return
	// }

	// 返回成功响应，同时包含新插入的 ID
	fmt.Fprintf(w, `{"status":0,"msg":"注册成功","account":%q,"id":%d}`, newUser.Account, newID)
}
