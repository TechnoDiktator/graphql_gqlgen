package graph

import (
	"context"
	"log"

	manualmodels "github.com/tarangrastogi/graphql_gqlgen/internal/manualmodel"
)

// UserCreated is the resolver for the userCreated field.
func (r *subscriptionResolver) UserCreated(
	ctx context.Context,
) (<-chan *manualmodels.User, error) {

	log.Println("[Resolver] Subscription.UserCreated")

	return r.UserCreatedChan, nil
}

// PostCreated is the resolver for the postCreated field.
func (r *subscriptionResolver) PostCreated(
	ctx context.Context,
) (<-chan *manualmodels.Post, error) {

	log.Println("[Resolver] Subscription.PostCreated")

	return r.PostCreatedChan, nil
}

// CommentCreated is the resolver for the commentCreated field.
func (r *subscriptionResolver) CommentCreated(
	ctx context.Context,
) (<-chan *manualmodels.Comment, error) {

	log.Println("[Resolver] Subscription.CommentCreated")

	return r.CommentCreatedChan, nil
}
