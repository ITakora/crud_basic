package models

import (
	"crud_basic/db"
	"errors"
)

type User struct {
	ID          int64
	Name        string `binding:"required"`
	Age         int    `binding:"required"`
	Deskription string `binding:"required"`
}

func (user *User) Delete() {

}

func (user *User) Update(parameterUrl string) error {
	query := `UPDATE users SET name = ?, age = ?, deskription = ? WHERE name = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(&user.Name, &user.Age, &user.Deskription, &parameterUrl)

	return err

}

func GetAllUsers() ([]User, error) {
	query := `SELECT * FROM users`
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Deskription)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (user *User) Save() error {
	query := `INSERT INTO users (name, age, deskription) VALUES (?, ?, ?) `

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return errors.New("error")
	}

	defer stmt.Close()
	_, err = stmt.Exec(user.Name, user.Age, user.Deskription)

	return err

}
