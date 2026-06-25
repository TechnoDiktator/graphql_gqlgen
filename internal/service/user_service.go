package service

import (
	"context"

	"github.com/tarangrastogi/graphql_gqlgen/internal/db_models"
	"github.com/tarangrastogi/graphql_gqlgen/internal/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Create(
	ctx context.Context,
	user *entity.User,
) (*entity.User, error) {

	return s.userRepo.Create(ctx, user)
}

func (s *UserService) GetByID(
	ctx context.Context,
	id int64,
) (*entity.User, error) {

	return s.userRepo.GetByID(ctx, id)
}

func (s *UserService) GetAll(
	ctx context.Context,
) ([]*entity.User, error) {

	return s.userRepo.GetAll(ctx)
}



