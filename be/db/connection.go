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
	conn, err := pgx.Connect(context.Background(), connUrl)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}
	return conn, nil
}
