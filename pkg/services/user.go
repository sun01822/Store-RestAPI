package services

import (
	"Store_RestAPI/pkg/domain"
	"Store_RestAPI/pkg/models"
	"errors"
	"gorm.io/gorm"
)

// Parent stuct to implement IUserService interface binding
type userService struct {
	repo domain.IUserRepo
}


// Interface binding
func UserInstance(userRepo domain.IUserRepo) domain.IUserService {
	return &userService{
		repo: userRepo,
	}
}

// All the methods of IUserService interface are implemented here

// GetUsers implements domain.IUserService.
func (service *userService) GetUsers(model *gorm.Model) ([]models.User, error) {
	var allUsers []models.User
	user, _ := service.repo.GetUsers(model)
	if len(user) == 0 {
		return nil, errors.New("No users found")
	}
	allUsers = append(allUsers, user...)
	return allUsers, nil
}

// GetUsersByID implements domain.IUserService.
func (service *userService) GetUserByID(ID uint) (models.User, error) {
	user, err := service.repo.GetUserByID(ID)
	if err != nil {
		return models.User{}, errors.New("User not found")
	}
	return user, nil
}

// CreateUser implements domain.IUserService.
func (service *userService) CreateUser(user *models.User) error {
	if err := service.repo.CreateUser(user); err != nil {
		return err
	}
	return nil
}

// DeleteUser implements domain.IUserService.
func (service *userService) DeleteUser(model *gorm.Model) error {
	if err := service.repo.DeleteUser(model); err != nil {
		return errors.New("User is not deleted")
	}
	return nil
}

// UpdateUser implements domain.IUserService.
func (service *userService) UpdateUser(user *models.User) error {
	if err := service.repo.UpdateUser(user); err != nil {
		return errors.New("User is not updated")
	}
	return nil
}

// LoginUser implements domain.IUserService.
func (service *userService) LoginUser(user *models.User) error {
	if err := service.repo.LoginUser(user); err != nil {
		return errors.New("Log in Failed")
	}
	return nil
}