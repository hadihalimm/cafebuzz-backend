package repository

import (
	"github.com/google/uuid"
	"github.com/hadihalimm/cafebuzz-backend/internal/config"
	"github.com/hadihalimm/cafebuzz-backend/internal/models"
)

type AccountRepository interface {
	Create(account *models.Account) (*models.Account, error)
	FindByUUID(uuid uuid.UUID) (*models.Account, error)
	FindByUsername(username string) (*models.Account, error)
	Update(account *models.Account) (*models.Account, error)
}

type accountRepository struct {
	db *config.Database
}

func NewAccountRepository(db *config.Database) AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) Create(account *models.Account) (*models.Account, error) {
	err := r.db.Gorm.Create(&account)
	if err.Error != nil {
		return nil, err.Error
	}
	return account, nil
}

func (r *accountRepository) FindByUUID(uuid uuid.UUID) (*models.Account, error) {
	var account models.Account
	err := r.db.Gorm.First(&account, uuid)
	if err.Error != nil {
		return nil, err.Error
	}
	return &account, nil

}

func (r *accountRepository) FindByUsername(username string) (*models.Account, error) {
	var account models.Account
	err := r.db.Gorm.Where("username = ?", username).First(&account)
	if err.Error != nil {
		return nil, err.Error
	}
	return &account, nil
}

func (r *accountRepository) Update(account *models.Account) (*models.Account, error) {
	err := r.db.Gorm.Save(&account)
	if err.Error != nil {
		return nil, err.Error
	}
	return account, nil
}
