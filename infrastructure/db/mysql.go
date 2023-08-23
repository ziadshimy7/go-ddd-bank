package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type DbConnection struct {
	DB *sql.DB
}

func NewMySqlConnection(driverName, dataSourceName string) (*DbConnection, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("db connection failure: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping failure: %v", err)
	}

	return &DbConnection{DB: db}, nil
}

// CloseDbConnection closes the db  connection
func (da *DbConnection) CloseDbConnection() {
	err := da.DB.Close()
	if err != nil {
		log.Fatalf("db close failure: %v", err)
	}
}
