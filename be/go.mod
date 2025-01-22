module spider // module spider 这行声明了该 Go 项目的模块名为 spider，即该项目的根目录或包名称

go 1.23.4 // 表示使用 Go 1.18 或更高版本

// 记录的是项目的依赖
require (
	github.com/gorilla/sessions v1.4.0
	github.com/jackc/pgx/v5 v5.7.2
)

require (
	github.com/gorilla/securecookie v1.1.2 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	golang.org/x/crypto v0.32.0 // indirect
	golang.org/x/text v0.21.0 // indirect
)
