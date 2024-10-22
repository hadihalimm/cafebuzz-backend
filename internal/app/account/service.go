package account

import (
	"errors"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/hadihalimm/cafebuzz-backend/internal/api/request"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(request request.RegisterRequest) (*Account, error)
	Login(request request.LoginRequest) (string, error)
	Update(uuid uuid.UUID, request request.AccountUpdateRequest) (*Account, error)
}

type service struct {
	repo     Repository
	validate *validator.Validate
}

func NewService(repository Repository, validate *validator.Validate) Service {
	return &service{
		repo:     repository,
		validate: validate,
	}
}

func (s *service) Register(request request.RegisterRequest) (*Account, error) {
	var accountReq Account

	validateError := s.validate.Struct(request)
	if validateError != nil {
		return nil, validateError
	}

	_, findError := s.repo.FindByUsername(request.Username)
	if findError == nil {
		return nil, errors.New("username already exists")
	}

	accountReq.Username = request.Username
	accountReq.FirstName = request.FirstName
	accountReq.LastName = request.LastName
	accountReq.Email = request.Email
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	accountReq.PasswordHash = string(hashedPassword)

	newAccount, createError := s.repo.Create(&accountReq)
	if createError != nil {
		return newAccount, createError
	}
	return newAccount, nil
}

func (s *service) Login(request request.LoginRequest) (string, error) {
	var accountFound Account

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
		"uuid": accountFound.UUID,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	token, TokenError := generateToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if TokenError != nil {
		return "", TokenError
	}

	return token, nil
}

func (s *service) Update(uuid uuid.UUID, request request.AccountUpdateRequest) (*Account, error) {
	accountFound, findError := s.repo.FindByUUID(uuid)
	if findError != nil {
		return nil, findError
	}

	accountFound.FirstName = request.FirstName
	accountFound.LastName = request.LastName
	accountFound.ProfilePicture = request.ProfilePicture
	accountFound.Bio = request.Bio
	updatedAccount, updateError := s.repo.Update(accountFound)
	if updateError != nil {
		return nil, updateError
	}
	return updatedAccount, nil
}
