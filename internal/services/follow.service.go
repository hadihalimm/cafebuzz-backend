package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/hadihalimm/cafebuzz-backend/internal/models"
	"github.com/hadihalimm/cafebuzz-backend/internal/repository"
)

type FollowService interface {
	Create(followerUUID uuid.UUID, followedUUID uuid.UUID, followType string) (*models.Follow, error)
	FindFollowingsByUUID(uuid uuid.UUID) ([]*models.PersonalAccount, []*models.CafeAccount, error)
	FindFollowersByUUID(uuid uuid.UUID) ([]*models.PersonalAccount, []*models.CafeAccount, error)
	Delete(followerUUID uuid.UUID, followedUUID uuid.UUID) error
}

type followService struct {
	repo     repository.FollowRepository
	validate *validator.Validate
}

func NewFollowService(repo repository.FollowRepository, validate *validator.Validate) FollowService {
	return &followService{repo: repo, validate: validate}
}

func (s *followService) Create(followerUUID uuid.UUID, followedUUID uuid.UUID, followType string) (*models.Follow, error) {
	var follow models.Follow
	follow.FollowerUUID = followerUUID
	follow.FollowedUUID = followedUUID
	follow.FollowType = followType

	newFollow, err := s.repo.Create(&follow)
	if err != nil {
		return nil, err
	}
	return newFollow, nil
}

func (s *followService) FindFollowingsByUUID(uuid uuid.UUID) ([]*models.PersonalAccount, []*models.CafeAccount, error) {
	personalFollowing, cafeFollowing, err := s.repo.FindFollowingByUUID(uuid)
	if err != nil {
		return nil, nil, err
	}
	return personalFollowing, cafeFollowing, nil
}

func (s *followService) FindFollowersByUUID(uuid uuid.UUID) ([]*models.PersonalAccount, []*models.CafeAccount, error) {
	personalFollowers, cafeFollowers, err := s.repo.FindFollowersByUUID(uuid)
	if err != nil {
		return nil, nil, err
	}
	return personalFollowers, cafeFollowers, nil
}

func (s *followService) Delete(followerUUID uuid.UUID, followedUUID uuid.UUID) error {
	return s.repo.Delete(followerUUID, followedUUID)
}
