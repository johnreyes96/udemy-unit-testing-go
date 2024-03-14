package database

import "database/sql"

type BadWords interface {
	FindAll() ([]string, error)
}

type BadWordsRepository struct {
	bd *sql.DB
}

func (dc *BadWordsRepository) FindAll() (badWordList []string, err error) {
	sql := "SELECT name FROM bad_word"
	rows, err := dc.bd.Query(sql)
	if err != nil {
		return nil, err
	}

	var badWord string
	for rows.Next() {
		if err := rows.Scan(&badWord); err != nil {
			return badWordList, err
		}

		badWordList = append(badWordList, badWord)
	}

	return badWordList, nil
}
