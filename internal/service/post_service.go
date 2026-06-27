package service


import (
	"context"

	entity "github.com/tarangrastogi/graphql_gqlgen/internal/db_models"
	"github.com/tarangrastogi/graphql_gqlgen/internal/repository"
)


type PostService interface {
	Create(ctx context.Context, post *entity.Post) (*entity.Post, error)
	GetByID(ctx context.Context, id int64) (*entity.Post, error)
	GetAll(ctx context.Context) ([]*entity.Post, error)
	GetByUserID(ctx context.Context, userID int64) ([]*entity.Post, error)

	GetByIDs(ctx context.Context, ids []int64) ([]*entity.Post, error)
}



type postService struct {
	postRepo repository.PostRepository
	userRepo repository.UserRepository
}

func NewPostService(
	postRepo repository.PostRepository,
	userRepo repository.UserRepository,
) PostService {

	return &postService{
		postRepo: postRepo,
		userRepo: userRepo,
	}
}


// Create creates a new post.
func (s *postService) Create(
	ctx context.Context,
	post *entity.Post,
) (*entity.Post, error) {

	// Ensure the user exists
	_, err := s.userRepo.GetByID(ctx, post.UserID)
	if err != nil {
		return nil, err
	}

	// Future business rules:
	// - title length
	// - profanity filter
	// - post limits
	// - permissions

	return s.postRepo.Create(ctx, post)
}
// GetByID fetches a post by its ID.
func (s *postService) GetByID(
	ctx context.Context,
	id int64,
) (*entity.Post, error) {

	return s.postRepo.GetByID(ctx, id)
}

// GetAll fetches all posts.
func (s *postService) GetAll(
	ctx context.Context,
) ([]*entity.Post, error) {

	return s.postRepo.GetAll(ctx)
}

// GetByUserID fetches all posts created by a user.
func (s *postService) GetByUserID(
	ctx context.Context,
	userID int64,
) ([]*entity.Post, error) {

	return s.postRepo.GetByUserID(ctx, userID)
}

func (s *postService) GetByIDs(
	ctx context.Context,
	ids []int64,
) ([]*entity.Post, error) {

	return s.postRepo.GetByIDs(ctx, ids)
}

// // Update updates an existing post.
// func (s *PostService) Update(
// 	ctx context.Context,
// 	post *entity.Post,
// ) error {

// 	return s.postRepo.Update(ctx, post)
// }

// // Delete deletes a post by its ID.
// func (s *PostService) Delete(
// 	ctx context.Context,
// 	id int64,
// ) error {

// 	return s.postRepo.Delete(ctx, id)
// }




