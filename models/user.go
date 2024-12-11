package models

import (
	db "notes.com/app/database"
	"notes.com/app/utils"
)

type User struct {
	Id       int64
	Email    string
	Password string
}

func GetUsers() ([]User, error) {
	query := `SELECT id, email FROM users`

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Email)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, err
}
func GetUserById(userId int64) (*User, error) {
	query := `SELECT id, email FROM users WHERE id = ?`

	row := db.DB.QueryRow(query, userId)

	var user User
	err := row.Scan(&user.Id, &user.Email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (user User) CreateUser() error {
	query := `INSERT INTO users (email, password) VALUES (?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	defer stmt.Close()
	res, err := stmt.Exec(&user.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := res.LastInsertId()
	user.Id = userId

	return err
}

func (user *User) LoginUser() error {
	query := `SELECT id, password FROM users WHERE email = ?`

	row := db.DB.QueryRow(query, user.Email)

	var retrievedPassword string
	err := row.Scan(&user.Id, &retrievedPassword)

	if err != nil {
		return err
	}

	err = utils.ComparePassword(user.Password, retrievedPassword)

	return err
}
