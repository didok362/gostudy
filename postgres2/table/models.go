package table

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Book struct {
	ID      int
	Name    string
	Author  string
	Review  *string
	Year    int
	IsRead  bool
	AddedAt time.Time
	ReadAt  *time.Time
}

func InsertBook(ctx context.Context, conn *pgx.Conn, book Book) error {
	SQLstr := `
	INSERT INTO books (name, author, review, year, is_read, added_at, read_at)
    VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := conn.Exec(ctx, SQLstr,
		book.Name,    // $1
		book.Author,  // $2
		book.Review,  // $3
		book.Year,    // $4
		book.IsRead,  // $5
		book.AddedAt, // $6
		book.ReadAt)  // $7

	return err
}

func ScanBooks(ctx context.Context, conn *pgx.Conn) ([]Book, error) {
	var books []Book
	SQLQuerry := "SELECT * FROM books"
	rows, err := conn.Query(ctx, SQLQuerry)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var book Book
		err := rows.Scan(
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
			return nil, err
		}
		books = append(books, book) // Добавляем в срез
	}
	return books, nil
}

func ReadBook(ctx context.Context, conn *pgx.Conn, bookID int) error {
	SQLQuerry := `
	UPDATE books
	SET is_read = true, read_at = NOW()
	WHERE id = $1;
	`
	_, err := conn.Exec(ctx, SQLQuerry, bookID)
	return err
}

func UpdateBook(ctx context.Context, conn *pgx.Conn, book Book) error {
	SQLQuerry := `
	UPDATE books
    SET name = $1, author = $2, review = $3, year = $4, is_read = $5, added_at = $6, read_at = $7
    WHERE id = $8;
	`
	_, err := conn.Exec(ctx, SQLQuerry,
		book.Name,    // $1
		book.Author,  // $2
		book.Review,  // $3
		book.Year,    // $4
		book.IsRead,  // $5
		book.AddedAt, // $6
		book.ReadAt,  // $7
		book.ID)      // $8
	return err
}

func DeleteBooks(ctx context.Context, conn *pgx.Conn, ids []int) error {
	SQLQuerry := `DELETE FROM books WHERE id = ANY($1);`
	_, err := conn.Exec(ctx, SQLQuerry, ids)
	return err
}

// CREATE TABLE IF NOT EXISTS books (
// 		id SERIAL PRIMARY KEY,
// 		name TEXT NOT NULL,
// 		author TEXT NOT NULL,
// 		review TEXT,
// 		year INT NOT NULL,
// 		is_read BOOLEAN NOT NULL DEFAULT false,
// 		added_at TIMESTAMPTZ NOT NULL,
// 		read_at  TIMESTAMPTZ
// 	);
// 	`
