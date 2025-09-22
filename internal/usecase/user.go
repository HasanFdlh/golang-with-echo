// internal/usecase/user_usecase.go
package usecase

import (
	"ms-golang-echo/internal/model"
	"ms-golang-echo/internal/repository"
)

type UserUsecase interface {
	Create(user *model.User) error
	GetByID(id uint) (*model.User, error)
	List() ([]model.User, error)
	Update(user *model.User) error
	Delete(id uint) error
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) UserUsecase {
	return &userUsecase{repo: r}
}

func (u *userUsecase) Create(user *model.User) error {
	return u.repo.Create(user)
}

func (u *userUsecase) GetByID(id uint) (*model.User, error) {
	return u.repo.GetByID(id)
}

func (u *userUsecase) List() ([]model.User, error) {
	return u.repo.List()
}

func (u *userUsecase) Update(user *model.User) error {
	return u.repo.Update(user)
}

func (u *userUsecase) Delete(id uint) error {
	return u.repo.Delete(id)
}
