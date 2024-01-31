package controllers

import (
	//"Store_RestAPI/pkg/config"
	"Store_RestAPI/pkg/domain"
	"Store_RestAPI/pkg/models"
	"Store_RestAPI/pkg/types"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

// IUserController is an interface for user controller
type IUserController interface {
	GetUsers(e echo.Context) error
	CreateUser(e echo.Context) error
	DeleteUser(e echo.Context) error
	UpdateUser(e echo.Context) error
	LoginUser(e echo.Context) error
	CreateAddress(e echo.Context) error
	GetUsersInfo(e echo.Context) error
	UpdateAddress(e echo.Context) error
}

type userController struct {
	Usersvc domain.IUserService
}

// UserInstance is a function to create an instance of user controller
func NewUserController(UserSvc domain.IUserService) IUserController {
	return &userController{
		Usersvc: UserSvc,
	}
}

// CreateUser implements IUserController.
func (controller *userController) CreateUser(e echo.Context) error {
	reqUser := &types.UserRequest{}
	if err := e.Bind(reqUser); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid data")
	}
	if err := reqUser.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	user := &models.User{
		Email:    reqUser.Email,
		Username: reqUser.Username,
		Password: reqUser.Password,
	}
	if err := controller.Usersvc.CreateUser(user); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "User created successfully")
}

// GetUsers implements IUserController.
func (controller *userController) GetUsers(e echo.Context) error {
	tempUserID := e.QueryParam("userID")
	UserID, err := strconv.ParseInt(tempUserID, 0, 0)
	if err != nil && tempUserID != "" {
		return e.JSON(http.StatusBadRequest, "Enter a valid User ID")
	}
	users, err := controller.Usersvc.GetUsers(&gorm.Model{ID: uint(UserID)})
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, users)
}

// GetUsers implements IUserController.
func (controller *userController) GetUsersInfo(e echo.Context) error {
	tempUserID := e.QueryParam("userID")
	UserID, err := strconv.ParseInt(tempUserID, 0, 0)
	if err != nil && tempUserID != "" {
		return e.JSON(http.StatusBadRequest, "Enter a valid User ID")
	}
	users, err := controller.Usersvc.GetUsersInfo(&gorm.Model{ID: uint(UserID)})
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, users)
}

// DeleteUser implements IUserController.
func (controller *userController) DeleteUser(e echo.Context) error {
	tempUserID := e.Param("id")
	UserID, err := strconv.ParseInt(tempUserID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Enter a valid User ID")
	}
	_, err = controller.Usersvc.GetUsers(&gorm.Model{ID: uint(UserID)})
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	if err := controller.Usersvc.DeleteUser(&gorm.Model{ID: uint(UserID)}); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "User deleted successfully")
}

// UpdateUser implements IUserController.
func (controller *userController) UpdateUser(e echo.Context) error {
	reqUser := &types.UserRequest{}
	if err := e.Bind(reqUser); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid data")
	}
	tempUserID := e.Param("id")
	UserID, err := strconv.ParseInt(tempUserID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Enter a valid User ID")
	}
	existingUser, err := controller.Usersvc.GetUsers(&gorm.Model{ID: uint(UserID)})
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	updateUser := &models.User{
		Model:     gorm.Model{ID: uint(UserID), UpdatedAt: time.Now(), CreatedAt: existingUser[0].CreatedAt, DeletedAt: existingUser[0].DeletedAt},
		Email:     reqUser.Email,
		Username:  reqUser.Username,
		Password:  reqUser.Password,
		Phone:     reqUser.Phone,
		AddressID: reqUser.AddressID,
		NameID:    reqUser.NameID,
	}

	if updateUser.Username == "" {
		updateUser.Username = existingUser[0].Username
	}
	if updateUser.Email == "" {
		updateUser.Email = existingUser[0].Email
	}
	if updateUser.Password == "" {
		updateUser.Password = existingUser[0].Password
	}
	if err := controller.Usersvc.UpdateUser(updateUser); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "User updated successfully")
}

// LoginUser implements IUserController.
func (controller *userController) LoginUser(e echo.Context) error {
	//config := config.LocalConfig
	reqUser := &types.UserRequest{}
	if err := e.Bind(reqUser); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid data")
	}
	if err := reqUser.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	user := &models.User{
		Email:    reqUser.Email,
		Password: reqUser.Password,
	}
	if err := controller.Usersvc.LoginUser(user); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	now := time.Now().UTC()
	ttl := time.Minute * 15
	claims := jwt.StandardClaims{
		ExpiresAt: now.Add(ttl).Unix(),
		IssuedAt:  now.Unix(),
		NotBefore: now.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, tokenString)
}


// CreateAddress implements IUserController.
func (controller *userController) CreateAddress(e echo.Context) error {
	tempUserID := e.Param("id")
	UserID, err := strconv.ParseInt(tempUserID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Enter a valid User ID")
	}
	existingUser, err := controller.Usersvc.GetUsers(&gorm.Model{ID: uint(UserID)})
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	reqUser := &types.UserAddressRequest{}
	if err := e.Bind(reqUser); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid data")
	}
	if err := reqUser.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	address := &models.Address{
		City:    reqUser.City,
		Street:  reqUser.Street,
		Number:  reqUser.Number,
		Zip:     reqUser.Zip,
	}
	if err := controller.Usersvc.CreateAddress(address); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	updateUser := &models.User{
		Model:     gorm.Model{ID: uint(UserID), UpdatedAt: time.Now(), CreatedAt: existingUser[0].CreatedAt, DeletedAt: existingUser[0].DeletedAt},
		AddressID: address.ID,
		Username: existingUser[0].Username,
		Email: existingUser[0].Email,
		Password: existingUser[0].Password,
		Phone: existingUser[0].Phone,
		NameID: existingUser[0].NameID,
	}
	if err := controller.Usersvc.UpdateUser(updateUser); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "Address created successfully")
}

// UpdateAddress implements IUserController.
func (controller *userController) UpdateAddress(e echo.Context) error {
	reqUser := &types.UserAddressRequest{}
	if err := e.Bind(reqUser); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid data")
	}
	if err := reqUser.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	tempUserID := e.Param("id")
	UserID, err := strconv.ParseInt(tempUserID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Enter a valid User ID")
	}
	existingUser, err := controller.Usersvc.GetUsers(&gorm.Model{ID: uint(UserID)})
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	updateAddress := &models.Address{
		ID: 	 existingUser[0].AddressID,
		City:     reqUser.City,
		Street:   reqUser.Street,
		Number:   reqUser.Number,
		Zip:      reqUser.Zip,
	}
	if err := controller.Usersvc.UpdateAddress(updateAddress); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "Address updated successfully")
}