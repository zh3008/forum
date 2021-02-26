package sql

import (
	"log"
)

//SubforumPostRepository ..
type SubforumPostRepository struct {
	store *Store
}

//CreateTableSubforumPost ..
func (r *SubforumPostRepository) CreateTableSubforumPost() error {

	tb, err := r.store.db.Prepare(`
		CREATE TABLE IF NOT EXISTS subforum_post(
		post_id INTEGER NOT NULL,
		subfoum_id INTEGER NOT NULL,
		FOREIGN KEY (post_id) REFERENCES post(post_id) ON DELETE CASCADE
		FOREIGN KEY (subforum_id) REFERENCES subforum(subforum_id) ON DELETE CASCADE
		
	`)

	if err != nil {
		return err
	}

	defer tb.Close()

	if _, err := tb.Exec(); err != nil {
		return err
	}
	log.Println("Table subforum_post created")

	return nil
}
