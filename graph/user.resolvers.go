package graph

import (
	"context"
	"log"
	"strconv"

	manualmodels "github.com/tarangrastogi/graphql_gqlgen/internal/manualmodel"
	"github.com/tarangrastogi/graphql_gqlgen/internal/mapper"
)

// Posts is the resolver for the posts field.
func (r *userResolver) Posts(ctx context.Context, obj *manualmodels.User) ([]*manualmodels.Post, error) {
	//panic(fmt.Errorf("not implemented: Posts - posts"))
	log.Printf("[Resolver] User.Posts | userID=%s", obj.ID)
	userID, err := strconv.ParseInt(obj.ID, 10, 64)
	if err != nil {
		return nil, err
	}

	posts, err := r.PostService.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return mapper.ToGraphQLPosts(posts), nil

}

// Comments is the resolver for the comments field.
func (r *userResolver) Comments(ctx context.Context, obj *manualmodels.User) ([]*manualmodels.Comment, error) {
	//panic(fmt.Errorf("not implemented: Comments - comments"))
	log.Printf("[Resolver] User.Comments | userID=%s", obj.ID)
	userID, err := strconv.ParseInt(obj.ID, 10, 64)
	if err != nil {
		return nil, err
	}

	comments, err := r.CommentService.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return mapper.ToGraphQLComments(comments), nil

}
