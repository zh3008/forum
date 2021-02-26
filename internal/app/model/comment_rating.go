package model

//CommentRating ..
type CommentRating struct {
	ID         int    `json:"comment_rating_id"`
	RatingType string `json:"comment_rating_type"`
	UserID     int    `json:"user_id"`
	CommentID  int    `json:"comment_id"`
}

//NewCommentRating ..
func NewCommentRating() *CommentRating {
	return &CommentRating{}
}
