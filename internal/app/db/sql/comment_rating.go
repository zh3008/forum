package sql

import (
	"log"

	"github.com/zh3008/forum.git/internal/app/model"
)

//CommentRatingRepository ..
type CommentRatingRepository struct {
	store *Store
}

//CreateTableCommentRating ..
func (r *CommentRatingRepository) CreateTableCommentRating() error {

	tb, err := r.store.db.Prepare(`
		CREATE TABLE IF NOT EXISTS comment_rating(
		comment_rating_id INTEGER PRIMARY KEY, 
		comment_rating_type TEXT NOT NULL,
		user_id INTEGER NOT NULL,
		comment_id INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES user(user_id) ON DELETE CASCADE,
		FOREIGN KEY (comment_id) REFERENCES comment(comment_id) ON DELETE CASCADE
	`)

	if err != nil {
		return err
	}

	defer tb.Close()

	if _, err := tb.Exec(); err != nil {
		return err
	}
	log.Println("Table comment_rating created")

	return nil
}

//CreateCommentRating ..
func (r *CommentRatingRepository) CreateCommentRating(commentRating *model.CommentRating) error {
	return r.store.db.QueryRow(
		`INSERT INTO comment_rating (comment_rating_type, comment_id, user_id)
		VALUES(?,?,?) 
		RETURNING comment_raiting_id`,
		commentRating.RatingType,
		commentRating.CommentID,
		commentRating.UserID,
	).Scan(&commentRating.ID)
}
