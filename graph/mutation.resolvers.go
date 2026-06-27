package graph

import (
	"context"
	"fmt"
	"log"

	"github.com/tarangrastogi/graphql_gqlgen/internal/auth"
	manualmodels "github.com/tarangrastogi/graphql_gqlgen/internal/manualmodel"
	"github.com/tarangrastogi/graphql_gqlgen/internal/mapper"
)

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(
	ctx context.Context,
	input manualmodels.CreatePostInput,
) (*manualmodels.Post, error) {

	log.Printf(
		"[Resolver] Mutation.CreatePost | title=%q",
		input.Title,
	)

	claims := auth.ForContext(ctx)
	if claims == nil {
		return nil, fmt.Errorf("unauthenticated")
	}

	entityPost := mapper.ToEntityPost(
		input,
		claims.UserID,
	)

	post, err := r.PostService.Create(ctx, entityPost)
	if err != nil {
		return nil, err
	}

	gqlPost := mapper.ToGraphQLPost(post)

	go func() {
		log.Println("Publishing post created")
		r.PostCreatedChan <- gqlPost
	}()

	return gqlPost, nil
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(
	ctx context.Context,
	input manualmodels.CreateCommentInput,
) (*manualmodels.Comment, error) {

	log.Printf(
		"[Resolver] Mutation.CreateComment | postID=%s",
		input.PostID,
	)

	claims := auth.ForContext(ctx)
	if claims == nil {
		return nil, fmt.Errorf("unauthenticated")
	}

	entityComment, err := mapper.ToEntityComment(
		input,
		claims.UserID,
	)
	if err != nil {
		return nil, err
	}

	comment, err := r.CommentService.Create(
		ctx,
		entityComment,
	)
	if err != nil {
		return nil, err
	}

	gqlComment := mapper.ToGraphQLComment(comment)

	go func() {
		log.Println("Publishing comment created")
		r.CommentCreatedChan <- gqlComment
	}()

	return gqlComment, nil
}

// Register is the resolver for the register field.
func (r *mutationResolver) Register(
	ctx context.Context,
	input manualmodels.RegisterInput,
) (*manualmodels.AuthPayload, error) {

	user, err := r.UserService.Register(ctx, input)
	if err != nil {
		return nil, err
	}

	token, err := auth.GenerateToken(
		user.ID,
		user.Email,
	)
	if err != nil {
		return nil, err
	}

	go func() {
		log.Println("Publishing user")
		r.UserCreatedChan <- mapper.ToGraphQLUser(user)
	}()

	return &manualmodels.AuthPayload{
		Token: token,
		User:  mapper.ToGraphQLUser(user),
	}, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(
	ctx context.Context,
	input manualmodels.LoginInput,
) (*manualmodels.AuthPayload, error) {

	user, err := r.UserService.Login(ctx, input)
	if err != nil {
		return nil, err
	}

	token, err := auth.GenerateToken(
		user.ID,
		user.Email,
	)
	if err != nil {
		return nil, err
	}

	return &manualmodels.AuthPayload{
		Token: token,
		User:  mapper.ToGraphQLUser(user),
	}, nil
}
