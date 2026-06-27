package graph

import (
	manualmodels "github.com/tarangrastogi/graphql_gqlgen/internal/manualmodel"
	"github.com/tarangrastogi/graphql_gqlgen/internal/service"
)

//"github.com/tarangrastogi/graphql_gqlgen/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.
type Resolver struct {
	UserService    service.UserService
	PostService    service.PostService
	CommentService service.CommentService

	UserCreatedChan    chan *manualmodels.User
	PostCreatedChan    chan *manualmodels.Post
	CommentCreatedChan chan *manualmodels.Comment
}
