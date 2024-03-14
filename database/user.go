package database

import (
	"database/sql"
	"udemy-unit-testing-go/entity"
)

type User interface {
	Add(user entity.User) error
}

type UserRepository struct {
	db *sql.DB
}

func (dc *UserRepository) Add(user entity.User) error {
	sql := "INSERT INTO user (name, email, description) VALUES (?, ?, ?)"

	if _, err := dc.db.Exec(sql, user.Name, user.Email, user.Description); err != nil {
		return err
	}

	return nil
}
