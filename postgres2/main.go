package main

import (
	"context"
	"fmt"
	"postgres2/db"
	"postgres2/table"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/k0kubun/pp"
)

func PrintBooks(ctx context.Context, conn *pgx.Conn) {
	books, err := table.ScanBooks(ctx, conn)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	pp.Println(books)
}

func main() {
	connStr := "postgres://postgres:123@localhost:5432/postgres"
	ctx := context.Background()
	conn, err := db.Connect(ctx, connStr)
	if err != nil {
		fmt.Println("Ошибка подключения:", err)
		return
	}
	fmt.Println("Подключено!")
	table.CreateTable(ctx, conn)
	BookToUpdate := table.Book{
		ID:      1, // ID существующей книги
		Name:    "Новое имя",
		Author:  "Новый автор",
		Review:  nil,
		Year:    2025,
		IsRead:  true,
		AddedAt: time.Now(),
		ReadAt:  nil,
	}

	table.UpdateBook(ctx, conn, BookToUpdate)
	PrintBooks(ctx, conn)
	table.DeleteBooks(ctx, conn, []int{2, 3})
}
