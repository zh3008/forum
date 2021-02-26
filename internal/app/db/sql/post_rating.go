package sql

import (
	"log"

	"github.com/zh3008/forum.git/internal/app/model"
)

//PostRatingRepository ..
type PostRatingRepository struct {
	store *Store
}

//CreateTablePostRating ..
func (r *PostRatingRepository) CreateTablePostRating() error {

	tb, err := r.store.db.Prepare(`
		CREATE TABLE IF NOT EXISTS post_rating(
		post_rating_id INTEGER PRIMARY KEY, 
		post_rating_type TEXT NOT NULL,
		user_id INTEGER NOT NULL,
		post_id INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES user(user_id) ON DELETE CASCADE,
		FOREIGN KEY (post_id) REFERENCES post(post_id) ON DELETE CASCADE
	`)

	if err != nil {
		return err
	}

	defer tb.Close()

	if _, err := tb.Exec(); err != nil {
		return err
	}
	log.Println("Table post_rating created")

	return nil
}

//CreatePostRating ..
func (r *PostRatingRepository) CreatePostRating(postRating *model.PostRating) error {
	return r.store.db.QueryRow(
		`INSERT INTO post_rating (post_rating_type, post_id, user_id)
		VALUES(?,?,?) 
		RETURNING post_raiting_id`,
		postRating.RatingType,
		postRating.PostID,
		postRating.UserID,
	).Scan(&postRating.ID)
}
