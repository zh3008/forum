package db

import "github.com/zh3008/forum.git/internal/app/model"

//UserRepository ..
type UserRepository interface {
	CreateUser(*model.User) error
	FindByUserID(userID int) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}

// //PostRepository ..
// type PostRepository interface {
// 	CreatePost(*model.Post) error
// 	FindByPostID(postID int) (*model.Post, error)
// 	SelectAllPosts() ([]*model.Post, error)
// 	SelectCreatedPosts(userID int) ([]*model.Post, error)
// 	SelectSubforumPosts(subforumID int) ([]*model.Post, error)
// 	SelectLikedPosts(id int) ([]*model.Post, error)
// }

// //CommentRepository ..
// type CommentRepository interface {
// 	CreateComment(*model.Comment) error
// 	SelectCommentsByPost(postID int) ([]*model.Comment, error)
// }

// //SubforumRepository ..
// type SubforumRepository interface {
// 	CreateSubforum(*model.Subforum) error
// 	SelectSubForumByID(subforumID int) (*model.Subforum, error)
// 	SelectSubForumByName(name int) (*model.Subforum, error)
// }

// //PostRatingRepository ..
// type PostRatingRepository interface {
// 	CreatePostRating(*model.PostRating) error
// 	SelectPostRatingType(*model.PostRating) (*model.PostRating, error)
// 	UpdatePostRating(*model.PostRating) error
// 	DeletePostRating(*model.PostRating) error
// }

// //CommentRatingRepository ..
// type CommentRatingRepository interface {
// 	CreateCommentRating(*model.CommentRating) error
// 	SelectCommentRatingType(*model.CommentRating) (*model.CommentRating, error)
// 	UpdateCommentRating(*model.CommentRating) error
// 	DeleteCommentRating(*model.CommentRating) error
// }

// //SessionRepository ..
// type SessionRepository interface {
// 	CreateSession()
// 	CheckSession()
// 	UpdateSession()
// 	DeleteSession()
// }

// //RoleRepository ..
// type RoleRepository interface {
// }
