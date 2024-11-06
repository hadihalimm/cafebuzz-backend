package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/hadihalimm/cafebuzz-backend/internal/api/request"
	"github.com/hadihalimm/cafebuzz-backend/internal/models"
	"github.com/hadihalimm/cafebuzz-backend/internal/repository"
)

type PostService interface {
	Create(request request.PostCreateRequest) (*models.Post, error)
	FindByID(id uint64) (*models.Post, error)
	FindAllByCreator(creatorUUID uuid.UUID) ([]*models.Post, error)
}

type postService struct {
	repo     repository.PostRepository
	validate *validator.Validate
}

func NewPostService(repo repository.PostRepository, validate *validator.Validate) PostService {
	return &postService{repo: repo, validate: validate}
}

func (s *postService) Create(request request.PostCreateRequest) (*models.Post, error) {
	var postReq models.Post

	validateError := s.validate.Struct(request)
	if validateError != nil {
		return nil, validateError
	}

	postReq.ImageURL = request.ImageURL
	postReq.Caption = request.Caption
	newPost, createError := s.repo.Create(&postReq)
	if createError != nil {
		return nil, createError
	}
	return newPost, nil
}

func (s *postService) FindByID(id uint64) (*models.Post, error) {
	postFound, findError := s.repo.FindByID(id)
	if findError != nil {
		return nil, findError
	}
	return postFound, nil
}

func (s *postService) FindAllByCreator(creatorUUID uuid.UUID) ([]*models.Post, error) {
	postFound, findError := s.repo.FindAllByCreator(creatorUUID)
	if findError != nil {
		return nil, findError
	}
	return postFound, nil
}
