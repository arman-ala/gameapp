package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLDB struct {
	db *sql.DB
}

func New() *MySQLDB {
	connection, err := sql.Open("mysql", "root:gameappRoo7t0lk2o20@(localhost:3308)/gameapp_db")
	if err != nil {
		panic(fmt.Errorf("can not open mysql db: %w", err))
	}

	connection.SetConnMaxLifetime(time.Minute * 3)
	connection.SetMaxOpenConns(10)
	connection.SetMaxIdleConns(10)

	return &MySQLDB{db: connection}
}
