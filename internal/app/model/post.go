package model

//Post ..
type Post struct {
	ID           int    `json:"post_id"`
	Title        string `json:"post_title"`
	CreationDate string `json:"post_creation_date"`
	Content      string `json:"post_content"`
	UserID       int    `json:"user_id"`
}

//NewPost ...
func NewPost() *Post {
	return &Post{}
}
