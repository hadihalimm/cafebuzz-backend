package account

import "github.com/hadihalimm/cafebuzz-backend/internal/config"

type Repository interface {
	Create(account Account) (Account, error)
	FindByUsername(username string) (Account, error)
}

type repository struct {
	db *config.Database
}

func NewRepository(db *config.Database) Repository {
	return &repository{db: db}
}

func (r *repository) Create(account Account) (Account, error) {
	result := r.db.Gorm.Create(&account)
	if result.Error != nil {
		return account, result.Error
	}
	return account, nil
}

func (r *repository) FindByUsername(username string) (Account, error) {
	account := Account{}
	err := r.db.Gorm.Where("username = ?", username).First(&account)
	if err != nil {
		return account, err.Error
	}
	return account, nil
}
