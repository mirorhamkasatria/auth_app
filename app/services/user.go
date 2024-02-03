package services

import (
	"errors"
	"fmt"

	"github.com/auth_app/app/models"
	r "github.com/auth_app/app/repositories"
	"github.com/auth_app/pkg/utils"
)

type ServiceUser interface {
	Login(user *models.User) (*string, error) //
	RegisterUser(user *models.User) (*string, error)
}

type serviceUser struct {
	repo r.RepoUser
}

func NewServiceUser(
	repo r.RepoUser) ServiceUser {
	return &serviceUser{
		repo: repo,
	}
}

func (s *serviceUser) Login(user *models.User) (*string, error) {
	data, err := s.repo.GetUsersByFilters(&models.User{
		Email: user.Email,
	})
	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(user.Password, (*data)[0].Password) {
		return nil, fmt.Errorf("invalid email and password")
	}

	token, err := utils.GenerateNewAccessToken(&(*data)[0])
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *serviceUser) RegisterUser(user *models.User) (*string, error) {

	data, err := s.repo.GetUsersByFilters(&models.User{
		Email: user.Email,
	})
	if err != nil {
		return nil, err
	}
	if len(*data) > 0 {
		return nil, errors.New("email already exist")
	}

	user.Password, _ = utils.HashPassword(user.Password)
	userId, err := s.repo.InsertUser(user)
	if err != nil {
		return nil, err
	}

	user.ID = *userId
	token, err := utils.GenerateNewAccessToken(user)
	if err != nil {
		return nil, err
	}

	return token, nil
}
