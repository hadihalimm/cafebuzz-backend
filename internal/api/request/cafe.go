package request

type CafeUpdateRequest struct {
	Name           string `json:"name" validate:"required,min=1,max=128"`
	Description    string `json:"description" validate:"alphanum,omitempty"`
	Address        string `json:"address" validate:"required"`
	ProfilePicture string `json:"profile_picture" validate:"url,omitempty"`
}
