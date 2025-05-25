package models

import (
	"database/sql"
	"errors"
	"time"
)

type Gist struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type GistModel struct {
	DB *sql.DB
}

func (m *GistModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *GistModel) Get(id int) (*Gist, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
    WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id)

	s := &Gist{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *GistModel) Latest() ([]*Gist, error) {
	return nil, nil
}
