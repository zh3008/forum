package sql

import (
	"log"

	"github.com/zh3008/forum.git/internal/app/model"
)

//CommentRepository ..
type CommentRepository struct {
	store *Store
}

//CreateTableComment ..
func (r *CommentRepository) CreateTableComment() error {

	tb, err := r.store.db.Prepare(`
		CREATE TABLE IF NOT EXISTS comment(
		comment_id INTEGER PRIMARY KEY, 
		comment_creation_date TEXT NOT NULL,
		comment_content TEXT NOT NULL,
		user_id INTEGER NOT NULL,
		post_id INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES user(user_id) ON DELETE CASCADE
		FOREIGN KEY (post_id) REFERENCES post(post_id) ON DELETE CASCADE

	`)

	if err != nil {
		return err
	}

	defer tb.Close()

	if _, err := tb.Exec(); err != nil {
		return err
	}
	log.Println("Table comment created")

	return nil
}

//CreateComment ..
func (r *CommentRepository) CreateComment(comment *model.Comment) error {
	return r.store.db.QueryRow(
		`INSERT INTO comment (comment_creation_date, comment_content, user_id, post_id)
		VALUES(?,?,?,?) 
		RETURNING comment_id`,
		comment.CreationDate,
		comment.Content,
		comment.UserID,
	).Scan(&comment.ID)
}

//SelectCommentsByPost ..
func (r *CommentRepository) SelectCommentsByPost(postID int) ([]*model.Comment, error) {
	comment := &model.Comment{}
	comments := []*model.Comment{}

	rows, err := r.store.db.Query(
		`SELECT comment_id, comment_creation_date, comment_content, user_id,  post_id
		FROM comment 
		WHERE post_id=?`,
		postID,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(
			&comment.ID,
			&comment.CreationDate,
			&comment.Content,
			&comment.PostID,
			&comment.UserID,
		)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)

	}

	return comments, nil
}
