package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	entity "github.com/tarangrastogi/graphql_gqlgen/internal/db_models"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)

	GetByID(ctx context.Context, id int64) (*entity.User, error)

	GetAll(ctx context.Context) ([]*entity.User, error)

	GetByIDs(ctx context.Context, ids []int64) ([]*entity.User, error)

	// Update(ctx context.Context, user *entity.User) error

	// Delete(ctx context.Context, id int64) error
}

type postgresUserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &postgresUserRepository{
		db: db,
	}
}

func (r *postgresUserRepository) Create(
	ctx context.Context,
	user *entity.User,
) (*entity.User, error) {

	query := `
	INSERT INTO users(name, age)
	VALUES ($1, $2)
	RETURNING id
	`

	err := r.db.QueryRow(
		ctx,
		query,
		user.Name,
		user.Age,
	).Scan(&user.ID)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *postgresUserRepository) GetByID(
	ctx context.Context,
	id int64,
) (*entity.User, error) {

	query := `
	SELECT id, name, age
	FROM users
	WHERE id = $1
	`

	user := &entity.User{}

	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Age,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *postgresUserRepository) GetAll(
	ctx context.Context,
) ([]*entity.User, error) {

	query := `
	SELECT id, name, age
	FROM users
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User

	for rows.Next() {
		user := &entity.User{}

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Age,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *postgresUserRepository) GetByIDs(ctx context.Context, ids []int64) ([]*entity.User, error) {

	query := `
		SELECT id, name, age
		FROM users
		WHERE id = ANY($1)
	`

	rows, err := r.db.Query(ctx, query, ids)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*entity.User, 0, len(ids))

	for rows.Next() {
		user := &entity.User{}

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Age,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil

}
