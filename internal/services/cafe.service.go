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

type CafeAccountService interface {
	Register(request request.CafeRegisterRequest) (*models.CafeAccount, error)
	Login(request request.LoginRequest) (string, error)
	Details(uuid uuid.UUID) (*models.CafeAccount, error)
	Update(uuid uuid.UUID, request request.CafeUpdateRequest) (*models.CafeAccount, error)
}

type cafeAccountService struct {
	repo     repository.CafeAccountRepository
	validate *validator.Validate
}

func NewCafeAccountService(repo repository.CafeAccountRepository, validate *validator.Validate) CafeAccountService {
	return &cafeAccountService{
		repo:     repo,
		validate: validate,
	}
}

func (s *cafeAccountService) Register(request request.CafeRegisterRequest) (*models.CafeAccount, error) {
	var cafeReq models.CafeAccount

	validateError := s.validate.Struct(request)
	if validateError != nil {
		return nil, validateError
	}

	cafeFound, _ := s.repo.FindByUsername(request.Username)
	if cafeFound != nil {
		return nil, errors.New("username already exists")
	}

	cafeReq.Account.Username = request.Username
	cafeReq.Account.Name = request.Name
	cafeReq.Account.Email = request.Email
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	cafeReq.Account.PasswordHash = string(hashedPassword)
	cafeReq.Address = request.Address

	newCafe, createError := s.repo.Create(&cafeReq)
	if createError != nil {
		return nil, createError
	}
	return newCafe, nil
}

func (s *cafeAccountService) Login(request request.LoginRequest) (string, error) {
	var cafeFound *models.CafeAccount

	validateError := s.validate.Struct(request)
	if validateError != nil {
		return "", validateError
	}

	cafeFound, findError := s.repo.FindByUsername(request.Username)
	if findError != nil {
		return "", findError
	}

	mismatchError := bcrypt.CompareHashAndPassword([]byte(cafeFound.Account.PasswordHash), []byte(request.Password))
	if mismatchError != nil {
		return "", mismatchError
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid":         cafeFound.Account.UUID,
		"account_type": "cafe",
		"exp":          time.Now().Add(time.Hour * 24).Unix(),
	})

	token, TokenError := generateToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if TokenError != nil {
		return "", TokenError
	}

	return token, nil
}

func (s *cafeAccountService) Details(uuid uuid.UUID) (*models.CafeAccount, error) {
	cafeFound, findError := s.repo.FindByUUID(uuid)
	if findError != nil {
		return nil, findError
	}
	return cafeFound, nil
}

func (s *cafeAccountService) Update(uuid uuid.UUID, request request.CafeUpdateRequest) (*models.CafeAccount, error) {
	validateError := s.validate.Struct(request)
	if validateError != nil {
		return nil, validateError
	}

	cafeFound, findError := s.repo.FindByUUID(uuid)
	if findError != nil {
		return nil, findError
	}

	cafeFound.Account.Name = request.Name
	cafeFound.Description = request.Description
	cafeFound.Address = request.Address
	cafeFound.Account.ProfilePicture = request.ProfilePicture

	updatedCafe, updateError := s.repo.Update(cafeFound)
	if updateError != nil {
		return nil, updateError
	}
	return updatedCafe, nil
}
