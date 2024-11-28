package repository

import (
	"github.com/google/uuid"
	"github.com/hadihalimm/cafebuzz-backend/internal/config"
	"github.com/hadihalimm/cafebuzz-backend/internal/models"
)

type CafeAccountRepository interface {
	Create(cafe *models.CafeAccount) (*models.CafeAccount, error)
	FindByUUID(uuid uuid.UUID) (*models.CafeAccount, error)
	FindByUsername(username string) (*models.CafeAccount, error)
	Update(cafe *models.CafeAccount) (*models.CafeAccount, error)
	Delete(uuid uuid.UUID) error
}

type cafeAccountRepository struct {
	db *config.Database
}

func NewCafeAccountRepository(db *config.Database) CafeAccountRepository {
	return &cafeAccountRepository{db: db}
}

func (r *cafeAccountRepository) Create(cafe *models.CafeAccount) (*models.CafeAccount, error) {
	err := r.db.Gorm.Create(&cafe)
	if err.Error != nil {
		return nil, err.Error
	}
	return cafe, nil
}

func (r *cafeAccountRepository) FindByUUID(uuid uuid.UUID) (*models.CafeAccount, error) {
	var cafe models.CafeAccount
	err := r.db.Gorm.First(&cafe, uuid)
	if err.Error != nil {
		return nil, err.Error
	}
	return &cafe, nil
}

func (r *cafeAccountRepository) FindByUsername(username string) (*models.CafeAccount, error) {
	var cafe models.CafeAccount
	err := r.db.Gorm.Where("username = ?", username).First(&cafe)
	if err.Error != nil {
		return nil, err.Error
	}
	return &cafe, nil
}

func (r *cafeAccountRepository) Update(cafe *models.CafeAccount) (*models.CafeAccount, error) {
	err := r.db.Gorm.Save(&cafe)
	if err.Error != nil {
		return nil, err.Error
	}
	return cafe, nil
}

func (r *cafeAccountRepository) Delete(uuid uuid.UUID) error {
	var cafe models.CafeAccount
	err := r.db.Gorm.Delete(&cafe, uuid)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
