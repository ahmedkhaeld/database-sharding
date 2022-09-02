package main

import (
	"database/sql"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"time"
)

const (
	maxOpenDbConn = 10
	maxDbLifetime = 5 * time.Minute
	maxIdleDbConn = 5
)

type DB struct {
	SQL *sql.DB
}

var Conn = &DB{}

func ConnectPostgres(dsn string) {
	d, err := openPostgres(dsn)
	if err != nil {
		panic(err)
	}

	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifetime)

	Conn.SQL = d

	err = pingDB(d)
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...", err)
	}
	log.Println("Connected to Postgres")

}

// pingDB tries to ping the database
func pingDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}
	return nil
}

// OpenPostgres creates a new database for the application
func openPostgres(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
