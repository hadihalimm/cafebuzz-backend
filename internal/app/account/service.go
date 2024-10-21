package account

import (
	"github.com/go-playground/validator/v10"
	"github.com/hadihalimm/cafebuzz-backend/internal/api/request"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(request request.RegisterRequest) (Account, error)
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
	newAccount := Account{}
	validateError := s.validate.Struct(request)
	if validateError != nil {
		return newAccount, validateError
	}
	newAccount.Username = request.Username
	newAccount.FirstName = request.FirstName
	newAccount.LastName = request.LastName
	newAccount.Email = request.Email
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	newAccount.PasswordHash = string(hashedPassword)

	repoError := s.repo.Create(newAccount)
	if repoError != nil {
		return newAccount, repoError
	}
	return newAccount, nil
}
