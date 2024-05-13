package models

import (
	"crud_basic/db"
	"errors"
	"fmt"
)

type User struct {
	ID          int64
	Name        string `binding:"required"`
	Age         int    `binding:"required"`
	Deskription string `binding:"required"`
}

func (user *User) Delete() {

}

func (user *User) Update() {

}

func (user *User) Save() error {
	query := `INSERT INTO users (name, age, deskription) VALUES (?, ?, ?) `

	stmt, err := db.DB.Prepare(query)
	fmt.Println(stmt)
	if err != nil {
		return errors.New("error")
	}

	defer stmt.Close()
	_, err = stmt.Exec(user.Name, user.Age, user.Deskription)

	return err

}
