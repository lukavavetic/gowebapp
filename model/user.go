package model

import (
	"database/sql"
	"fmt"
)

type User struct {
	Id        int
	Email     string
	FirstName string
}

func GetUserById() (*User, error) {
	result := &User{}

	row := db.QueryRow(`SELECT id, email, firstname FROM users WHERE id = 1`)

	err := row.Scan(&result.Id, &result.Email, &result.FirstName)

	switch {
		case err == sql.ErrNoRows:
			return nil, fmt.Errorf("user not found")
		case err != nil:
			return nil, err
	}

	return result, nil
}