package service


import (
	"context"

	entity "github.com/tarangrastogi/graphql_gqlgen/internal/db_models"
	"github.com/tarangrastogi/graphql_gqlgen/internal/repository"
)

type PostService struct {
	postRepo repository.PostRepository
	userRepo repository.UserRepository
}

func NewPostService(
	postRepo repository.PostRepository,
	userRepo repository.UserRepository,
) *PostService {
	return &PostService{
		postRepo: postRepo,
		userRepo: userRepo,
	}
}

// Create creates a new post.
func (s *PostService) Create(
	ctx context.Context,
	post *entity.Post,
) (*entity.Post, error) {

	// Business logic can be added here later:
	// - Validate the user exists.
	// - Check permissions.
	// - Sanitize content.

	return s.postRepo.Create(ctx, post)
}

// GetByID fetches a post by its ID.
func (s *PostService) GetByID(
	ctx context.Context,
	id int64,
) (*entity.Post, error) {

	return s.postRepo.GetByID(ctx, id)
}

// GetAll fetches all posts.
func (s *PostService) GetAll(
	ctx context.Context,
) ([]*entity.Post, error) {

	return s.postRepo.GetAll(ctx)
}

// GetByUserID fetches all posts created by a user.
func (s *PostService) GetByUserID(
	ctx context.Context,
	userID int64,
) ([]*entity.Post, error) {

	return s.postRepo.GetByUserID(ctx, userID)
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




