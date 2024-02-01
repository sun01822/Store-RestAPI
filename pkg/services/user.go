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
		return nil, errors.New("no users found")
	}
	allUsers = append(allUsers, user...)
	return allUsers, nil
}

func (service *userService) GetUsersInfo(model *gorm.Model) ([]models.UserInfo, error) {
	var allUsers []models.UserInfo
	user, _ := service.repo.GetUsersInfo(model)
	if len(user) == 0 {
		return nil, errors.New("no users found")
	}
	allUsers = append(allUsers, user...)
	return allUsers, nil
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


// CreateAddress implements domain.IUserService.
func (service *userService) CreateAddress(address *models.Address) error {
	if err := service.repo.CreateAddress(address); err != nil {
		return errors.New("Address is not created")
	}
	return nil
}

// GetAddress implements domain.IUserService.
func (service *userService) GetAddress(model *gorm.Model) ([]models.Address, error) {
	var allAddress []models.Address
	address, _ := service.repo.GetAddress(model)
	if len(address) == 0 {
		return nil, errors.New("no address found")
	}
	allAddress = append(allAddress, address...)
	return allAddress, nil
}

// UpdateAddress implements domain.IUserService.
func (service *userService) UpdateAddress(address *models.Address) error {
	if err := service.repo.UpdateAddress(address); err != nil {
		return errors.New("Address is not updated")
	}
	return nil
}

// CreateGeoLocation implements domain.IUserService.
func (service *userService) CreateGeoLocation(geoLocation *models.GeoLocation) error {
	if err := service.repo.CreateGeoLocation(geoLocation); err != nil {
		return errors.New("GeoLocation is not created")
	}
	return nil
}

// UpdateGeoLocation implements domain.IUserService.
func (service *userService) UpdateGeoLocation(geoLocation *models.GeoLocation) error {
	if err := service.repo.UpdateGeoLocation(geoLocation); err != nil {
		return errors.New("GeoLocation is not updated")
	}
	return nil
}

// CreateName implements domain.IUserService.
func (service *userService) CreateName(name *models.Name) error {
	if err := service.repo.CreateName(name); err != nil {
		return errors.New("Name is not created")
	}
	return nil
}

// GetName implements domain.IUserService.
func (service *userService) GetName(model *gorm.Model) ([]models.Name, error) {
	var allName []models.Name
	name, _ := service.repo.GetName(model)
	if len(name) == 0 {
		return nil, errors.New("no name found")
	}
	allName = append(allName, name...)
	return allName, nil
}

// UpdateName implements domain.IUserService.
func (service *userService) UpdateName(name *models.Name) error {
	if err := service.repo.UpdateName(name); err != nil {
		return errors.New("Name is not updated")
	}
	return nil
}

