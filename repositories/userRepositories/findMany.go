package userRepositories

import (
	"crud/database"
	"crud/entities"
)

func FindMany() ([]entities.User, error) {
	var users []entities.User
	db, err := database.Connect()
	if err != nil {
		return users, err
	}
	defer db.Close()

	rows, err := database.ExecuteQuery("select * from users")
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user entities.User

		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}