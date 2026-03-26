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
	var books []table.Book
	for i := 0; i <= 45; i += 5 {
		SQLQuery := "SELECT * FROM books LIMIT 5 OFFSET $1;"
		var rows pgx.Rows
		rows, err = conn.Query(ctx, SQLQuery, i)
		if err != nil {
			fmt.Println(err)
		}
		for rows.Next() {
			var book table.Book
			err = rows.Scan(
				&book.ID,
				&book.Name,
				&book.Author,
				&book.Review,
				&book.Year,
				&book.IsRead,
				&book.AddedAt,
				&book.ReadAt,
			)
			if err != nil {
				fmt.Println(err)
			}
			books = append(books, book)
		}
		pp.Println(books)
		books = []table.Book{}
	}
}
