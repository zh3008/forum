package sql

import (
	"log"

	"github.com/zh3008/forum.git/internal/app/model"
)

//PostRepository ..
type PostRepository struct {
	store *Store
}

//CreateTablePost ..
func (r *PostRepository) CreateTablePost() error {

	tb, err := r.store.db.Prepare(`
		CREATE TABLE IF NOT EXISTS post(
		post_id INTEGER PRIMARY KEY, 
		post_title TEXT NOT NULL,
		post_creation_date TEXT NOT NULL,
		post_content TEXT NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES user(user_id) ON DELETE CASCADE
	`)

	if err != nil {
		return err
	}

	defer tb.Close()

	if _, err := tb.Exec(); err != nil {
		return err
	}
	log.Println("Table post created")

	return nil
}

//CreatePost ..
func (r *PostRepository) CreatePost(post *model.Post) error {
	return r.store.db.QueryRow(
		`INSERT INTO post (post_title, post_creation_date, post_content, user_id)
		VALUES(?,?,?,?) 
		RETURNING post_id`,
		post.Title,
		post.CreationDate,
		post.Content,
		post.UserID,
	).Scan(&post.ID)
}

//FindByPostID ..
func (r *PostRepository) FindByPostID(postID int) (*model.Post, error) {
	post := &model.Post{}

	if err := r.store.db.QueryRow(
		`SELECT post_id, post_title, post_creation_date, post_content, user_id,  
		FROM post 
		WHERE post_id=?`,
		postID,
	).Scan(
		&post.ID,
		&post.Title,
		&post.CreationDate,
		&post.Content,
		&post.UserID,
	); err != nil {
		return nil, err
	}

	return post, nil
}

//SelectAllPosts ..
func (r *PostRepository) SelectAllPosts() ([]*model.Post, error) {
	post := &model.Post{}
	posts := []*model.Post{}

	rows, err := r.store.db.Query(
		`SELECT post_id, post_title, post_creation_date, post_content, user_id  
		FROM post`,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.CreationDate,
			&post.Content,
			&post.UserID,
		)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)

	}

	return posts, nil

}

//SelectCreatedPosts ..
func (r *PostRepository) SelectCreatedPosts(userID int) ([]*model.Post, error) {
	post := &model.Post{}
	posts := []*model.Post{}

	rows, err := r.store.db.Query(
		`SELECT post_id, post_title, post_creation_date, post_content, user_id,  
		FROM post 
		WHERE user_id=?`,
		userID,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.CreationDate,
			&post.Content,
			&post.UserID,
		)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)

	}

	return posts, nil
}

//SelectLikedPosts ..
// func (r *PostRepository) SelectLikedPosts(id int) ([]*model.Post, error) {
// }
