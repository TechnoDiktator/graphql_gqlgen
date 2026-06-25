package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/tarangrastogi/graphql_gqlgen/graph"
	"github.com/tarangrastogi/graphql_gqlgen/graph/model"
	"github.com/tarangrastogi/graphql_gqlgen/internal/db"
	"github.com/tarangrastogi/graphql_gqlgen/internal/repository"
	"github.com/tarangrastogi/graphql_gqlgen/internal/service"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	pool , err :=  db.NewPostgresConnection()

	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(pool)
	postRepo := repository.NewPostRepository(pool)
	commentRepo := repository.NewCommentRepository(pool)


	userService := service.NewUserService(userRepo)
	postService := service.NewPostService(postRepo ,  userRepo)
	commentService := service.NewCommentService(commentRepo , userRepo , postRepo)


	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		UserCreatedChan:    make(chan *model.User),
		PostCreatedChan:    make(chan *model.Post),
		CommentCreatedChan: make(chan *model.Comment),
		UserService: userService,
		PostService: postService,
		CommentService: commentService,
	}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})

	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
