package handlers

import (
	"fmt"
	"net/http"
)

// 注意有中间件
func GetUserSession(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-spider")
	account := session.Values["account"]
	role := session.Values["role"]

	// 修正了 JSON 字符串的格式
	fmt.Fprintf(w, `{"status": 0, "account": "%s", "role": "%s"}`, account, role)
}
