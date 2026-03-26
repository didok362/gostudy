package table

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) {
	SQLstr := `
	CREATE TABLE IF NOT EXISTS books (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		author TEXT NOT NULL,
		review TEXT,
		year INT NOT NULL,
		is_read BOOLEAN NOT NULL DEFAULT false,
		added_at TIMESTAMPTZ NOT NULL,
		read_at  TIMESTAMPTZ
	);
	`
	_, err := conn.Exec(ctx, SQLstr)
	if err != nil {
		fmt.Println(err.Error())
	}
}
