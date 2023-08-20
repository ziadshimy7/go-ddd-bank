package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Adapter struct {
	DB *sql.DB
}

func NewMySqlConnection(driverName, dataSourceName string) (*Adapter, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("db connection failure: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping failure: %v", err)
	}

	return &Adapter{DB: db}, nil
}

// CloseDbConnection closes the db  connection
func (da *Adapter) CloseDbConnection() {
	err := da.DB.Close()
	if err != nil {
		log.Fatalf("db close failure: %v", err)
	}
}
