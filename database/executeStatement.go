package database

import (
	"database/sql"
)

func ExecuteStatement(query string, args ...any) (sql.Result, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	statement, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	
	result, err := statement.Exec(args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func teste (){
	
}