package model

//PostRating ..
type PostRating struct {
	ID         int    `json:"post_rating_id"`
	RatingType string `json:"post_rating_type"`
	UserID     int    `json:"user_id"`
	PostID     int    `json:"post_id"`
}

//NewPostRating ..
func NewPostRating() *PostRating {
	return &PostRating{}
}
