package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	if err := createTodosTable(db); err != nil {
		return nil, err
	}

	return db, nil
}
func createTodosTable(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS todos (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        completed BOOLEAN NOT NULL
    );`
	_, err := db.Exec(query)
	return err
}
