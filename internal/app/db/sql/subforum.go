package sql

import (
	"log"

	"github.com/zh3008/forum.git/internal/app/model"
)

//SubforumRepository ..
type SubforumRepository struct {
	store *Store
}

//CreateTableSubforum ..
func (r *SubforumRepository) CreateTableSubforum() error {

	tb, err := r.store.db.Prepare(`
		CREATE TABLE IF NOT EXISTS subforum(
		subforum_id INTEGER PRIMARY KEY, 
		subforum_name TEXT NOT NULL,
		
	`)

	if err != nil {
		return err
	}

	defer tb.Close()

	if _, err := tb.Exec(); err != nil {
		return err
	}
	log.Println("Table subforum created")

	return nil
}

//CreateSubforum ..
func (r *SubforumRepository) CreateSubforum(subforum *model.Subforum) error {
	return r.store.db.QueryRow(
		`INSERT INTO subforum (subforum_name)
		VALUES(?) 
		RETURNING subforum_id`,
		subforum.Name,
	).Scan(&subforum.ID)
}

//SelectSubForumByID ..
func (r *SubforumRepository) SelectSubForumByID(subforumID int) (*model.Subforum, error) {
	subforum := &model.Subforum{}

	if err := r.store.db.QueryRow(
		`SELECT subforum_id, subforum_name,  
		FROM subforum
		WHERE subforum_id=?`,
		subforumID,
	).Scan(
		&subforum.ID,
		&subforum.Name,
	); err != nil {
		return nil, err
	}

	return subforum, nil
}

//SelectSubForumByName ..
func (r *SubforumRepository) SelectSubForumByName(name int) (*model.Subforum, error) {
	subforum := &model.Subforum{}

	if err := r.store.db.QueryRow(
		`SELECT subforum_id, subforum_name,  
		FROM subforum 
		WHERE subforum_name=?`,
		name,
	).Scan(
		&subforum.ID,
		&subforum.Name,
	); err != nil {
		return nil, err
	}

	return subforum, nil
}
