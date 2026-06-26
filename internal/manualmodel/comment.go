package manualmodels

/*
I mentioned "don't generate the structs." More precisely, gqlgen does generate models by default, but for any GraphQL type 
that you explicitly map in gqlgen.yml, it reuses your Go type instead of generating one.

*/



type Comment struct {
	ID      string `json:"id"`
	Content string `json:"content"`

	Post   *Post `json:"post"`
	Author *User `json:"author"`

	// Internal fields    
	PostID string `json:"-"`
	UserID string `json:"-"`
}


type CreateCommentInput struct {
	PostID  string `json:"postId"`
	UserID  string `json:"userId"`
	Content string `json:"content"`
}
