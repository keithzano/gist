package models

import (
	"database/sql"
	"time"
)

type Gist struct {
	ID      int
	Tittle  string
	Content string
	Created time.Time
	Expires time.Time
}

type GistModel struct {
	DB *sql.DB
}

func (m GistModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, nil
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return int(id), nil
}

func (m GistModel) Get(id int) (*Gist, error) {
	return nil, nil
}

func (m *GistModel) Latest() ([]*Gist, error) {
	return nil, nil
}
