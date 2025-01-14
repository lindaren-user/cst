// db/db.go
package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// ConnectDB 函数返回一个数据库连接
func ConnectDB() (*pgx.Conn, error) {
	connUrl := "postgres://testuser:123456@localhost:5432/testdb"
	conn, err := pgx.Connect(context.Background(), connUrl) // context.Background()：通常用于没有特定上下文的情况下，作为根上下文传递
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}
	return conn, nil
}
