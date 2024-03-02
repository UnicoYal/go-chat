package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/jackc/pgx/v5"
)

const (
	dbDSN = "host=localhost port=54321 dbname=go_chat user=igor password=12345 sslmode=disable"
)

func main() {
	ctx := context.Background()

	con, err := pgx.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatal("Failed to connect to database %v", err)
	}

	defer con.Close(ctx)

	res, err := con.Exec(ctx, "INSERT INTO message (from_user, content) VALUES ($1, $2)", gofakeit.Name(), gofakeit.Address().Street)

	if err != nil {
		log.Fatal("Failed to insert into message_table: %v", err)
	}

	log.Printf("inserted %d rows", res.RowsAffected())

	rows, err := con.Query(ctx, "SELECT * FROM message")

	if err != nil {
		log.Fatal("Failed to select from message_table: %v", err)
	}

	defer rows.Close()

	var id int
		var fromUser string
		var content string
		var createdAt time.Time
		var updatedAt sql.NullTime

	for rows.Next() {
		err = rows.Scan(&id, &fromUser, &content, &createdAt, &updatedAt)

		if err != nil {
			log.Fatal("Error while scanning rows %v", err)
		}

		log.Printf("id: %d, from_user: %s, content: %s, created_at: %v, updated_at: %v\n", id, fromUser, content, createdAt, updatedAt)
	}
}
