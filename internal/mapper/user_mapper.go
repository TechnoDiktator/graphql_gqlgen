package mapper

import (
	"strconv"

	"github.com/tarangrastogi/graphql_gqlgen/internal/manualmodel"
	entity "github.com/tarangrastogi/graphql_gqlgen/internal/db_models"
)

func ToGraphQLUser(u *entity.User) *manualmodels.User {
	if u == nil {
		return  nil 
	}
	return &manualmodels.User{
		ID: strconv.FormatInt(u.ID , 10),
		Name : u.Name,
		Age : int32(u.Age),
	}

}


func ToEntityUser(input manualmodels.RegisterInput) *entity.User {
	return &entity.User{
		Name:  input.Name,
		Email: input.Email,
		Age:   input.Age,
	}
}
func ToGraphQLUsers(users []*entity.User) []*manualmodels.User {

	result := make([]*manualmodels.User, 0, len(users))

	for _, user := range users {
		result = append(result, ToGraphQLUser(user))
	}

	return result
}