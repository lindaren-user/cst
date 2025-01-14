package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func CheckAccountHandler(w http.ResponseWriter, req *http.Request) {
	// 获取用户名
	account := req.URL.Query().Get("account")

	// 检查用户名是否存在（用 pgx 连接 PostgreSQL）
	conn, err := pgx.Connect(context.Background(), "postgres://testuser:123456@localhost:5432/testdb")
	if err != nil {
		http.Error(w, "无法连接数据库", http.StatusInternalServerError)
		return
	}
	defer conn.Close(context.Background())

	var count int
	err = conn.QueryRow(context.Background(), "SELECT COUNT(*) FROM t_user WHERE account=$1", account).Scan(&count)
	if err != nil {
		fmt.Fprintf(w, `{"status":1,"msg":"查询失败: %s"}`, err.Error())
		return
	}

	if count > 0 {
		fmt.Fprintln(w, `{"status":1,"msg":"用户名重复"}`)
	} else {
		fmt.Fprintln(w, `{"status":0,"msg":"用户名允许"}`)
	}
}
