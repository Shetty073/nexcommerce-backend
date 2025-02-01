package schemas

// LoginSchema holds the structure for user login data
type LoginSchema struct {
	Email    string `json:"email" validate:"omitempty,email"`
	Username string `json:"username" validate:"omitempty,alphanum"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}
