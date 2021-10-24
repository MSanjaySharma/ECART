package service

import (
	"ecommerce/api/repository"
	"ecommerce/models"
)

//UserService UserService struct
type UserService struct {
	repo repository.UserRepository
}

//NewUserService : get injected user repo
func NewUserService(repo repository.UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

//Save -> saves users entity
func (u UserService) CreateUser(user models.UserRegister) error {
	return u.repo.CreateUser(user)
}

//FindAll -> calls user repo find all method
func (u UserService) FindAll(user models.User, keyword string) (*[]models.User, int64, error) {
	return u.repo.FindAll(user, keyword)
}

//Login -> Gets validated user
func (u UserService) LoginUser(user models.UserLogin) (*models.User, error) {
	return u.repo.LoginUser(user)

}
