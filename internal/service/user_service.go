package service

import (
	"context"

	entity "github.com/tarangrastogi/graphql_gqlgen/internal/db_models"
	"github.com/tarangrastogi/graphql_gqlgen/internal/repository"
)

type UserService interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	GetByID(ctx context.Context, id int64) (*entity.User, error)
	GetAll(ctx context.Context) ([]*entity.User, error)

	GetByIDs(ctx context.Context, ids []int64) ([]*entity.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) Create(
	ctx context.Context,
	user *entity.User,
) (*entity.User, error) {

	return s.userRepo.Create(ctx, user)
}

func (s *userService) GetByID(
	ctx context.Context,
	id int64,
) (*entity.User, error) {

	return s.userRepo.GetByID(ctx, id)
}

func (s *userService) GetAll(
	ctx context.Context,
) ([]*entity.User, error) {

	return s.userRepo.GetAll(ctx)
}

func (s *userService) GetByIDs(
	ctx context.Context,
	ids []int64,
) ([]*entity.User, error) {

	return s.userRepo.GetByIDs(ctx, ids)
}
