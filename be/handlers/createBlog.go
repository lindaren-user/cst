package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"spider/db"

	"github.com/jackc/pgx/v5/pgtype"
)

// 定义结构体来映射接收到的 JSON 数据
type Article struct {
	Title   pgtype.Text `json:"title"`
	Content pgtype.Text `json:"content"`
}

func CreateBlogsHandler(w http.ResponseWriter, r *http.Request) {
	// 验证请求方法
	// if r.Method != http.MethodPost {
	// 	http.Error(w, "请求方法错误", http.StatusMethodNotAllowed)
	// 	return
	// }

	// 连接数据库
	conn, err := db.ConnectDB()
	if err != nil {
		http.Error(w, fmt.Sprintf("无法连接数据库, %s", err.Error()), http.StatusInternalServerError)
		return
	}
	defer conn.Close(context.Background())

	var article Article

	// 解析json
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&article)
	if err != nil {
		http.Error(w, fmt.Sprintf("错误的json:, %s", err.Error()), http.StatusBadRequest)
		return
	}

	session, _ := store.Get(r, "session-spider")
	a := session.Values["account"]

	_, err = conn.Exec(context.Background(), "INSERT INTO t_article (author, title, content) VALUES ($1, $2, $3)", a, article.Title, article.Content)
	if err != nil {
		http.Error(w, fmt.Sprintf("数据库操作失败: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, `{"status": 0, "msg": "创建成功"}`)
}
