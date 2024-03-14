package database

import (
	"database/sql"

	"udemy-unit-testing-go/entity"
)

type AttemptHistory interface {
	IncrementFailure(user entity.User) error
	CountFailures(user entity.User) (int, error)
}

type AttemptHistoryRepository struct {
	db *sql.DB
}

func NewAttemptHistoryRepository(db *sql.DB) *AttemptHistoryRepository {
	return &AttemptHistoryRepository{
		db: db,
	}
}

func (dc *AttemptHistoryRepository) IncrementFailure(user entity.User) error {
	sql := "INSERT INTO attempt_history (user_id) VALUES (?)"
	if _, err := dc.db.Exec(sql, user.ID); err != nil {
		return err
	}

	return nil
}

func (dc *AttemptHistoryRepository) CountFailures(user entity.User) (count int, err error) {
	sql := "SELECT count(user_id) FROM attempt_history WHERE user_id = ?"
	row := dc.db.QueryRow(sql)
	if err = row.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}
