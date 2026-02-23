package validation

type Login struct {
	Email    string `json:"email" validate:"required,email,max=50" example:"fake@example.com"`
	Password string `json:"password" validate:"required,min=8,max=20" example:"password1"`
}
