package domain

import (
	"Store_RestAPI/pkg/models"
	"gorm.io/gorm"
)

// for database Repository operation (call from service)
type IUserRepo interface {
	GetUsers(*gorm.Model)([]models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(*gorm.Model) error
	LoginUser(user *models.User) error
	GetUsersInfo(*gorm.Model)([]models.UserInfo, error)


	CreateAddress(address *models.Address) error
	GetAddress(*gorm.Model)([]models.Address, error)
	UpdateAddress(address *models.Address) error



	CreateGeoLocation(geoLocation *models.GeoLocation) error
	UpdateGeoLocation(geoLocation *models.GeoLocation) error

}

// for service operation (response to contorller || call from controller)
type IUserService interface {
	GetUsers(*gorm.Model)([]models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(*gorm.Model) error
	LoginUser(user *models.User) error
	GetUsersInfo(*gorm.Model)([]models.UserInfo, error)


	CreateAddress(address *models.Address) error
	GetAddress(*gorm.Model)([]models.Address, error)
	UpdateAddress(address *models.Address) error


	CreateGeoLocation(geoLocation *models.GeoLocation) error
	UpdateGeoLocation(geoLocation *models.GeoLocation) error
}