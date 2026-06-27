package manualmodels


/*
I mentioned "don't generate the structs." More precisely, gqlgen does generate models by default, 
but for any GraphQL type that you explicitly map in gqlgen.yml, it reuses your Go type instead of generating one.

*/

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int32  `json:"age"`

	Posts    []*Post    `json:"posts"`
	Comments []*Comment `json:"comments"`
}



//Now we will do register user
// type CreateUserInput struct {
// 	Name string `json:"name"`
// 	Age  int32  `json:"age"`
// }