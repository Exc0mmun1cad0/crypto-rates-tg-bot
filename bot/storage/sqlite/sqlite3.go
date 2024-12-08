package sqlite3

import (
	"database/sql"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(path string) (*Storage, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Init() error {
	q := `CREATE TABLE IF NOT EXISTS tokens (username TEXT, token TEXT)`

	_, err := s.db.Exec(q)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Add(username string, tokenList string) error {
	q := `INSERT INTO tokens (username, token) VALUES (?, ?)`

	for _, token := range strings.Fields(tokenList) {
		if _, err := s.db.Exec(q, username, token); err != nil {
			return err
		}
	}

	return nil
}

func (s *Storage) GetAll(username string) ([]string, error) {
	var tokenList []string

	q, _ := s.db.Prepare(`SELECT token FROM tokens WHERE username = ?`)
	rows, _ := q.Query(username)
	for rows.Next() {
		var token string
		rows.Scan(&token)
		tokenList = append(tokenList, token)
	}

	return tokenList, nil
}

func (s *Storage) Delete(username string, tokenList string) error {
	q := `DELETE FROM tokens WHERE username = ? AND token = ?`

	for _, token := range strings.Fields(tokenList) {
		if _, err := s.db.Exec(q, username, token); err != nil {
			return err
		}
	}

	return nil
}

func (s *Storage) DeleteAll(username string) error {
	q := `DELETE FROM tokens WHERE username = ?`
	_, err := s.db.Exec(q, username)
	if err != nil {
		return err
	}

	return nil
}
