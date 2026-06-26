package manualmodels

/*
I mentioned "don't generate the structs." More precisely, gqlgen does generate models by default, but for any GraphQL 
type that you explicitly map in gqlgen.yml, it reuses your Go type instead of generating one.

*/


type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`

	Author   *User      `json:"author"`
	Comments []*Comment `json:"comments"`

	// Internal field
	UserID string `json:"-"`
}


type CreatePostInput struct {
	UserID  string `json:"userId"`
	Title   string `json:"title"`
	Content string `json:"content"`
}