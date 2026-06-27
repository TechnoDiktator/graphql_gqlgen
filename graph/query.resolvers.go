package graph

import (
	"context"
	"strconv"

	manualmodels "github.com/tarangrastogi/graphql_gqlgen/internal/manualmodel"
	"github.com/tarangrastogi/graphql_gqlgen/internal/mapper"
)

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*manualmodels.User, error) {
	//panic(fmt.Errorf("not implemented: Users - users"))

	users, err := r.UserService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.ToGraphQLUsers(users), nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*manualmodels.User, error) {
	//panic(fmt.Errorf("not implemented: User - user"))
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	user, err := r.UserService.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return mapper.ToGraphQLUser(user), nil
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*manualmodels.Post, error) {
	//panic(fmt.Errorf("not implemented: Posts - posts"))

	posts, err := r.PostService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.ToGraphQLPosts(posts), nil
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*manualmodels.Post, error) {
	//panic(fmt.Errorf("not implemented: Post - post"))
	postID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	post, err := r.PostService.GetByID(ctx, postID)
	if err != nil {
		return nil, err
	}

	return mapper.ToGraphQLPost(post), nil
}
