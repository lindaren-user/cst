package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"spider/db"
)

type User struct {
	Id           int    `json:"id"`
	Account      string `json:"account"`
	Name         string `json:"name"`
	Mobile_Phone string `json:"mobile_phone"`
}

func GetAllAccountsHandler(w http.ResponseWriter, r *http.Request) {
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

	rows, err := conn.Query(context.Background(), "SELECT id, account, name, mobile_phone FROM t_user WHERE category = 'user'")
	if err != nil {
		http.Error(w, fmt.Sprintf("数据库操作失败: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err = rows.Scan(&user.Id, &user.Account, &user.Name, &user.Mobile_Phone); err != nil {
			http.Error(w, fmt.Sprintf("读取数据库结果出错: %v", err), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	// 检查是否存在查询错误
	if err := rows.Err(); err != nil {
		http.Error(w, "读取数据库结果出错", http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(map[string]interface{}{
		"status": 0,
		"users":  users,
	})
	if err != nil {
		fmt.Fprintf(w, `{"status": 1, "msg": "展示出错"}`)
		return
	}
}
