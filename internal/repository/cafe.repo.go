package repository

import (
	"github.com/google/uuid"
	"github.com/hadihalimm/cafebuzz-backend/internal/config"
	"github.com/hadihalimm/cafebuzz-backend/internal/models"
)

type CafeRepository interface {
	Create(cafe *models.Cafe) (*models.Cafe, error)
	FindByUUID(uuid uuid.UUID) (*models.Cafe, error)
	FindByUsername(username string) (*models.Cafe, error)
	Update(cafe *models.Cafe) (*models.Cafe, error)
}

type cafeRepository struct {
	db *config.Database
}

func NewCafeRepository(db *config.Database) CafeRepository {
	return &cafeRepository{db: db}
}

func (r *cafeRepository) Create(cafe *models.Cafe) (*models.Cafe, error) {
	err := r.db.Gorm.Create(&cafe)
	if err.Error != nil {
		return nil, err.Error
	}
	return cafe, nil
}

func (r *cafeRepository) FindByUUID(uuid uuid.UUID) (*models.Cafe, error) {
	var cafe models.Cafe
	err := r.db.Gorm.First(&cafe, uuid)
	if err.Error != nil {
		return nil, err.Error
	}
	return &cafe, nil
}

func (r *cafeRepository) FindByUsername(username string) (*models.Cafe, error) {
	var cafe models.Cafe
	err := r.db.Gorm.Where("username = ?", username).First(&cafe)
	if err.Error != nil {
		return nil, err.Error
	}
	return &cafe, nil
}

func (r *cafeRepository) Update(cafe *models.Cafe) (*models.Cafe, error) {
	err := r.db.Gorm.Save(&cafe)
	if err.Error != nil {
		return nil, err.Error
	}
	return cafe, nil
}
