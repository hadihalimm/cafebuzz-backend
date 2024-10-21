package account

import "github.com/hadihalimm/cafebuzz-backend/internal/config"

type Repository interface {
	Create(account Account) error
}

type repository struct {
	db config.Database
}

func NewRepository(db config.Database) Repository {
	return &repository{db: db}
}

func (r *repository) Create(account Account) error {
	err := r.db.Create(&account)
	if err != nil {
		return err
	}
	return nil
}
