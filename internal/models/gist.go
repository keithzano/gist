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

	gist := &Gist{}

	err := row.Scan(&gist.ID, &gist.Title, &gist.Content, &gist.Created, &gist.Expires)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return gist, nil
}

func (m *GistModel) Latest() ([]*Gist, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
    WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	gists := []*Gist{}

	for rows.Next() {
		g := &Gist{}

		err = rows.Scan(&g.ID, &g.Title, &g.Content, &g.Created, &g.Expires)
		if err != nil {
			return nil, err
		}
		gists = append(gists, g)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return gists, nil
}
