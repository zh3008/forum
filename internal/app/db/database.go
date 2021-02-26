package db

import "database/sql"

//Database ..
type Database struct {
	// config *Config
	db *sql.DB
}

//NewDB ..
func NewDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// CreateTables ..
func CreateTables(db *Database) error {
	return nil
}
