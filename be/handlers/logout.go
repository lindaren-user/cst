package handlers

import (
	"fmt"
	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// 获取当前会话
	session, err := store.Get(r, "session-spider")
	if err != nil {
		http.Error(w, "无法获取会话", http.StatusInternalServerError)
		return
	}

	// 删除会话中的所有数据
	session.Values = nil

	// 设置会话过期时间，使其立即失效
	session.Options.MaxAge = -1

	// 保存清除后的会话
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, fmt.Sprintf("无法保存会话, %s", err.Error()), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, `{"status": 0, "msg": "退出成功"}`)
}
