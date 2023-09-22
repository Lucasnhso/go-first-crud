package userRepositories

import (
	"crud/database"
	"crud/entities"
	"errors"
)

func FindById(id uint64) (entities.User, error) {
	var user entities.User
	db, err := database.Connect()
	if err != nil {
		return user, err
	}
	defer db.Close()

	row, err := database.ExecuteQuery("select * from users where id = ?", id)
	if err != nil {
		return user, err
	}
	defer row.Close()

	if row.Next() {
		if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return user, err
		}
	} else {
		return user, errors.New("Failed to find users")
	}

	return user, nil
}