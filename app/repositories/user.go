package repositories

import (
	"fmt"

	"github.com/auth_app/app/models"
	"gorm.io/gorm"
)

type RepoUser interface {
	GetUsersByFilters(filter *models.User) (*[]models.User, error)
	InsertUser(data *models.User) (*int, error)
}

type repoUser struct {
	db *gorm.DB
}

func NewRepoUser(db *gorm.DB) RepoUser {
	return &repoUser{
		db: db,
	}
}

func (r *repoUser) GetUsersByFilters(filter *models.User) (*[]models.User, error) {
	var Users []models.User
	query := r.db.Model(Users)
	if filter.Email != "" {
		query = query.Where("email = ? ", filter.Email)
	}

	if err := query.Find(&Users).Error; err != nil {
		return nil, fmt.Errorf("failed to get user by filters %w", err)
	}
	return &Users, nil
}

func (r *repoUser) InsertUser(data *models.User) (*int, error) {
	if err := r.db.Create(&data).Error; err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}

	return &data.ID, nil
}
