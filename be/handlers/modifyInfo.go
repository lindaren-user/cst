package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"spider/db"
	"strings"
)

type ModifyingUser struct {
	Account string `json:"account"`
	Pwd     string `json:"pwd"`
}

func ModifyInfoHandler(w http.ResponseWriter, r *http.Request) {
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

	var modifyingUser ModifyingUser

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&modifyingUser)
	if err != nil {
		http.Error(w, "错误的json", http.StatusBadRequest)
		return
	}

	_, err = conn.Exec(context.Background(), "UPDATE t_user SET account = $1, user_token = crypt($2, gen_salt('bf')) WHERE id = $3", modifyingUser.Account, modifyingUser.Pwd, id)
	if err != nil {
		http.Error(w, fmt.Sprintf("数据库操作失败: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, `{"status": 0, "msg": "修改成功"}`)
}
