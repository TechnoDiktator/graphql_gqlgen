package graph

import (
	"context"
	"fmt"
	"log"

	manualmodels "github.com/tarangrastogi/graphql_gqlgen/internal/manualmodel"
	"github.com/tarangrastogi/graphql_gqlgen/internal/mapper"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input manualmodels.CreateUserInput) (*manualmodels.User, error) {
	//panic(fmt.Errorf("not implemented: CreateUser - createUser"))

	log.Printf(
		"[Resolver] Mutation.CreateUser | name=%s age=%d",
		input.Name,
		input.Age,
	)

	entityUser := mapper.ToEntityUser(input)

	user, err := r.UserService.Create(ctx, entityUser)
	if err != nil {
		return nil, err
	}

	graphqluser := mapper.ToGraphQLUser(user)

	go func() {
		log.Panicln("Publishing user")
		r.UserCreatedChan <- graphqluser
	}()
	return graphqluser, nil
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input manualmodels.CreatePostInput) (*manualmodels.Post, error) {
	//panic(fmt.Errorf("not implemented: CreatePost - createPost"))
	log.Printf(
		"[Resolver] Mutation.CreatePost | userID=%s title=%q",
		input.UserID,
		input.Title,
	)

	fmt.Printf("\n gql model object %s", input)
	entityPost, err := mapper.ToEntityPost(input)
	if err != nil {
		return nil, err
	}
	fmt.Printf("\n db entity created %s", entityPost)
	post, err := r.PostService.Create(ctx, entityPost)
	if err != nil {
		return nil, err
	}
	fmt.Printf("\n db object created %s", input)
	gqlPost := mapper.ToGraphQLPost(post)

	go func() {
		r.PostCreatedChan <- gqlPost
	}()

	return gqlPost, nil
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, input manualmodels.CreateCommentInput) (*manualmodels.Comment, error) {
	//panic(fmt.Errorf("not implemented: CreateComment - createComment"))
	log.Printf(
		"[Resolver] Mutation.CreateComment | postID=%s userID=%s",
		input.PostID,
		input.UserID,
	)
	entityComment, err := mapper.ToEntityComment(input)
	if err != nil {
		return nil, err
	}

	comment, err := r.CommentService.Create(ctx, entityComment)
	if err != nil {
		return nil, err
	}

	gqlComment := mapper.ToGraphQLComment(comment)

	go func() {
		r.CommentCreatedChan <- gqlComment
	}()

	return gqlComment, nil
}
