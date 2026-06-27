package manualmodels

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}


type AuthPayload struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}