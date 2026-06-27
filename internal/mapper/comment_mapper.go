package mapper

import (
	"strconv"

	entity "github.com/tarangrastogi/graphql_gqlgen/internal/db_models"
	manualmodels "github.com/tarangrastogi/graphql_gqlgen/internal/manualmodel"
)

func ToGraphQLComment(comment *entity.Comment) *manualmodels.Comment {

	if comment == nil {
		return nil
	}

	return &manualmodels.Comment{
		ID:      strconv.FormatInt(comment.ID, 10),
		Content: comment.Content,
		PostID:  strconv.FormatInt(comment.PostID, 10),
		UserID:  strconv.FormatInt(comment.UserID, 10),
	}
}

func ToEntityComment(input manualmodels.CreateCommentInput) (*entity.Comment, error) {

	postID, err := strconv.ParseInt(input.PostID, 10, 64)
	if err != nil {
		return nil, err
	}

	return &entity.Comment{
		PostID:  postID,
		Content: input.Content,
	}, nil
}

func ToGraphQLComments(comments []*entity.Comment) []*manualmodels.Comment {
	if comments == nil {
		return []*manualmodels.Comment{}
	}

	result := make([]*manualmodels.Comment, 0, len(comments))

	for _, comment := range comments {
		result = append(result, ToGraphQLComment(comment))
	}

	return result
}
