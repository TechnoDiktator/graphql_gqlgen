package service

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/tarangrastogi/graphql_gqlgen/internal/auth"
	entity "github.com/tarangrastogi/graphql_gqlgen/internal/db_models"
	manualmodels "github.com/tarangrastogi/graphql_gqlgen/internal/manualmodel"
	"github.com/tarangrastogi/graphql_gqlgen/internal/repository"
)

type UserService interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)

	GetByID(ctx context.Context, id int64) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)

	GetAll(ctx context.Context) ([]*entity.User, error)
	GetByIDs(ctx context.Context, ids []int64) ([]*entity.User, error)

	Register(ctx context.Context, input manualmodels.RegisterInput) (*entity.User, error)
	Login(ctx context.Context, input manualmodels.LoginInput) (*entity.User, error)
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

func (s *userService) GetByEmail(
	ctx context.Context,
	email string,
) (*entity.User, error) {

	return s.userRepo.GetByEmail(ctx, email)
}

func (s *userService) Register(
	ctx context.Context,
	input manualmodels.RegisterInput,
) (*entity.User, error) {

	// Check if email already exists

	user, err := s.userRepo.GetByEmail(ctx, input.Email)

	if err == nil {
		return nil, errors.New("email already registered")
	}

	if !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}

	hash, err := auth.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user = &entity.User{
		Name:         input.Name,
		Email:        input.Email,
		PasswordHash: hash,
		Age:          input.Age,
	}

	return s.userRepo.Create(ctx, user)
}

func (s *userService) Login(
	ctx context.Context,
	input manualmodels.LoginInput,
) (*entity.User, error) {

	user, err := s.userRepo.GetByEmail(ctx, input.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = auth.CheckPassword(
		input.Password,
		user.PasswordHash,
	)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}
