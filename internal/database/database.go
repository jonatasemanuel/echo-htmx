package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	//"github.com/jonatasemanuel/echo-htmx/internal/database"
)

type Anime struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type DB struct {
	conn *sqlx.DB
}

func NewDB(dns string) *DB {
	conn, err := sqlx.Open("sqlite3", dns)
	if err != nil {
		log.Fatal("Failed to connect to database: %v", err)
	}
	defer conn.Close()

	// create tables
	schema := `
		CREATE TABLE IF NOT EXISTS anime (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL UNIQUE
	);`
	if _, err := conn.Exec(schema); err != nil {
		log.Fatal("Failed to create schema: %v", err)
	}

	return &DB{conn: conn}
}

func (db *DB) Close() {
	if err := db.conn.Close(); err != nil {
		log.Printf("Failed to close database connection: %v", err)
	}
}

func (db *DB) CreateAnime(name string) (Anime, error) {
	query := `INSERT INTO anime (anime) VALUES (?) RETURNING id,name`
	var anime Anime
	err := db.conn.QueryRowx(query, name).StructScan(&anime)
	return anime, err
}

func (db *DB) ListAnime() ([]Anime, error) {
	var animes []Anime
	query := `SELECT id, name FROM anime ORDER BY id`
	err := db.conn.Select(&animes, query)
	return animes, err
}
