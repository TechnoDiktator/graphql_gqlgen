package graph

import (
	"context"
	"log"
	"strconv"

	manualmodels "github.com/tarangrastogi/graphql_gqlgen/internal/manualmodel"
	"github.com/tarangrastogi/graphql_gqlgen/internal/mapper"
)

// Post is the resolver for the post field.
func (r *commentResolver) Post(ctx context.Context, obj *manualmodels.Comment) (*manualmodels.Post, error) {
	//panic(fmt.Errorf("not implemented: Post - post"))
	log.Printf(
		"[Resolver] Comment.Post | commentID=%s postID=%s",
		obj.ID,
		obj.PostID,
	)
	postID, err := strconv.ParseInt(obj.PostID, 10, 64)
	if err != nil {
		return nil, err
	}

	post, err := r.PostService.GetByID(ctx, postID)
	if err != nil {
		return nil, err
	}

	return mapper.ToGraphQLPost(post), nil
}

// Author is the resolver for the author field.
func (r *commentResolver) Author(ctx context.Context, obj *manualmodels.Comment) (*manualmodels.User, error) {
	//panic(fmt.Errorf("not implemented: Author - author"))
	log.Printf(
		"[Resolver] Comment.Author | commentID=%s userID=%s",
		obj.ID,
		obj.UserID,
	)
	userID, err := strconv.ParseInt(obj.UserID, 10, 64)
	if err != nil {
		return nil, err
	}

	user, err := r.UserService.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return mapper.ToGraphQLUser(user), nil

}
