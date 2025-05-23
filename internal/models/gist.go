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
	return 0, nil
}

func (m GistModel) Get(id int) (*Gist, error) {
	return nil, nil
}

func (m *GistModel) Latest() ([]*Gist, error) {
	return nil, nil
}
