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
	err := u.db.Find(&users).Error
	if err != nil {
		return []models.User{}, err
	}
	return users, nil
}

// GetUserByID implements domain.IUserRepo interface
func (u *userRepo) GetUserByID(ID uint) (models.User, error) {
	var user models.User
	err := u.db.Where("id = ?", ID).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
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
