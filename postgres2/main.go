package main

import (
	"context"
	"fmt"
	"postgres2/db"
)

func main() {
	connStr := "postgres://postgres:123@localhost:5432/postgres"
	ctx := context.Background()
	conn, err := db.Connect(ctx, connStr)
	if err != nil {
		fmt.Println("Ошибка подключения:", err)
		return
	}
	fmt.Println("Подключено!", conn)
}
