package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	connectString := "root:1234@tcp(localhost)/go-crud"

	db, err := sql.Open("mysql", connectString)
	if err != nil {
		return nil, err
	} 
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
