package repository

import (
	"github.com/google/uuid"
	"github.com/hadihalimm/cafebuzz-backend/internal/config"
	"github.com/hadihalimm/cafebuzz-backend/internal/models"
)

type PersonalAccountRepository interface {
	Create(account *models.PersonalAccount) (*models.PersonalAccount, error)
	FindByUUID(uuid uuid.UUID) (*models.PersonalAccount, error)
	FindByUsername(username string) (*models.PersonalAccount, error)
	Update(account *models.PersonalAccount) (*models.PersonalAccount, error)
	Delete(uuid uuid.UUID) error
}

type personalAccountRepository struct {
	db *config.Database
}

func NewPersonalAccountRepository(db *config.Database) PersonalAccountRepository {
	return &personalAccountRepository{db: db}
}

func (r *personalAccountRepository) Create(account *models.PersonalAccount) (*models.PersonalAccount, error) {
	err := r.db.Gorm.Create(&account)
	if err.Error != nil {
		return nil, err.Error
	}
	return account, nil
}

func (r *personalAccountRepository) FindByUUID(uuid uuid.UUID) (*models.PersonalAccount, error) {
	var account models.PersonalAccount
	err := r.db.Gorm.First(&account, uuid)
	if err.Error != nil {
		return nil, err.Error
	}
	return &account, nil

}

func (r *personalAccountRepository) FindByUsername(username string) (*models.PersonalAccount, error) {
	var account models.PersonalAccount
	err := r.db.Gorm.Where("username = ?", username).First(&account)
	if err.Error != nil {
		return nil, err.Error
	}
	return &account, nil
}

func (r *personalAccountRepository) Update(account *models.PersonalAccount) (*models.PersonalAccount, error) {
	err := r.db.Gorm.Save(&account)
	if err.Error != nil {
		return nil, err.Error
	}
	return account, nil
}

func (r *personalAccountRepository) Delete(uuid uuid.UUID) error {
	var account models.PersonalAccount
	err := r.db.Gorm.Delete(&account, uuid)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
