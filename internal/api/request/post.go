package request

type PostCreateRequest struct {
	ImageURL string `json:"image_url" validate:"required"`
	Caption  string `json:"caption" validate:"required"`
}
