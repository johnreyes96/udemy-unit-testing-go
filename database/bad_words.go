package database

import "database/sql"

type BadWords interface {
	findAll() ([]string, error)
}

type BadWordsRepository struct {
	bd *sql.DB
}

func newBadWordsRepository(db *sql.DB) *BadWordsRepository {
	return &BadWordsRepository{bd: db}
}

func (dc *BadWordsRepository) findAll() (badWordList []string, err error) {
	sql := "SELECT name FROM bad_word"
	rows, err := dc.bd.Query(sql)
	if err != nil {
		return err
	}

	var badWord string
	for rows.Next() {
		err := rows.Scan(&badWord)
		if err != nil {
			return badWordList, err
		}

		badWordList = append(badWordList, badWord)
	}

	return badWordList, nil
}
