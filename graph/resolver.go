package graph

import "github.com/tarangrastogi/graphql_gqlgen/graph/model"

//"github.com/tarangrastogi/graphql_gqlgen/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct {
	UserCreatedChan    chan *model.User
	PostCreatedChan    chan *model.Post
	CommentCreatedChan	 chan *model.Comment
}
