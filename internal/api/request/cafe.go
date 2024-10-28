package request

type CafeUpdateRequest struct {
	Name           string `json:"name" validate:"required,min=1,max=128"`
	Description    string `json:"description" validate:"alphanumeric"`
	Address        string `json:"address" validate:"alphanumeric,min=8"`
	ProfilePicture string `json:"profile_picture" validate:"url"`
}
