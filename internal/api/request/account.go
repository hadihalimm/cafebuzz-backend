package request

type AccountUpdateRequest struct {
	Name           string `json:"name" validate:"required,min=1,max=128"`
	ProfilePicture string `json:"profile_picture" validate:"url"`
	Bio            string `json:"bio" validate:"alphanumeric"`
}
