package main

import (
	"context" // 用于在 Go 中传递上下文（如超时控制、取消信号等）
	"fmt"
	"net/http" // 用于处理 HTTP 请求和响应
	"net/url"  // 用于解析和处理 URL

	"github.com/gorilla/sessions" // 一个流行的库，用于处理会话（Session）
	"github.com/jackc/pgx/v5"     // 用于连接和操作 PostgreSQL 数据库
)

// 用于在客户端和服务器之间存储和管理会话数据
var store = sessions.NewCookieStore([]byte("alongstory"))

func handler(w http.ResponseWriter, req *http.Request) {
	// 这里需要开启本机pgsql
	connUrl := "postgres://testuser:123456@localhost:5432/testdb"
	conn, err := pgx.Connect(context.Background(), connUrl)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Fprintf(w, `{"status":-1,"msg":"无法连接数据库, %s"}`, url.QueryEscape(err.Error()))
		return
	}
	defer conn.Close(context.Background())

	var name, cert string
	n := req.URL.Query().Get("name")
	c := req.URL.Query().Get("pwd")

	session, _ := store.Get(req, "session-spider")

	if n == "" {
		n, _ = session.Values["name"].(string)
	}

	if c == "" {
		c, _ = session.Values["pwd"].(string)
	}

	if n != "" {
		session.Values["name"] = n
	}

	if c != "" {
		session.Values["pwd"] = c
	}

	// 将会话数据保存到响应中
	session.Save(req, w)

	// 使用 conn.QuerRow 执行一个 SQL 查询
	// 如果查询成功，返回用户名 (name) 和认证令牌 (cert)
	err = conn.QueryRow(context.Background(), "selyect account,user_token from t_user where account=$1 and 		user_token=crypt($2,user_token)", n, c).Scan(&name, &cert)
	if err != nil {
		fmt.Fprintf(w, `{"status":-1,"msg":"%s"}`, url.QueryEscape(err.Error()))
		return
	}

	fmt.Fprintf(w, `{"status":0,"msg":"%s"}`, "成功")
}
func main() {
	http.HandleFunc("/api/login", handler)

	// 打印服务器启动提示
	fmt.Println("Server is starting on http://localhost:6616...")

	// 在本机的6616端口之下，是这个http服务
	http.ListenAndServe(":6616", nil) // nil：使用默认的 ServeMux 路由器来分发请求
}
