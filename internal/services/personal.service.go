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

type PersonalAccountService interface {
	Register(request request.AccountRegisterRequest) (*models.PersonalAccount, error)
	Login(request request.LoginRequest) (string, error)
	Details(uuid uuid.UUID) (*models.PersonalAccount, error)
	Update(uuid uuid.UUID, request request.AccountUpdateRequest) (*models.PersonalAccount, error)
}

type personalAccountService struct {
	repo     repository.PersonalAccountRepository
	validate *validator.Validate
}

func NewPersonalAccountService(repository repository.PersonalAccountRepository, validate *validator.Validate) PersonalAccountService {
	return &personalAccountService{
		repo:     repository,
		validate: validate,
	}
}

func (s *personalAccountService) Register(request request.AccountRegisterRequest) (*models.PersonalAccount, error) {
	var accountReq models.PersonalAccount

	validateError := s.validate.Struct(request)
	if validateError != nil {
		return nil, validateError
	}

	existedAccount, _ := s.repo.FindByUsername(request.Username)
	if existedAccount != nil {
		return nil, errors.New("username already exists")
	}

	accountReq.Account.Username = request.Username
	accountReq.Account.Name = request.Name
	accountReq.Account.Email = request.Email
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	accountReq.Account.PasswordHash = string(hashedPassword)

	newAccount, createError := s.repo.Create(&accountReq)
	if createError != nil {
		return newAccount, createError
	}
	return newAccount, nil
}

func (s *personalAccountService) Login(request request.LoginRequest) (string, error) {
	var accountFound *models.PersonalAccount

	validateError := s.validate.Struct(request)
	if validateError != nil {
		return "", validateError
	}

	accountFound, findError := s.repo.FindByUsername(request.Username)
	if findError != nil {
		return "", findError
	}

	mismatchError := bcrypt.CompareHashAndPassword([]byte(accountFound.Account.PasswordHash), []byte(request.Password))
	if mismatchError != nil {
		return "", mismatchError
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid":         accountFound.Account.UUID,
		"account_type": "personal",
		"exp":          time.Now().Add(time.Hour * 24).Unix(),
	})

	token, TokenError := generateToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if TokenError != nil {
		return "", TokenError
	}

	return token, nil
}

func (s *personalAccountService) Details(uuid uuid.UUID) (*models.PersonalAccount, error) {
	accountFound, findError := s.repo.FindByUUID(uuid)
	if findError != nil {
		return nil, findError
	}
	return accountFound, nil
}

func (s *personalAccountService) Update(uuid uuid.UUID, request request.AccountUpdateRequest) (*models.PersonalAccount, error) {
	validateError := s.validate.Struct(request)
	if validateError != nil {
		return nil, validateError
	}

	accountFound, findError := s.repo.FindByUUID(uuid)
	if findError != nil {
		return nil, findError
	}

	accountFound.Account.Name = request.Name
	accountFound.Account.ProfilePicture = request.ProfilePicture
	accountFound.Bio = request.Bio
	updatedAccount, updateError := s.repo.Update(accountFound)
	if updateError != nil {
		return nil, updateError
	}
	return updatedAccount, nil
}
