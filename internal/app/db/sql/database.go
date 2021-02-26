package sql

import (
	"database/sql"

	"github.com/zh3008/forum.git/internal/app/db"
)

// Store ...
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
	// postRepository          *PostRepository
	// commentRepository       *CommentRepository
	// subforumRepository      *SubforumRepository
	// postRatingRepository    *PostRatingRepository
	// commentRatingRepository *CommentRatingRepository
	//oleRepository          *RoleRepository
	// sessionRepository *SessionRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User ...
func (s *Store) User() db.UserRepository {
	if s.userRepository != nil {
		return s.userRepository.store.User()
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository.store.User()
}

// // Post ...
// func (s *Store) Post() db.PostRepository {
// 	if s.postRepository != nil {
// 		return s.postRepository.store.Post()
// 	}

// 	s.postRepository = &PostRepository{
// 		store: s,
// 	}

// 	return s.postRepository.store.Post()
// }

// // Comment ...
// func (s *Store) Comment() db.CommentRepository {
// 	if s.commentRepository != nil {
// 		return s.commentRepository.store.Comment()
// 	}

// 	s.commentRepository = &CommentRepository{
// 		store: s,
// 	}

// 	return s.commentRepository.store.Comment()
// }

// // Subforum ...
// func (s *Store) Subforum() db.SubforumRepository {
// 	if s.subforumRepository != nil {
// 		return s.subforumRepository.store.Subforum()
// 	}

// 	s.subforumRepository = &SubforumRepository{
// 		store: s,
// 	}

// 	return s.subforumRepository.store.Subforum()
// }

// // PostRating ...
// func (s *Store) PostRating() db.PostRatingRepository {
// 	if s.postRatingRepository != nil {
// 		return s.postRatingRepository.store.PostRating()
// 	}

// 	s.postRatingRepository = &PostRatingRepository{
// 		store: s,
// 	}

// 	return s.postRatingRepository.store.PostRating()
// }

// // CommentRating ...
// func (s *Store) CommentRating() db.CommentRatingRepository {
// 	if s.commentRatingRepository != nil {
// 		return s.commentRatingRepository.store.CommentRating()
// 	}

// 	s.commentRatingRepository = &CommentRatingRepository{
// 		store: s,
// 	}

// 	return s.commentRatingRepository.store.CommentRating()
// }
