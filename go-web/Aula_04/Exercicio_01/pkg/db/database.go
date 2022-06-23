package db

import (
	"database/sql"
	"log"
)

var (
	StorageDB *sql.DB
)

func init() {
	dataSource := "root:AM@U3c=62@tcp(localhost:3306)/storage"
	var err error
	StorageDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}
	log.Println("database configured")
}
