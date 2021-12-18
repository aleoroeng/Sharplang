package service

import (
	"database/sql"
	"fmt"
	"sharplang/src/model"
)

type Users struct {
	model.Users
}

func NewUsers() *Users {
	var user Users
	return &user
}

func (u *Users) GetAllUsers(tx *sql.Tx) []*model.Users {

	var users []*model.Users
	var user model.Users
	defer tx.Rollback()
	rows, rowsErr := tx.Query("SELECT * from users")

	if rowsErr != nil {
		return nil
	}

	for rows.Next() {
		rows.Scan(&user.ID, &user.DateUserCreated, &user.DateUserLastUpdated, &user.Username, &user.Password, &user.FirstName, &user.LastName)
		users = append(users, &user)
	}

	if txError := tx.Commit(); txError != nil {
		return nil
	}

	return users
}

//Fetches user from users table by id, returns user if found or nil if no row was found
func (u *Users) GetUsersById(tx *sql.Tx, id int) *model.Users {
	var user model.Users
	defer tx.Rollback()

	row := tx.QueryRow("SELECT * from users WHERE id = $1", id)

	//row.Scan returns exactly one row, columns are served in alphabetical order, so arguments must match up for given column
	errNoRows := row.Scan(&user.ID, &user.DateUserCreated, &user.DateUserLastUpdated, &user.FirstName, &user.LastName, &user.Password, &user.Username)

	if errNoRows != nil {
		fmt.Println(errNoRows)
		return nil
	}

	if err := tx.Commit(); err != nil {
		return nil
	}
	fmt.Println(user.FirstName)
	return &user
}
