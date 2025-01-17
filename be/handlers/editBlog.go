package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"spider/db"
	"strings"
)

func EditBlogHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := db.ConnectDB()
	if err != nil {
		http.Error(w, "无法连接数据库", http.StatusInternalServerError)
		return
	}
	defer conn.Close(context.Background())

	path := r.URL.Path
	parts := strings.Split(path, "/")

	id := parts[len(parts)-1]

	var article Article

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&article)
	if err != nil {
		http.Error(w, "错误的json", http.StatusBadRequest)
		return
	}

	_, err = conn.Exec(context.Background(), "UPDATE t_article SET title = $1, content = $2 WHERE id = $3", article.Title, article.Content, id)
	if err != nil {
		http.Error(w, "数据库操作失败", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, `{"status": 0, "msg": "修改成功"}`)
}
