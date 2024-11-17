package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	DB *sql.DB
}

var dbConn = &DB{}

func ConnectDB(dns string) (*DB, error) {
	conn, err := sql.Open("sqlite3", dns)
	if err != nil {
		log.Fatal("Failed to connect to database: %v", err)
	}

	err = testDB(conn)
	if err != nil {
		return nil, err
	}
	// schema := `
	// 	CREATE TABLE IF NOT EXISTS anime (
	// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 		name TEXT NOT NULL UNIQUE
	// );`
	dbConn.DB = conn
	return dbConn, nil
}

func testDB(conn *sql.DB) error {
	err := conn.Ping()
	if err != nil {
		fmt.Println("Error", err)
		return err
	}

	fmt.Println("***Pinged database succesfully! ***")
	return nil

}
