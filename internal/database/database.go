package database

import (
	"context"
	"database/sql"
	_ "embed"
	"log"
	// "reflect"

	_ "github.com/mattn/go-sqlite3"

	"github.com/jonatasemanuel/echo-htmx/internal/database/db_test"
)

//go:embed schema.sql
var ddl string

func Run() error {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return err
	}

	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return err
	}

	queries := db_test.New(db)

	// list all authors
	animes, err := queries.ListAnimes(ctx)
	if err != nil {
		return err
	}
	log.Println(animes)

	return nil
}
