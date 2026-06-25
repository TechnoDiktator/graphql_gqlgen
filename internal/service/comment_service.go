package service

import (
	"context"

	entity "github.com/tarangrastogi/graphql_gqlgen/internal/db_models"
	"github.com/tarangrastogi/graphql_gqlgen/internal/repository"
)

type CommentService struct {
	commentRepo repository.CommentRepository
	userRepo    repository.UserRepository
	postRepo    repository.PostRepository
}

func NewCommentService(
	commentRepo repository.CommentRepository,
	userRepo repository.UserRepository,
	postRepo repository.PostRepository,
) *CommentService {
	return &CommentService{
		commentRepo: commentRepo,
		userRepo:    userRepo,
		postRepo:    postRepo,
	}
}

// Create creates a new comment.
func (s *CommentService) Create(
	ctx context.Context,
	comment *entity.Comment,
) (*entity.Comment, error) {

	// Business logic can be added here later:
	// - Verify the user exists.
	// - Verify the post exists.
	// - Check if commenting is allowed.

	return s.commentRepo.Create(ctx, comment)
}

// GetByID fetches a comment by its ID.
func (s *CommentService) GetByID(
	ctx context.Context,
	id int64,
) (*entity.Comment, error) {

	return s.commentRepo.GetByID(ctx, id)
}

// GetAll fetches all comments.
func (s *CommentService) GetAll(
	ctx context.Context,
) ([]*entity.Comment, error) {

	return s.commentRepo.GetAll(ctx)
}

// GetByUserID fetches all comments made by a user.
func (s *CommentService) GetByUserID(
	ctx context.Context,
	userID int64,
) ([]*entity.Comment, error) {

	return s.commentRepo.GetByUserID(ctx, userID)
}

// GetByPostID fetches all comments on a post.
func (s *CommentService) GetByPostID(
	ctx context.Context,
	postID int64,
) ([]*entity.Comment, error) {

	return s.commentRepo.GetByPostID(ctx, postID)
}

// Update updates an existing comment.
// func (s *CommentService) Update(
// 	ctx context.Context,
// 	comment *entity.Comment,
// ) error {

// 	return s.commentRepo.Update(ctx, comment)
// }

// // Delete deletes a comment by its ID.
// func (s *CommentService) Delete(
// 	ctx context.Context,
// 	id int64,
// ) error {

// 	return s.commentRepo.Delete(ctx, id)
// }
