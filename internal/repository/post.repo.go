package repository

import (
	"github.com/google/uuid"
	"github.com/hadihalimm/cafebuzz-backend/internal/config"
	"github.com/hadihalimm/cafebuzz-backend/internal/models"
)

type PostRepository interface {
	Create(post *models.Post) (*models.Post, error)
	FindByID(id uint64) (*models.Post, error)
	FindAllByCreator(creatorUUID uuid.UUID) ([]*models.Post, error)
	Delete(id uint64) error
}

type postRepository struct {
	db *config.Database
}

func NewPostRepository(db *config.Database) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) Create(post *models.Post) (*models.Post, error) {
	err := r.db.Gorm.Create(&post)
	if err.Error != nil {
		return nil, err.Error
	}
	return post, nil
}

func (r *postRepository) FindByID(id uint64) (*models.Post, error) {
	var post models.Post
	err := r.db.Gorm.First(&post, id)
	if err.Error != nil {
		return nil, err.Error
	}
	return &post, nil
}

func (r *postRepository) FindAllByCreator(creatorUUID uuid.UUID) ([]*models.Post, error) {
	var post []*models.Post
	err := r.db.Gorm.Order("created_at desc").Where("creator_uuid = ?", creatorUUID).Find(&post)
	if err.Error != nil {
		return nil, err.Error
	}
	return post, nil
}

func (r *postRepository) Delete(id uint64) error {
	var post models.Post
	err := r.db.Gorm.Delete(post, id)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
