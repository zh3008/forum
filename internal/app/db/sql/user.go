package sql

import (
	"database/sql"
	"log"

	"github.com/zh3008/forum.git/internal/app/db"
	"github.com/zh3008/forum.git/internal/app/model"
)

//UserRepository ..
type UserRepository struct {
	store *Store
}

//CreateTableUser ..
func (r *UserRepository) CreateTableUser() error {

	tb, err := r.store.db.Prepare(`
		CREATE TABLE IF NOT EXISTS user(
		user_id INTEGER PRIMARY KEY, 
		username TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
	`)

	if err != nil {
		return err
	}

	defer tb.Close()

	if _, err := tb.Exec(); err != nil {
		return err
	}
	log.Println("Table user created")

	return nil
}

//CreateUser ..
func (r *UserRepository) CreateUser(user *model.User) error {
	// if err := user.Validate(); err != nil {
	// 	return err
	// }
	if err := user.BeforeCreate(); err != nil {
		return err
	}
	return r.store.db.QueryRow(
		`INSERT INTO user ( username, email, password)
		VALUES(?,?,?,?,?) 
		RETURNING user_id`,

		user.Username,
		user.Email,
		user.Password,
	).Scan(&user.ID)

}

//FindByUserID ..
func (r *UserRepository) FindByUserID(userID int) (*model.User, error) {
	user := &model.User{}

	if err := r.store.db.QueryRow(
		`SELECT user_id, username, email, password
		FROM user 
		WHERE user_id=?`,
		userID,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, db.ErrRecordNotFound
		}
	}

	return user, nil
}

//FindByUsername ..
func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	user := &model.User{}

	if err := r.store.db.QueryRow(
		`SELECT user_id, username, email, password
		FROM user 
		WHERE username=?`,
		username,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, db.ErrRecordNotFound
		}
	}

	return user, nil
}

//FindByEmail ..
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}

	if err := r.store.db.QueryRow(
		`SELECT user_id, username, email, password
		FROM user 
		WHERE email=?`,
		email,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, db.ErrRecordNotFound
		}
	}

	return user, nil
}
