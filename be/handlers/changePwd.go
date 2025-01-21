package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"spider/db"
)

type PwdPair struct {
	Old string `json:"oldPwd"`
	New string `json:"newPwd"`
}

func ChangePwdHandler(w http.ResponseWriter, r *http.Request) {
	// 数据库连接
	conn, err := db.ConnectDB()
	if err != nil {
		http.Error(w, fmt.Sprintf("无法连接数据库, %s", err.Error()), http.StatusInternalServerError)
		return
	}
	defer conn.Close(context.Background())

	// 获取当前会话
	session, _ := store.Get(r, "session-spider")

	// 验证账号密码
	a := session.Values["account"]

	var pwdPair PwdPair
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&pwdPair)
	if err != nil {
		http.Error(w, fmt.Sprintf("错误的json:, %s", err.Error()), http.StatusBadRequest)
		return
	}

	var account string
	err = conn.QueryRow(context.Background(), "SELECT account FROM t_user WHERE account=$1 AND user_token=crypt($2, user_token)", a, pwdPair.Old).Scan(&account)
	if err != nil {
		// 返回错误信息
		fmt.Fprintf(w, `{"status":1,"msg":"原密码错误"}`)
		return
	}

	// 使用 Exec 执行 UPDATE 语句
	_, err = conn.Exec(context.Background(), "UPDATE t_user SET user_token = crypt($1, gen_salt('bf')) WHERE account = $2", pwdPair.New, a)
	if err != nil {
		http.Error(w, fmt.Sprintf("数据库操作失败: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, `{"status":0, "msg":"修改成功"}`)
}
