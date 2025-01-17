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
		http.Error(w, "无法连接数据库", http.StatusInternalServerError)
		return
	}
	defer conn.Close(context.Background())

	path := r.URL.Path
	parts := strings.Split(path, "/")

	id := parts[len(parts)-1]

	_, err = conn.Exec(context.Background(), "DELETE FROM t_article WHERE id = $1", id)
	if err != nil {
		http.Error(w, "数据库操作失败", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, `{"status": 0, "msg": "删除成功"}`)
}
