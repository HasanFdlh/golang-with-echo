package service

import (
	"ms-golang-echo/internal/model"
	"ms-golang-echo/internal/repository"
)

type UserService interface {
	GetByID(id uint) (*model.User, error)
	List() ([]model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) GetByID(id uint) (*model.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) List() ([]model.User, error) {
	return s.repo.List()
}
