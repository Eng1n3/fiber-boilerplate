package validation

type Register struct {
	Username string `json:"username" validate:"required,min=3,max=50" example:"John Doe"`
	Email    string `json:"email" validate:"required,email,max=50" example:"fake@example.com"`
	Password string `json:"password" validate:"required,min=8,max=20" example:"password1"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email,max=50" example:"fake@example.com"`
	Password string `json:"password" validate:"required,min=8,max=20" example:"password1"`
}
