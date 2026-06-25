package mapper

import (
	"strconv"

	entity "github.com/tarangrastogi/graphql_gqlgen/internal/db_models"
	"github.com/tarangrastogi/graphql_gqlgen/graph/model"
)

func ToGraphQLComment(comment *entity.Comment) *model.Comment {

	if comment == nil {
		return nil
	}

	return &model.Comment{
		ID:      strconv.FormatInt(comment.ID, 10),
		Content: comment.Content,
	}
}

func ToEntityComment(input model.CreateCommentInput) (*entity.Comment , error) {

	userID, err := strconv.ParseInt(input.UserID, 10, 64)
	
	if err != nil {
		return nil , err
	}

	postID, err := strconv.ParseInt(input.PostID, 10, 64)
	if err != nil {
			return nil , err
	}

	return &entity.Comment{
		UserID:  userID,
		PostID:  postID,
		Content: input.Content,
	} , nil
}

func ToGraphQLComments(comments []*entity.Comment) []*model.Comment {

	result := make([]*model.Comment, 0, len(comments))

	for _, comment := range comments {
		result = append(result, ToGraphQLComment(comment))
	}

	return result
}