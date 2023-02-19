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
	user, err1 := c.service.FindUser(userId)
	if err1 != nil {
		return user, err
	}
	return user, nil
}
