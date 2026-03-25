package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	connStr := "postgres://postgres:123@localhost:5432/postgres"

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatal("Ошибка подключения:", err)
	}
	defer conn.Close(ctx)

	err = conn.Ping(ctx)
	if err != nil {
		log.Fatal("Ping не прошёл:", err)
	}

	fmt.Println("Подключение успешно!")
}
