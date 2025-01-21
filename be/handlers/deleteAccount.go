package handlers

import (
	"context"
	"fmt"
	"net/http"
	"spider/db"
	"strings"
)

func DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-spider")
	if session.Values["role"] != "admin" {
		fmt.Fprintf(w, `{"status": 1, "msg": "没有权限"}`)
		return
	}

	conn, err := db.ConnectDB()
	if err != nil {
		http.Error(w, fmt.Sprintf("无法连接数据库, %s", err.Error()), http.StatusInternalServerError)
		return
	}
	defer conn.Close(context.Background())

	path := r.URL.Path
	parts := strings.Split(path, "/")

	id := parts[len(parts)-1]

	var account string
	err = conn.QueryRow(context.Background(), "DELETE FROM t_user WHERE id = $1 RETURNING account", id).Scan(&account)
	if err != nil {
		http.Error(w, fmt.Sprintf("数据库操作失败1: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	// 顺带删除该用户的博客
	_, err = conn.Exec(context.Background(), "DELETE FROM t_article WHERE author = $1", account)
	if err != nil {
		http.Error(w, fmt.Sprintf("数据库操作失败2: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, `{"status": 0, "msg": "删除成功"}`)
}
