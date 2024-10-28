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

type AccountService interface {
	Register(request request.AccountRegisterRequest) (*models.Account, error)
	Login(request request.LoginRequest) (string, error)
	Details(uuid uuid.UUID) (*models.Account, error)
	Update(uuid uuid.UUID, request request.AccountUpdateRequest) (*models.Account, error)
}

type accountService struct {
	repo     repository.AccountRepository
	validate *validator.Validate
}

func NewAccountService(repository repository.AccountRepository, validate *validator.Validate) AccountService {
	return &accountService{
		repo:     repository,
		validate: validate,
	}
}

func (s *accountService) Register(request request.AccountRegisterRequest) (*models.Account, error) {
	var accountReq models.Account

	validateError := s.validate.Struct(request)
	if validateError != nil {
		return nil, validateError
	}

	_, findError := s.repo.FindByUsername(request.Username)
	if findError == nil {
		return nil, errors.New("username already exists")
	}

	accountReq.Username = request.Username
	accountReq.Name = request.Name
	accountReq.Email = request.Email
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	accountReq.PasswordHash = string(hashedPassword)

	newAccount, createError := s.repo.Create(&accountReq)
	if createError != nil {
		return newAccount, createError
	}
	return newAccount, nil
}

func (s *accountService) Login(request request.LoginRequest) (string, error) {
	var accountFound *models.Account

	validateError := s.validate.Struct(request)
	if validateError != nil {
		return "", validateError
	}

	accountFound, findError := s.repo.FindByUsername(request.Username)
	if findError != nil {
		return "", findError
	}

	mismatchError := bcrypt.CompareHashAndPassword([]byte(accountFound.PasswordHash), []byte(request.Password))
	if mismatchError != nil {
		return "", mismatchError
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid":         accountFound.UUID,
		"account_type": "personal",
		"exp":          time.Now().Add(time.Hour * 24).Unix(),
	})

	token, TokenError := generateToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if TokenError != nil {
		return "", TokenError
	}

	return token, nil
}

func (s *accountService) Details(uuid uuid.UUID) (*models.Account, error) {
	accountFound, findError := s.repo.FindByUUID(uuid)
	if findError != nil {
		return nil, findError
	}
	return accountFound, nil
}

func (s *accountService) Update(uuid uuid.UUID, request request.AccountUpdateRequest) (*models.Account, error) {
	accountFound, findError := s.repo.FindByUUID(uuid)
	if findError != nil {
		return nil, findError
	}

	accountFound.Name = request.Name
	accountFound.ProfilePicture = request.ProfilePicture
	accountFound.Bio = request.Bio
	updatedAccount, updateError := s.repo.Update(accountFound)
	if updateError != nil {
		return nil, updateError
	}
	return updatedAccount, nil
}
