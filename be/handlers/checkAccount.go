package handlers

import (
	"context"
	"fmt"
	"net/http"
	"spider/db"
)

func CheckAccountHandler(w http.ResponseWriter, req *http.Request) {
	// 获取用户名
	account := req.URL.Query().Get("account")

	// 检查用户名是否存在（用 pgx 连接 PostgreSQL）
	conn, err := db.ConnectDB()
	if err != nil {
		http.Error(w, fmt.Sprintf("无法连接数据库, %s", err.Error()), http.StatusInternalServerError)
		return
	}
	defer conn.Close(context.Background())

	var count int
	err = conn.QueryRow(context.Background(), "SELECT COUNT(*) FROM t_user WHERE account=$1", account).Scan(&count)
	if err != nil {
		http.Error(w, fmt.Sprintf("数据库操作失败: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	if count > 0 {
		fmt.Fprintln(w, `{"status":1,"msg":"用户名已存在"}`)
	} else {
		fmt.Fprintln(w, `{"status":0,"msg":"用户名允许"}`)
	}
}
