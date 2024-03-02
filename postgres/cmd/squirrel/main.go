package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	dbDSN = "host=localhost port=54321 dbname=go_chat user=igor password=12345 sslmode=disable"
)

func main() {
	ctx := context.Background()

	pool, err := pgxpool.Connect(ctx, dbDSN)

	if err != nil {
		log.Fatal("Error while connecting %v", err)
	}

	defer pool.Close()

	builderInsert := sq.Insert("message").
		PlaceholderFormat(sq.Dollar).
		Columns("from_user", "content").
		Values(gofakeit.Name(), gofakeit.Address().Street).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()

	if err != nil {
		log.Fatal("Error while inserting: %v", err)
	}

	var messageId int

	err = pool.QueryRow(ctx, query, args...).Scan(&messageId)

	if err != nil {
		log.Fatal("Error while scanning: %v", err)
	}

	log.Printf("Scanning id: %d", messageId)


	builderSelect := sq.Select("id", "from_user", "content", "created_at", "updated_at").
		From("message").
		PlaceholderFormat(sq.Dollar).
		OrderBy("id DESC").
		Limit(100)

	query, args, err = builderSelect.ToSql()

	if err != nil {
		log.Fatal("Error while selecting: %v", err)
	}

	rows, err := pool.Query(ctx, query, args...)

	if err != nil {
		log.Fatal("Error while query: %v", err)
	}

	var (
		id int
		fromUser string
		content string
		createdAt time.Time
		updatedAt sql.NullTime
	)

	for rows.Next(){
		err = rows.Scan(&id, &fromUser, &content, &createdAt, &updatedAt)

		if err != nil {
			log.Fatal("Error while scanning v2: %v", err)
		}

		log.Printf("id: %d, from_user: %s, content: %s, created_at: %v, updated_at: %v\n", id, fromUser, content, createdAt, updatedAt)
	}


	builderUpdate := sq.Update("message").
		PlaceholderFormat(sq.Dollar).
		Set("from_user", gofakeit.Name()).
		Set("content", gofakeit.Address().Street).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": messageId})

	query, args, err = builderUpdate.ToSql()

	if err != nil {
		log.Fatal("Error while updating: %v", err)
	}

	res, err := pool.Exec(ctx, query, args...)

	if err != nil {
		log.Fatal("Error while updating: %v", err)
	}

	log.Printf("updated %d rows", res.RowsAffected())


	builderGetOne := sq.Select("id", "from_user", "content", "created_at", "updated_at").
		From("message").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": messageId}).
		Limit(1)

	query, args, err = builderGetOne.ToSql()

	if err != nil {
		log.Fatal("Error while getting one v1: %v", err)
	}

	err = pool.QueryRow(ctx, query, args...).Scan(&id, &fromUser, &content, &createdAt, &updatedAt)

	if err != nil {
		log.Fatal("Error while getting one v2: %v", err)
	}

	log.Printf("id: %d, from_user: %s, content: %s, created_at: %v, updated_at: %v\n", id, fromUser, content, createdAt, updatedAt)
}
