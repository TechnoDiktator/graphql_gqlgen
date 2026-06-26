package mapper

import (
	"strconv"

	entity "github.com/tarangrastogi/graphql_gqlgen/internal/db_models"
	"github.com/tarangrastogi/graphql_gqlgen/internal/manualmodel"
)

func ToGraphQLPost(post *entity.Post) *manualmodels.Post {

	if post == nil {
		return nil
	}

	return &manualmodels.Post{
		ID:      strconv.FormatInt(post.ID, 10),
		Title:   post.Title,
		Content: post.Content,
	}
}

func ToEntityPost(input manualmodels.CreatePostInput) (*entity.Post ,  error) {

	userID, err := strconv.ParseInt(input.UserID, 10, 64)
	if err != nil {
		return nil, err
	}

	return &entity.Post{
		UserID:  userID,
		Title:   input.Title,
		Content: input.Content,
	} , nil
}

func ToGraphQLPosts(posts []*entity.Post) []*manualmodels.Post {

	result := make([]*manualmodels.Post, 0, len(posts))

	for _, post := range posts {
		result = append(result, ToGraphQLPost(post))
	}

	return result
}