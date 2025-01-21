package handlers

import (
	"context"
	"fmt"
	"net/http"
	"spider/db"
	"strings"
)

func DeleteBlogHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := db.ConnectDB()
	if err != nil {
		http.Error(w, fmt.Sprintf("无法连接数据库, %s", err.Error()), http.StatusInternalServerError)
		return
	}
	defer conn.Close(context.Background())

	path := r.URL.Path
	parts := strings.Split(path, "/")

	id := parts[len(parts)-1]

	var a string
	err = conn.QueryRow(context.Background(), "SELECT author FROM t_article WHERE id = $1", id).Scan(&a)
	if err != nil {
		http.Error(w, fmt.Sprintf("数据库操作失败: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	session, _ := store.Get(r, "session-spider")
	if a != session.Values["account"] {
		if session.Values["role"] != "admin" {
			fmt.Fprintf(w, `{"status": 1, "msg": "没有权限"}`)
			return
		}
	}

	_, err = conn.Exec(context.Background(), "DELETE FROM t_article WHERE id = $1", id)
	if err != nil {
		http.Error(w, fmt.Sprintf("数据库操作失败: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, `{"status": 0, "msg": "删除成功"}`)
}
