package request

type RegisterRequest struct {
	Username  string `json:"username" validate:"required,min=1,max=64"`
	FirstName string `json:"first_name" validate:"required,min=1,max=128"`
	LastName  string `json:"last_name" validate:"required,min=1,max=128"`
	Email     string `json:"email" validate:"required,min=1,max=128"`
	Password  string `json:"password" validate:"required,min=8,max=32"`
}
