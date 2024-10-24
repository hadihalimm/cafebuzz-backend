package repository

import (
	"github.com/google/uuid"
	"github.com/hadihalimm/cafebuzz-backend/internal/config"
	"github.com/hadihalimm/cafebuzz-backend/internal/models"
)

type AccountRepository interface {
	Create(account *models.Account) (*models.Account, error)
	FindByUUID(uuid uuid.UUID) (*models.Account, error)
	FindByUsername(username string) (models.Account, error)
	Update(account *models.Account) (*models.Account, error)
}

type repository struct {
	db *config.Database
}

func NewAccountRepository(db *config.Database) AccountRepository {
	return &repository{db: db}
}

func (r *repository) Create(account *models.Account) (*models.Account, error) {
	result := r.db.Gorm.Create(&account)
	if result.Error != nil {
		return nil, result.Error
	}
	return account, nil
}

func (r *repository) FindByUUID(uuid uuid.UUID) (*models.Account, error) {
	var account models.Account
	result := r.db.Gorm.First(&account, uuid)
	if result.Error != nil {
		return nil, result.Error
	}
	return &account, nil

}

func (r *repository) FindByUsername(username string) (models.Account, error) {
	account := models.Account{}
	err := r.db.Gorm.Where("username = ?", username).First(&account)
	if err != nil {
		return account, err.Error
	}
	return account, nil
}

func (r *repository) Update(account *models.Account) (*models.Account, error) {
	result := r.db.Gorm.Save(&account)
	if result.Error != nil {
		return nil, result.Error
	}
	return account, nil
}
