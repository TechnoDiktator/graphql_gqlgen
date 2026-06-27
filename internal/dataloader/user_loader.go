package loader

import (
	"context"
	"log"

	"github.com/graph-gophers/dataloader/v7"

	entity "github.com/tarangrastogi/graphql_gqlgen/internal/db_models"
	"github.com/tarangrastogi/graphql_gqlgen/internal/service"
)

type UserLoader struct {
	batchLoader *dataloader.Loader[int64, *entity.User]
}

func NewUserLoader(userService service.UserService) *UserLoader {
	/*
		One thing to think about

		Right now your Load() executes the thunk immediately:

		return l.batchloader.Load(ctx, id)()

		That's perfectly fine because every resolver just wants:

		(*User, error)

		However, it's useful to understand what's happening under the hood.

		This:

		thunk := l.batchloader.Load(ctx, id)

		does not query the database immediately.

		The library waits briefly (by default only a few milliseconds), collects other Load() calls made during that window, and
		then invokes your batchFn once with all the accumulated keys. When you later execute:

		user, err := thunk()

		the thunk blocks until the batched query has completed and returns the corresponding result.

	*/
	batchFn := func(ctx context.Context, keys []int64) []*dataloader.Result[*entity.User] {
		log.Printf("[UserLoader] ================ Loading users: %v", keys)

		users, err := userService.GetByIDs(ctx, keys)
		if err != nil {

			results := make([]*dataloader.Result[*entity.User], len(keys))

			for i := range results {
				results[i] = &dataloader.Result[*entity.User]{
					Error: err,
				}
			}

			return results
		}

		// Build lookup map
		userMap := make(map[int64]*entity.User)

		for _, user := range users {
			userMap[user.ID] = user
		}

		results := make([]*dataloader.Result[*entity.User], len(keys))

		for i, id := range keys {

			user, ok := userMap[id]

			if !ok {
				results[i] = &dataloader.Result[*entity.User]{
					Error: ErrUserNotFound(id),
				}
				continue
			}

			results[i] = &dataloader.Result[*entity.User]{
				Data: user,
			}
		}

		return results
	}

	return &UserLoader{
		batchLoader: dataloader.NewBatchedLoader(batchFn),
	}
}

func (l *UserLoader) Load(
	ctx context.Context,
	id int64,
) (*entity.User, error) {

	return l.batchLoader.Load(ctx, id)() //this is called a thunk
}

/*
Notice something interesting.

Load() doesn't immediately hit the database.

Instead,

l.loader.Load(ctx, id)

returns a Thunk:

func() (*entity.User, error)

The database isn't queried until you execute the thunk:

() // execute

This is how the library collects many Load() calls before executing a single batch query.

Why is there a map?

Suppose GraphQL asks for

7
2
15

The loader calls

GetByIDs([]int64{7,2,15})

Your SQL

SELECT * FROM users
WHERE id = ANY($1)

might return

2
15
7

or even

15
7
2

SQL does not guarantee the order unless you explicitly use ORDER BY.

But the DataLoader contract says:

Return results in the same order as the requested keys.

That's why we create

userMap[id] = user

and then iterate over

keys

instead of

users

This is one of the most important implementation details of a DataLoader.


*/
