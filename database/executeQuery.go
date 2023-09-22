package database

import (
	"database/sql"
	"fmt"
)

func ExecuteQuery(query string, args ...any) (*sql.Rows, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// defer rows.Close()
	
	if err != nil {
		return nil, err
	}

	return rows, nil
}