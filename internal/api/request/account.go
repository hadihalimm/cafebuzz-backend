package request

type AccountUpdateRequest struct {
	FirstName      string `json:"first_name" validate:"required,min=1,max=128"`
	LastName       string `json:"last_name" validate:"required,min=1,max=128"`
	ProfilePicture string `json:"profile_picture" validate:"url"`
	Bio            string `json:"bio" validate:"alphanumeric"`
}
