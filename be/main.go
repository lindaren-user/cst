package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/sessions" // 一个流行的库，用于处理会话（Session）
	"github.com/jackc/pgx/v5"
)

var store = sessions.NewCookieStore([]byte("alongstory"))

func handler(w http.ResponseWriter, req *http.Request) {
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

	session.Save(req, w)

	err = conn.QueryRow(context.Background(), "select account,user_token from t_user where account=$1 and 		user_token=crypt($2,user_token)", n, c).Scan(&name, &cert)
	if err != nil {
		fmt.Fprintf(w, `{"status":-1,"msg":"%s"}`, url.QueryEscape(err.Error()))
		return
	}

	fmt.Fprintf(w, `{"status":0,"msg":"%s"}`, "成功")
}
func main() {
	http.HandleFunc("/api/login", handler)
	http.ListenAndServe(":6616", nil)
}
