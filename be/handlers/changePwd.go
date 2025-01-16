package handlers

import (
	"context"
	"fmt"
	"net/http"
	"spider/db"
)

func ChangePwdHandler(w http.ResponseWriter, req *http.Request) {
	// 数据库连接
	conn, err := db.ConnectDB()
	if err != nil {
		http.Error(w, "无法连接数据库", http.StatusInternalServerError)
		return
	}
	defer conn.Close(context.Background())

	// 获取当前会话
	session, _ := store.Get(req, "session-spider")

	// 验证账号密码
	a := session.Values["account"]
	o := req.URL.Query().Get("oldPwd")

	var account, cert string
	err = conn.QueryRow(context.Background(), "SELECT account, user_token FROM t_user WHERE account=$1 AND user_token=crypt($2, user_token)", a, o).Scan(&account, &cert)
	if err != nil {
		// 返回错误信息
		fmt.Fprintf(w, `{"status":1,"msg":"原密码错误: %s"}`, err.Error())
		return
	}

	n := req.URL.Query().Get("newPwd")

	// 使用 Exec 执行 UPDATE 语句
	_, err = conn.Exec(context.Background(), "UPDATE t_user SET user_token = crypt($1, gen_salt('bf')) WHERE account = $2", n, a)
	if err != nil {
		http.Error(w, "数据库操作失败", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, `{"status":0, "msg":"修改成功"}`)
}
