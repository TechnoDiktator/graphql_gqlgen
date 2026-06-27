package main

import (
	"fmt"

	"github.com/tarangrastogi/graphql_gqlgen/internal/auth"
)

func main() {

	token, err := auth.GenerateToken(
		24,
		"tarang@example.com",
	)
	if err != nil 	{
		fmt.Println("Something is wrong with the auth logic")
	}
	fmt.Println(token)

	claims, err := auth.ParseToken(token)

	fmt.Println(claims.UserID)
	fmt.Println(claims.Email)

}
