package request

type AccountRegisterRequest struct {
	Username string `json:"username" validate:"required,min=1,max=64"`
	Name     string `json:"name" validate:"required,min=1,max=128"`
	Email    string `json:"email" validate:"required,min=1,max=128"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type CafeRegisterRequest struct {
	Username string `json:"username" validate:"required,min=1,max=64"`
	Name     string `json:"name" validate:"required,min=1,max=128"`
	Email    string `json:"email" validate:"required,min=1,max=128"`
	Password string `json:"password" validate:"required,min=8,max=32"`
	Address  string `json:"address" validate:"required,min=8"`
}
