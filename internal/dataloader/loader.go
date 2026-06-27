package loader

import (
	"fmt"

	"github.com/tarangrastogi/graphql_gqlgen/internal/service"
)

/*
Excellent. At this point your project resembles a real backend:

✅ Clean architecture (Resolver → Service → Repository)
✅ PostgreSQL
✅ Custom GraphQL models
✅ Mappers
✅ Nested resolvers
✅ CRUD
✅ Basic subscriptions
✅ Resolver files split by responsibility

This is a solid foundation.

Next Phase: Solve the N+1 Problem with DataLoaders ⭐⭐⭐⭐⭐

If I were interviewing someone for a backend GraphQL role, this would be one of the first things I'd ask after CRUD.

Why?

Take this query:

query {
  posts {
    id
    title
    author {
      id
      name
    }
  }
}

Assume there are 100 posts.

Today your execution looks like:

Query.Posts()
        │
        ▼
PostRepository.GetAll()
             │
             ▼
        1 SQL query

Then GraphQL starts resolving every author:

Post.Author(post1)
        │
        ▼
UserRepository.GetByID(1)

Post.Author(post2)
        │
        ▼
UserRepository.GetByID(2)

Post.Author(post3)
        │
        ▼
UserRepository.GetByID(3)

...

Total:

1 query

+

100 queries

=

101 SQL queries

This is the famous N+1 problem.

What a DataLoader does

Instead of

1

2

3

5

8

9

executing immediately,

the loader collects them for a few milliseconds.

Then it executes

SELECT *
FROM users
WHERE id IN (1,2,3,5,8,9);

One query instead of six.

Where does the loader live?

We'll create something like

internal/

    dataloader/

        loader.go

        user_loader.go

        post_loader.go

        comment_loader.go
Request lifecycle

Instead of

HTTP Request

↓

Resolver

↓

Service

we'll have

HTTP Request

↓

Create DataLoaders

↓

Put them in Context

↓

Resolver

↓

Loader

↓

Repository

Every request gets its own loaders.



*/

/*
PostResolver.Author()
        │
        ▼
UserLoader.Load(7)
        │
        ▼
(graph-gophers batches requests)
        │
        ▼
Batch function receives:
[7, 3, 5, 9]
        │
        ▼
UserService.GetByIDs(...)
        │
        ▼
UserRepository.GetByIDs(...)
        │
        ▼
SELECT * FROM users WHERE id = ANY($1)
*/

func ErrUserNotFound(id int64) error {
	return fmt.Errorf("user %d not found", id)
}

type Loaders struct {
	UserLoader    *UserLoader
	PostLoader    *PostLoader
	CommentLoader *CommentLoader
}

func NewLoaders(userService service.UserService, postService service.PostService, commentService service.CommentService) *Loaders {

	return &Loaders{
		UserLoader: NewUserLoader(userService),
		// PostLoader:    NewPostLoader(postService),       // later
		// CommentLoader: NewCommentLoader(commentService), // later
	}
}
