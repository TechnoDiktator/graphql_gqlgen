package graph

import (
	"context"
	"log"
	"strconv"

	loader "github.com/tarangrastogi/graphql_gqlgen/internal/dataloader"
	manualmodels "github.com/tarangrastogi/graphql_gqlgen/internal/manualmodel"
	"github.com/tarangrastogi/graphql_gqlgen/internal/mapper"
)

/*IMPLEMENTING BATCH LOADER IN posts{
	jdswcl;kjadclkhjsdclk
	author{
		swjdckhdlcik
	}

}*/

//===========================OLD VERSION
// Author is the resolver for the author field.
// func (r *postResolver) Author(ctx context.Context, obj *manualmodels.Post) (*manualmodels.User, error) {
// 	//panic(fmt.Errorf("not implemented: Author - author"))
// 	log.Printf(
// 		"[Resolver] Post.Author | postID=%s userID=%s",
// 		obj.ID,
// 		obj.UserID,
// 	)
// 	userID, err := strconv.ParseInt(obj.UserID, 10, 64)
// 	if err != nil {
// 		return nil, err
// 	}

// 	user, err := r.UserService.GetByID(ctx, userID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return mapper.ToGraphQLUser(user), nil
// }

// ==================================NEW VERSIOn with DATA LOADER for BATCHING
// Author is the resolver for the author field.
func (r *postResolver) Author(
	ctx context.Context,
	obj *manualmodels.Post,
) (*manualmodels.User, error) {

	log.Printf(
		"[Resolver] Post.Author | postID=%s userID=%s",
		obj.ID,
		obj.UserID,
	)

	userID, err := strconv.ParseInt(obj.UserID, 10, 64)
	if err != nil {
		return nil, err
	}

	// Get the request-scoped loaders from the context
	loaders := loader.For(ctx)

	// This does NOT hit the database immediately.
	// It queues the request so multiple resolver calls can be batched together.
	user, err := loaders.UserLoader.Load(ctx, userID)
	if err != nil {
		return nil, err
	}

	return mapper.ToGraphQLUser(user), nil


	/*
			Query.Posts()
				│
				▼
		returns 100 posts
				│
				▼
		Post.Author(post1)
				│
				├── UserLoader.Load(ctx, 5)
				│
		Post.Author(post2)
				│
				├── UserLoader.Load(ctx, 2)
				│
		Post.Author(post3)
				│
				├── UserLoader.Load(ctx, 5)
				│
		Post.Author(post4)
				│
				├── UserLoader.Load(ctx, 8)
				│
				▼
		UserLoader batches the keys
				▼
		batchFn(ctx, []int64{5,2,8})
				▼
		UserService.GetByIDs(...)
				▼
		Repository.GetByIDs(...)
				▼
		SELECT * FROM users WHERE id = ANY($1)
	*/

}
/*
Where you are now

At this point, you've built something that closely matches how DataLoader is used in production GraphQL services. You have:

✅ Layered architecture (Resolvers → Services → Repositories)
✅ Request-scoped DataLoaders
✅ Batch loading
✅ Request-level caching
✅ Elimination of the N+1 problem for Post.Author

This is a significant milestone. From here, you can apply the same pattern to Comment.Author and Comment.Post, and eventually 
introduce middleware, authentication, pagination, tracing, and other production concerns.

*/




// Comments is the resolver for the comments field.
func (r *postResolver) Comments(ctx context.Context, obj *manualmodels.Post) ([]*manualmodels.Comment, error) {
	//panic(fmt.Errorf("not implemented: Comments - comments"))
	log.Printf("[Resolver] Post.Comments | postID=%s", obj.ID)
	postID, err := strconv.ParseInt(obj.ID, 10, 64)
	if err != nil {
		return nil, err
	}

	comments, err := r.CommentService.GetByPostID(ctx, postID)
	if err != nil {
		return nil, err
	}

	return mapper.ToGraphQLComments(comments), nil

}


/*

{
  "data": {
    "createUser": {
      "id": "24",
      "name": "Tarang",
      "age": 28
    }
  }
}

{
  "data": {
    "createUser": {
      "id": "25",
      "name": "valice",
      "age": 28
    }
  }
}

{
  "data": {
    "createUser": {
      "id": "26",
      "name": "mogambo",
      "age": 56
    }
  }
}

*/