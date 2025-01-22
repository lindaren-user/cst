package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"spider/db"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

// Go 的结构体字段需要是大写字母开头的，否则无法被 JSON 编码器（json.NewEncoder）访问到
type Blog struct {
	Id      pgtype.Int4 `json:"id"`
	Author  pgtype.Text `json:"author"`
	Title   pgtype.Text `json:"title"`
	Content pgtype.Text `json:"content"`
}

func ShowAllBlogsHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-spider")
	a := session.Values["account"]
	role := session.Values["role"]

	conn, err := db.ConnectDB()
	if err != nil {
		http.Error(w, fmt.Sprintf("无法连接数据库, %s", err.Error()), http.StatusInternalServerError)
		return
	}
	defer conn.Close(context.Background())

	var rows pgx.Rows
	if role == "user" {
		rows, err = conn.Query(context.Background(), "SELECT id, author, title, content FROM t_article WHERE author = $1", a)
	} else if role == "admin" {
		rows, err = conn.Query(context.Background(), "SELECT id, author, title, content FROM t_article")
	} else {
		fmt.Fprintf(w, `{"status":1,"msg":"没有权限"}`)
		return
	}

	if err != nil {
		http.Error(w, fmt.Sprintf("数据库操作失败: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var blogs []Blog
	for rows.Next() {
		var blog Blog

		if err = rows.Scan(&blog.Id, &blog.Author, &blog.Title, &blog.Content); err != nil {
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
		http.Error(w, fmt.Sprintf("展示出错, %s", err.Error()), http.StatusInternalServerError)
		return
	}
}
