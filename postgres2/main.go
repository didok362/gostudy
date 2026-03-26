package main

import (
	"context"
	"fmt"
	"postgres2/db"
	"postgres2/table"

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
	// for i := 1; i <= 50; i++ {
	// 	book := table.Book{
	// 		Name:    fmt.Sprintf("Книга номер %d", i),
	// 		Author:  fmt.Sprintf("Автор %d", i),
	// 		Year:    1900 + i,
	// 		IsRead:  false,
	// 		AddedAt: time.Now(),
	// 		Review:  nil,
	// 		ReadAt:  nil,
	// 	}
	// 	table.InsertBook(ctx, conn, book)
	// }
	// fmt.Println("Добавлено 50 книг!")
	table.ListPages(ctx, conn, 5) // Выведет по 5 книг на страницу
}
