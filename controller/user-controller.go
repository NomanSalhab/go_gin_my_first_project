package controller

import (
	"errors"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
	"github.com/NomanSalhab/go_gin_my_first_project/service"
	"github.com/NomanSalhab/go_gin_my_first_project/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController interface {
	FindAllUsers() []entity.User
	SaveUser(ctx *gin.Context) error
	FindUser(ctx *gin.Context) (entity.User, error)
	LoginUser(ctx *gin.Context) (entity.User, error)
	EditUser(ctx *gin.Context) error
	DeleteUser(ctx *gin.Context) error
	UserAddAddress(ctx *gin.Context) error
	UserAddressesList(ctx *gin.Context) ([]entity.Address, error)
	UserCircles(ctx *gin.Context) (entity.UserCirclesResponse, error)
}

type userController struct {
	service service.UserService
}

var validate *validator.Validate

func NewUserController(service service.UserService) UserController {
	validate = validator.New()
	validate.RegisterValidation("is-full-name", validators.ValidateFullUserName)
	return &userController{
		service: service,
	}
}

func (c *userController) FindAllUsers() []entity.User {
	return c.service.FindAll()
}

func (c *userController) SaveUser(ctx *gin.Context) error {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}

	err = validate.Struct(user)
	if err != nil {
		return err
	}

	if c.service.Save(user).Phone == "" {
		return errors.New("user phone number is already registered")
	}
	return nil
}

func (c *userController) FindUser(ctx *gin.Context) (entity.User, error) {

	var userId entity.UserInfoRequest
	var user entity.User
	err := ctx.ShouldBindJSON(&userId)
	if err != nil {
		return user, err
	}
	user, err = c.service.FindUser(userId)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (c *userController) LoginUser(ctx *gin.Context) (entity.User, error) {

	var userAuth entity.UserLoginRequest
	var user entity.User
	err := ctx.ShouldBindJSON(&userAuth)
	if err != nil {
		return user, err
	}
	user, err = c.service.LoginUser(userAuth)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (c *userController) EditUser(ctx *gin.Context) error {

	// var userId entity.UserInfoRequest
	var user entity.UserEditRequest
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}
	err = c.service.EditUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (c *userController) DeleteUser(ctx *gin.Context) error {

	// var userId entity.UserInfoRequest
	var user entity.UserInfoRequest
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}
	err = c.service.DeleteUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (c *userController) UserAddressesList(ctx *gin.Context) ([]entity.Address, error) {

	var addressUserID entity.UserAddressesRequest
	var address []entity.Address
	err := ctx.ShouldBindJSON(&addressUserID)
	if err != nil {
		return address, err
	}
	address, err = c.service.FindUserAddresses(addressUserID)
	if err != nil {
		return address, err
	}
	return address, nil
}

func (c *userController) UserAddAddress(ctx *gin.Context) error {

	var addressInfo entity.AddAddressRequest
	err := ctx.ShouldBindJSON(&addressInfo)
	if err != nil {
		return err
	}
	err = c.service.UserAddAddress(addressInfo)
	/*if err != nil {
		return err
	}*/
	return err
}

func (c *userController) UserCircles(ctx *gin.Context) (entity.UserCirclesResponse, error) {

	var addressUserID entity.UserInfoRequest
	err := ctx.ShouldBindJSON(&addressUserID)
	if err != nil {
		return entity.UserCirclesResponse{}, err
	}
	circles, err := c.service.UserCircles(addressUserID)
	return circles, err
}
