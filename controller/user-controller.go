package controller

import (
	"net/http"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
	"github.com/NomanSalhab/go_gin_my_first_project/service"
	"github.com/NomanSalhab/go_gin_my_first_project/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController interface {
	FindAll() []entity.User
	Save(ctx *gin.Context) error /*entity.Video*/
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.UserService
}

var validate *validator.Validate

func New(service service.UserService) UserController {
	validate = validator.New()
	validate.RegisterValidation("is-full-name", validators.ValidateFullUserName)
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.User {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error /*entity.Video*/ {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}

	err = validate.Struct(user)
	if err != nil {
		return err
	}

	c.service.Save(user)
	return nil /*video*/
}

func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
