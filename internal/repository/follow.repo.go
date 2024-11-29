package repository

import (
	"github.com/google/uuid"
	"github.com/hadihalimm/cafebuzz-backend/internal/config"
	"github.com/hadihalimm/cafebuzz-backend/internal/models"
)

type FollowRepository interface {
	Create(follow *models.Follow) (*models.Follow, error)
	FindFollowingByUUID(uuid uuid.UUID) ([]*models.PersonalAccount, []*models.CafeAccount, error)
	FindFollowersByUUID(uuid uuid.UUID) ([]*models.PersonalAccount, []*models.CafeAccount, error)
	Delete(followerUUID uuid.UUID, followedUUID uuid.UUID) error
}

type followRepository struct {
	db *config.Database
}

func NewFollowRepository(db *config.Database) FollowRepository {
	return &followRepository{db: db}
}

func (r *followRepository) Create(follow *models.Follow) (*models.Follow, error) {
	err := r.db.Gorm.Create(&follow)
	if err.Error != nil {
		return nil, err.Error
	}

	return follow, nil
}

func (r *followRepository) FindFollowingByUUID(uuid uuid.UUID) ([]*models.PersonalAccount, []*models.CafeAccount, error) {
	var followings []models.Follow
	var followedAccounts []*models.PersonalAccount
	var followedCafes []*models.CafeAccount

	if err := r.db.Gorm.Order("created_at desc").Where("follower_uuid = ?", uuid).Find(&followings).Error; err != nil {
		return nil, nil, err
	}

	for _, f := range followings {
		if f.FollowType == "personal" {
			var p models.PersonalAccount
			if err := r.db.Gorm.First(&p, f.FollowedUUID).Error; err == nil {
				followedAccounts = append(followedAccounts, &p)
			}
		} else if f.FollowType == "cafe" {
			var c models.CafeAccount
			if err := r.db.Gorm.First(&c, f.FollowedUUID).Error; err == nil {
				followedCafes = append(followedCafes, &c)
			}
		}
	}
	return followedAccounts, followedCafes, nil
}

func (r *followRepository) FindFollowersByUUID(uuid uuid.UUID) ([]*models.PersonalAccount, []*models.CafeAccount, error) {
	var followers []models.Follow
	var followersAccounts []*models.PersonalAccount
	var followersCafes []*models.CafeAccount

	if err := r.db.Gorm.Order("created_at desc").Where("followed_uuid = ?", uuid).Find(&followers).Error; err != nil {
		return nil, nil, err
	}

	for _, f := range followers {
		if f.FollowType == "personal" {
			var p models.PersonalAccount
			if err := r.db.Gorm.First(&p, f.FollowerUUID).Error; err == nil {
				followersAccounts = append(followersAccounts, &p)
			}
		} else if f.FollowType == "cafe" {
			var c models.CafeAccount
			if err := r.db.Gorm.First(&c, f.FollowerUUID).Error; err == nil {
				followersCafes = append(followersCafes, &c)
			}
		}
	}
	return followersAccounts, followersCafes, nil
}

func (r *followRepository) Delete(followerUUID uuid.UUID, followedUUID uuid.UUID) error {
	var follow models.Follow
	err := r.db.Gorm.Where("follower_uuid = ? AND followed_uuid = ?", followerUUID, followedUUID).Delete(&follow)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
