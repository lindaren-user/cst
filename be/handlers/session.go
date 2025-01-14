package handlers

import (
	"github.com/gorilla/sessions"
)

// 用于在客户端和服务器之间存储和管理会话数据
var store = sessions.NewCookieStore([]byte("alongstory"))
