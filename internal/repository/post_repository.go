package repository

import (
	context "context"
	"github.com/tarangrastogi/graphql_gqlgen/internal/db_models"

	 "github.com/jackc/pgx/v5/pgxpool"
)
type PostRepository interface {
	Create(ctx context.Context, post *entity.Post) (*entity.Post, error)
	GetByID(ctx context.Context, id int64) (*entity.Post, error)
	GetAll(ctx context.Context) ([]*entity.Post, error)
	GetByUserID(ctx context.Context, userID int64) ([]*entity.Post, error)
}
// 	Update(ctx context.Context, post *entity.Post) error

// 	Delete(ctx context.Context, id int64) error

type postRepository struct {
	db *pgxpool.Pool
}


func newPostRepository(db *pgxpool.Pool) PostRepository {
	return &postRepository{
		db: db,
	}
}


func (r *postRepository) Create(
	ctx context.Context,
	post *entity.Post,
) (*entity.Post, error) {

	query := `
	INSERT INTO posts(user_id, title, content)
	VALUES ($1, $2, $3)
	RETURNING id
	`

	err := r.db.QueryRow(
		ctx,
		query,
		post.UserID,
		post.Title,
		post.Content,
	).Scan(&post.ID)

	if err != nil {
		return nil, err
	}

	return post, nil
}
func (r *postRepository) GetByID(ctx context.Context, id int64) (*entity.Post, error) {

	query := `
		SELECT id, user_id, title, content
		FROM posts
		WHERE id = $1
	`

	post := &entity.Post{}

	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&post.ID,
		&post.UserID,
		&post.Title,
		&post.Content,
	)

	if err != nil {
		return nil, err
	}

	return post, nil
}


func (r *postRepository) GetAll(ctx context.Context) ([]*entity.Post, error) {

	query := `
		SELECT id, user_id, title, content
		FROM posts
	`
	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}	
	defer rows.Close()

	posts := []*entity.Post{}

	for rows.Next() {
		post := &entity.Post{}
		err := rows.Scan(	
		&post.ID,
		&post.UserID,
		&post.Title,
		&post.Content,
		)
		if err != nil {
			return nil, err
		}		
		posts = append(posts, post)
	}
	return posts, nil
}


func (r *postRepository) GetByUserID(ctx context.Context, userID int64) ([]*entity.Post, error) {
	query := `
		SELECT id, user_id, title, content
		FROM posts
		WHERE user_id = $1
	`
	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}	
	defer rows.Close()

	posts := []*entity.Post{}

	for rows.Next() {
		post := &entity.Post{}
		err := rows.Scan(	
		&post.ID,
		&post.UserID,
		&post.Title,
		&post.Content,
		)
		if err != nil {
			return nil, err
		}		
		posts = append(posts, post)
	}
	return posts, nil

}








