package service

import (
	"context"

	entity "github.com/tarangrastogi/graphql_gqlgen/internal/db_models"
	"github.com/tarangrastogi/graphql_gqlgen/internal/repository"
)

type CommentService interface {
	Create(ctx context.Context, comment *entity.Comment) (*entity.Comment, error)
	GetByID(ctx context.Context, id int64) (*entity.Comment, error)
	GetAll(ctx context.Context) ([]*entity.Comment, error)
	GetByUserID(ctx context.Context, userID int64) ([]*entity.Comment, error)
	GetByPostID(ctx context.Context, postID int64) ([]*entity.Comment, error)

	GetByIDs(ctx context.Context, ids []int64) ([]*entity.Comment, error)
}

type commentService struct {
	commentRepo repository.CommentRepository
	userRepo    repository.UserRepository
	postRepo    repository.PostRepository
}

func NewCommentService(
	commentRepo repository.CommentRepository,
	userRepo repository.UserRepository,
	postRepo repository.PostRepository,
) CommentService {
	return &commentService{
		commentRepo: commentRepo,
		userRepo:    userRepo,
		postRepo:    postRepo,
	}
}

func (s *commentService) Create(
	ctx context.Context,
	comment *entity.Comment,
) (*entity.Comment, error) {

	// Verify user exists
	if _, err := s.userRepo.GetByID(ctx, comment.UserID); err != nil {
		return nil, err
	}

	// Verify post exists
	if _, err := s.postRepo.GetByID(ctx, comment.PostID); err != nil {
		return nil, err
	}

	return s.commentRepo.Create(ctx, comment)
}

// GetByID fetches a comment by its ID.
func (s *commentService) GetByID(
	ctx context.Context,
	id int64,
) (*entity.Comment, error) {

	return s.commentRepo.GetByID(ctx, id)
}

// GetAll fetches all comments.
func (s *commentService) GetAll(
	ctx context.Context,
) ([]*entity.Comment, error) {

	return s.commentRepo.GetAll(ctx)
}

// GetByUserID fetches all comments made by a user.
func (s *commentService) GetByUserID(
	ctx context.Context,
	userID int64,
) ([]*entity.Comment, error) {

	return s.commentRepo.GetByUserID(ctx, userID)
}

// GetByPostID fetches all comments on a post.
func (s *commentService) GetByPostID(
	ctx context.Context,
	postID int64,
) ([]*entity.Comment, error) {

	return s.commentRepo.GetByPostID(ctx, postID)
}
func (s *commentService) GetByIDs(
	ctx context.Context,
	ids []int64,
) ([]*entity.Comment, error) {

	return s.commentRepo.GetByIDs(ctx, ids)
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
