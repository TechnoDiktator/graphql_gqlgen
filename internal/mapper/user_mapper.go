package mapper

import (
	"strconv"

	"github.com/tarangrastogi/graphql_gqlgen/graph/model"
	entity "github.com/tarangrastogi/graphql_gqlgen/internal/db_models"
)

func ToGraphQLUser(u *entity.User) *model.User {
	if u == nil {
		return  nil 
	}
	return &model.User{
		ID: strconv.FormatInt(u.ID , 10),
		Name : u.Name,
		Age : int32(u.Age),
	}

}


func ToEntityUser(input model.CreateUserInput) *entity.User {
	return &entity.User{
		Name: input.Name,
		Age:  int(input.Age),
	}
}

func ToGraphQLUsers(users []*entity.User) []*model.User {

	result := make([]*model.User, 0, len(users))

	for _, user := range users {
		result = append(result, ToGraphQLUser(user))
	}

	return result
}