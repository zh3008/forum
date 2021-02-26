package model

//Comment ..
type Comment struct {
	ID           int    `json:"comment_id"`
	CreationDate string `json:"comment_creation_date"`
	Content      string `json:"comment_content"`
	UserID       int    `json:"user_id"`
	PostID       int    `json:"post_id"`
}

//NewComment ..
func NewComment() *Comment {
	return &Comment{}
}
