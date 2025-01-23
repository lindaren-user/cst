package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"spider/db"

	"github.com/jackc/pgx/v5/pgtype"
)

type Logining struct {
	Account string `json:"account"`
	Pwd     string `json:"pwd"`
	Role    string `json:"role"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// 获取数据库连接
	conn, err := db.ConnectDB()
	if err != nil {
		http.Error(w, fmt.Sprintf("无法连接数据库, %s", err.Error()), http.StatusInternalServerError)
		return
	}
	defer conn.Close(context.Background())

	var logining Logining
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&logining)
	if err != nil {
		http.Error(w, fmt.Sprintf("错误的json:, %s", err.Error()), http.StatusBadRequest)
		return
	}

	// 执行 SQL 查询验证用户
	var account, role pgtype.Text
	err = conn.QueryRow(context.Background(), "SELECT account, category FROM t_user WHERE account=$1 AND category=$2 AND user_token=crypt($3, user_token)", logining.Account, logining.Role, logining.Pwd).Scan(&account, &role)
	if err != nil {
		// 返回错误信息
		fmt.Fprintf(w, `{"status":1,"msg":"用户名、密码或身份错误"}`)
		return
	}

	// 如果会话不存在，会自动生成一个新的会话 ID
	session, err := store.Get(r, "session-spider")
	if err != nil {
		http.Error(w, fmt.Sprintf("无法获取会话: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	// 使用字符串存储账号和角色，而不是 pgtype.Text
	session.Values["account"] = logining.Account
	session.Values["role"] = logining.Role

	// 使用了 gorilla/sessions，不需要手动设置 http.Cookie
	// Set-Cookie: session-spider=MTczNjk5Njg1NXx...; Path=/; HttpOnly; Secure
	// 使用 session.Save(req, w) 保存会话数据，这时 gorilla/sessions 会自动在响应中设置 session-spider Cookie
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, fmt.Sprintf("无法保存会话, %s", err.Error()), http.StatusInternalServerError)
		return
	}

	// 返回成功响应
	fmt.Fprintf(w, `{"status":0,"msg":"登录成功"}`)
}
