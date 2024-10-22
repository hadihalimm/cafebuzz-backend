package account

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/hadihalimm/cafebuzz-backend/internal/api/request"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(request request.RegisterRequest) (Account, error)
	// Login(request request.LoginRequest) error
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

func (s *service) Register(request request.RegisterRequest) (Account, error) {
	req := &Account{}
	validateError := s.validate.Struct(request)
	if validateError != nil {
		return *req, validateError
	}

	_, findError := s.repo.FindByUsername(request.Username)
	if findError == nil {
		return *req, errors.New("username already exists")
	}

	req.Username = request.Username
	req.FirstName = request.FirstName
	req.LastName = request.LastName
	req.Email = request.Email
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	req.PasswordHash = string(hashedPassword)

	newAccount, createError := s.repo.Create(*req)
	if createError != nil {
		return newAccount, createError
	}
	return newAccount, nil
}
