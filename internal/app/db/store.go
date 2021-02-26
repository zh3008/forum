package db

// Store ...
type Store interface {
	User() UserRepository
	// Post() PostRepository
	// Comment() CommentRepository
	// Subforum() SubforumRepository
	// CommentRating() CommentRatingRepository
	// PostRating() PostRatingRepository
	// Session() SessionRepository
	// Role() RoleRepository
}
