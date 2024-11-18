package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	// "github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	DB *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifetime = 5 * time.Minute

func ConnectDB(dns string) (*DB, error) {
	conn, err := sql.Open("sqlite3", dns)
	if err != nil {
		log.Fatal("Failed to connect to database: %v", err)
	}
	conn.SetMaxOpenConns(maxIdleDbConn)
	conn.SetMaxIdleConns(maxIdleDbConn)
	conn.SetConnMaxLifetime(maxDbLifetime)

	err = testDB(conn)
	if err != nil {
		return nil, err
	}

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
