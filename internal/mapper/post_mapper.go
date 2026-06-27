package mapper

import (
	"strconv"

	entity "github.com/tarangrastogi/graphql_gqlgen/internal/db_models"
	manualmodels "github.com/tarangrastogi/graphql_gqlgen/internal/manualmodel"
)

func ToGraphQLPost(post *entity.Post) *manualmodels.Post {
	return &manualmodels.Post{
		ID:      strconv.FormatInt(post.ID, 10),
		Title:   post.Title,
		Content: post.Content,

		UserID: strconv.FormatInt(post.UserID, 10), // <-- REQUIRED
	}
}

func ToEntityPost(input manualmodels.CreatePostInput) *entity.Post {
	return &entity.Post{
		Title:   input.Title,
		Content: input.Content,
	}
}


func ToGraphQLPosts(posts []*entity.Post) []*manualmodels.Post {

	result := make([]*manualmodels.Post, 0, len(posts))

	for _, post := range posts {
		result = append(result, ToGraphQLPost(post))
	}

	return result
}
