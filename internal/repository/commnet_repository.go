package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	entity "github.com/tarangrastogi/graphql_gqlgen/internal/db_models"
)

type CommentRepository interface {
	Create(ctx context.Context, comment *entity.Comment) (*entity.Comment, error)

	GetByID(ctx context.Context, id int64) (*entity.Comment, error)

	GetAll(ctx context.Context) ([]*entity.Comment, error)

	GetByUserID(ctx context.Context, userID int64) ([]*entity.Comment, error)

	GetByPostID(ctx context.Context, postID int64) ([]*entity.Comment, error)
	GetByIDs(ctx context.Context, ids []int64) ([]*entity.Comment, error)
}

type commentRepository struct {
	db *pgxpool.Pool
}

func NewCommentRepository(db *pgxpool.Pool) CommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (r *commentRepository) Create(ctx context.Context, comment *entity.Comment) (*entity.Comment, error) {

	query := `INSERT INTO comments(user_id, post_id, content)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := r.db.QueryRow(
		ctx,
		query,
		comment.UserID,
		comment.PostID,
		comment.Content,
	).Scan(&comment.ID)

	if err != nil {
		return nil, err
	}

	return comment, nil

}

func (r *commentRepository) GetByID(
	ctx context.Context,
	id int64,
) (*entity.Comment, error) {

	query := `
		SELECT id, user_id, post_id, content
		FROM comments
		WHERE id = $1
	`

	comment := &entity.Comment{}

	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&comment.ID,
		&comment.UserID,
		&comment.PostID,
		&comment.Content,
	)

	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (r *commentRepository) GetAll(
	ctx context.Context,
) ([]*entity.Comment, error) {

	query := `
		SELECT id, user_id, post_id, content
		FROM comments
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*entity.Comment

	for rows.Next() {

		comment := &entity.Comment{}

		err := rows.Scan(
			&comment.ID,
			&comment.UserID,
			&comment.PostID,
			&comment.Content,
		)

		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *commentRepository) GetByUserID(
	ctx context.Context,
	userID int64,
) ([]*entity.Comment, error) {

	query := `
		SELECT id, user_id, post_id, content
		FROM comments
		WHERE user_id = $1
	`

	rows, err := r.db.Query(
		ctx,
		query,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*entity.Comment

	for rows.Next() {

		comment := &entity.Comment{}

		err := rows.Scan(
			&comment.ID,
			&comment.UserID,
			&comment.PostID,
			&comment.Content,
		)

		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *commentRepository) GetByPostID(
	ctx context.Context,
	postID int64,
) ([]*entity.Comment, error) {

	query := `
		SELECT id, user_id, post_id, content
		FROM comments
		WHERE post_id = $1
	`

	rows, err := r.db.Query(
		ctx,
		query,
		postID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*entity.Comment

	for rows.Next() {

		comment := &entity.Comment{}

		err := rows.Scan(
			&comment.ID,
			&comment.UserID,
			&comment.PostID,
			&comment.Content,
		)

		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *commentRepository) GetByIDs(ctx context.Context, ids []int64) ([]*entity.Comment, error) {
	return nil, fmt.Errorf("not implemented getBYIDs Yet")
}
