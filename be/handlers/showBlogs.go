package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"spider/db"
)

// Go 的结构体字段需要是大写字母开头的，否则无法被 JSON 编码器（json.NewEncoder）访问到
type Blog struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func ShowBlogsHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-spider")
	a := session.Values["account"]

	conn, err := db.ConnectDB()
	if err != nil {
		http.Error(w, "无法连接数据库", http.StatusInternalServerError)
		return
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT id, title, content FROM t_article WHERE author = $1", a)
	if err != nil {
		http.Error(w, "数据库查询出错", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var blogs []Blog
	for rows.Next() {
		var blog Blog

		if err = rows.Scan(&blog.Id, &blog.Title, &blog.Content); err != nil {
			http.Error(w, "数据库查询出错", http.StatusInternalServerError)
			return
		}

		blogs = append(blogs, blog)
	}

	// 检查是否存在查询错误
	if err := rows.Err(); err != nil {
		http.Error(w, "读取数据库结果出错", http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(map[string]interface{}{
		"status": 0,
		"blogs":  blogs,
	})
	if err != nil {
		fmt.Fprintf(w, `{"status": 1, "msg": "展示出错"}`)
		return
	}
}
