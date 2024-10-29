package services

import (
	"errors"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/hadihalimm/cafebuzz-backend/internal/api/request"
	"github.com/hadihalimm/cafebuzz-backend/internal/models"
	"github.com/hadihalimm/cafebuzz-backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type CafeService interface {
	Register(request request.CafeRegisterRequest) (*models.Cafe, error)
	Login(request request.LoginRequest) (string, error)
	Details(uuid uuid.UUID) (*models.Cafe, error)
	Update(uuid uuid.UUID, request request.CafeUpdateRequest) (*models.Cafe, error)
}

type cafeService struct {
	repo     repository.CafeRepository
	validate *validator.Validate
}

func NewCafeService(repo repository.CafeRepository, validate *validator.Validate) CafeService {
	return &cafeService{
		repo:     repo,
		validate: validate,
	}
}

func (s *cafeService) Register(request request.CafeRegisterRequest) (*models.Cafe, error) {
	var cafeReq models.Cafe

	validateError := s.validate.Struct(request)
	if validateError != nil {
		return nil, validateError
	}

	cafeFound, _ := s.repo.FindByUsername(request.Username)
	if cafeFound != nil {
		return nil, errors.New("username already exists")
	}

	cafeReq.Username = request.Username
	cafeReq.Name = request.Name
	cafeReq.Email = request.Email
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	cafeReq.PasswordHash = string(hashedPassword)
	cafeReq.Address = request.Address

	newCafe, createError := s.repo.Create(&cafeReq)
	if createError != nil {
		return nil, createError
	}
	return newCafe, nil
}

func (s *cafeService) Login(request request.LoginRequest) (string, error) {
	var cafeFound *models.Cafe

	validateError := s.validate.Struct(request)
	if validateError != nil {
		return "", validateError
	}

	cafeFound, findError := s.repo.FindByUsername(request.Username)
	if findError != nil {
		return "", findError
	}

	mismatchError := bcrypt.CompareHashAndPassword([]byte(cafeFound.PasswordHash), []byte(request.Password))
	if mismatchError != nil {
		return "", mismatchError
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid":         cafeFound.UUID,
		"account_type": "cafe",
		"exp":          time.Now().Add(time.Hour * 24).Unix(),
	})

	token, TokenError := generateToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if TokenError != nil {
		return "", TokenError
	}

	return token, nil
}

func (s *cafeService) Details(uuid uuid.UUID) (*models.Cafe, error) {
	cafeFound, findError := s.repo.FindByUUID(uuid)
	if findError != nil {
		return nil, findError
	}
	return cafeFound, nil
}

func (s *cafeService) Update(uuid uuid.UUID, request request.CafeUpdateRequest) (*models.Cafe, error) {
	cafeFound, findError := s.repo.FindByUUID(uuid)
	if findError != nil {
		return nil, findError
	}

	cafeFound.Name = request.Name
	cafeFound.Description = request.Description
	cafeFound.Address = request.Address
	cafeFound.ProfilePicture = request.ProfilePicture

	updatedCafe, updateError := s.repo.Update(cafeFound)
	if updateError != nil {
		return nil, updateError
	}
	return updatedCafe, nil
}
