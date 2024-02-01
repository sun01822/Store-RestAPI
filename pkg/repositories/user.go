package repositories

import (
	"Store_RestAPI/pkg/domain"
	"Store_RestAPI/pkg/models"
	"Store_RestAPI/pkg/utils"
	"errors"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}


// interface binding
func UserDBInstance(db *gorm.DB) domain.IUserRepo {
	return &userRepo{
		db: db,
	}
}

// LoginUser implements domain.IUserRepo interface
func (u *userRepo) LoginUser(user *models.User) error {
	// Find the user by user_name
	var existingUser models.User
	if err := u.db.Where("user_name = ?", user.Username).First(&existingUser).Error; err != nil {
		return err
	}

	// Compare the stored hashed password, with the hashed version of the password that was received
	if err := utils.ComparePasswords(existingUser.Password, []byte(user.Password)); err != nil {
		return err
	}
	// Otherwise, we are good to go, so return a nil error.
	return nil
}

// GetUsers implements domain.IUserRepo interface
func (u *userRepo) GetUsers(model *gorm.Model) ([]models.User, error) {
	var users []models.User
	var err error
	userID := model.ID

	if userID != 0 {
		err = u.db.Where("id = ?", userID).Find(&users).Error
	} else {
		err = u.db.Find(&users).Error
	}
	if err != nil {
		return []models.User{}, err
	}
	return users, nil
}

// GetUsers implements domain.IUserRepo interface
func (u *userRepo) GetUsersInfo(model *gorm.Model) ([]models.UserInfo, error) {
	var users []models.UserInfo
	var err error
	userID := model.ID

	if userID != 0 {
		err = u.db.Raw("SELECT * FROM users Inner Join addresses on users.address_id = addresses.id Inner Join geo_locations on addresses.geo_location_id = geo_locations.id  where users.id = ?", userID, "and users.deleted_at is null").Scan(&users).Error
	} else {
		err = u.db.Raw("SELECT * FROM users Inner Join addresses on users.address_id = addresses.id Inner Join geo_locations on addresses.geo_location_id = geo_locations.id where users.deleted_at is null").Scan(&users).Error
	}
	if err != nil {
		return []models.UserInfo{}, err
	}
	return users, nil
}

// CreateUser implements domain.IUserRepo interface
func (u *userRepo) CreateUser(user *models.User) error {
	userEmail := user.Email
	userName := user.Username

	// Check if the user already exists
	var existingUser models.User
	if err := u.db.Where("email = ?", userEmail).First(&existingUser).Error; err == nil {
		return errors.New("email already exists")
	}
	if err := u.db.Where("user_name = ?", userName).First(&existingUser).Error; err == nil {
		return errors.New("username already exists")
	}

	// HashPassword
	user.Password = utils.HashPassword(user.Password)

	err := u.db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser implements domain.IUserRepo interface
func (u *userRepo) DeleteUser(model *gorm.Model) error {
	var user models.User
	err := u.db.Where("id = ?", model.ID).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateUser implements domain.IUserRepo interface
func (u *userRepo) UpdateUser(user *models.User) error {
	err := u.db.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// CreateAddress implements domain.IUserRepo.
func (u *userRepo) CreateAddress(address *models.Address) error {
	err := u.db.Create(&address).Error
	if err != nil {
		return err
	}
	return nil
}

// GetAddress implements domain.IUserRepo.
func (u *userRepo) GetAddress(model *gorm.Model) ([]models.Address, error) {
	var address []models.Address
	var err error
	addressID := model.ID
	if addressID == 0 {
		err = u.db.Find(&address).Error
	} else {
		err = u.db.Where("id = ?", addressID).Find(&address).Error
	}
	if err != nil {
		return []models.Address{}, err
	}
	return address, nil
}

// UpdateAddress implements domain.IUserRepo.
func (u *userRepo) UpdateAddress(address *models.Address) error {
	err := u.db.Save(&address).Error
	if err != nil {
		return err
	}
	return nil
}

// CreateGeoLocation implements domain.IUserRepo.
func (u *userRepo) CreateGeoLocation(geoLocation *models.GeoLocation) error {
	err := u.db.Create(&geoLocation).Error
	if err != nil {
		return err
	}
	return nil
}


// UpdateGeoLocation implements domain.IUserRepo.
func (u *userRepo) UpdateGeoLocation(geoLocation *models.GeoLocation) error {
	err := u.db.Save(&geoLocation).Error
	if err != nil {
		return err
	}
	return nil
}
