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
	FindActiveUsers() []entity.User
	FindNotActiveUsers() []entity.User
	SaveUser(ctx *gin.Context) error
	FindUser(ctx *gin.Context, sentId int) (entity.User, error)
	LoginUser(ctx *gin.Context) (entity.User, error)
	EditUser(ctx *gin.Context) error
	DeleteUser(ctx *gin.Context) error
	UserAddAddress(ctx *gin.Context) error
	UserDeleteAddress(ctx *gin.Context) error
	UserAddressesList(ctx *gin.Context, sentId int) ([]entity.Address, error)
	UserCircles(ctx *gin.Context, sentId int) (entity.UserCirclesResponse, error)
	ActivateUser(ctx *gin.Context) error
	DeactivateUser(ctx *gin.Context) error

	SpecializeUser(ctx *gin.Context) error
	NormalizeUser(ctx *gin.Context) error
	ChangeUserRole(ctx *gin.Context) error
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

func (c *userController) FindActiveUsers() []entity.User {
	return c.service.FindActiveUsers()
}

func (c *userController) FindNotActiveUsers() []entity.User {
	return c.service.FindNotActiveUsers()
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

	err = c.service.Save(user)
	if err != nil {
		return err
	}
	return nil
}

func (c *userController) FindUser(ctx *gin.Context, sentId int) (entity.User, error) {

	if sentId == 0 {
		var user entity.User
		return user, errors.New("user id cannot be zero")
		// var userId entity.UserInfoRequest
		// err := ctx.ShouldBindJSON(&userId)
		// if err != nil {
		// 	return user, err
		// }
		// user, err = c.service.FindUser(userId)
		// if err != nil {
		// 	return user, err
		// }
		// return user, nil
	} else {
		userId := entity.UserInfoRequest{
			ID: sentId,
		}
		var user entity.User
		// err := ctx.ShouldBindJSON(&userId)
		// if err != nil {
		// 	return user, err
		// }
		user, err := c.service.FindUser(userId)
		if err != nil {
			return user, err
		}
		return user, nil
	}
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

func (c *userController) ActivateUser(ctx *gin.Context) error {
	var userId entity.UserInfoRequest
	err := ctx.ShouldBindJSON(&userId)
	if err != nil {
		return err
	}
	err = validate.Struct(userId)
	if err != nil {
		return err
	}
	err = c.service.ActivateUser(userId)
	if err != nil {
		return err
	}
	return nil
}

func (c *userController) DeactivateUser(ctx *gin.Context) error {
	var userId entity.UserInfoRequest
	err := ctx.ShouldBindJSON(&userId)
	if err != nil {
		return err
	}
	err = validate.Struct(userId)
	if err != nil {
		return err
	}
	err = c.service.DeactivateUser(userId)
	if err != nil {
		return err
	}
	return nil
}

func (c *userController) SpecializeUser(ctx *gin.Context) error {
	var userId entity.UserInfoRequest
	err := ctx.ShouldBindJSON(&userId)
	if err != nil {
		return err
	}
	err = validate.Struct(userId)
	if err != nil {
		return err
	}
	err = c.service.SpecializeUser(userId)
	if err != nil {
		return err
	}
	return nil
}

func (c *userController) NormalizeUser(ctx *gin.Context) error {
	var userId entity.UserInfoRequest
	err := ctx.ShouldBindJSON(&userId)
	if err != nil {
		return err
	}
	err = validate.Struct(userId)
	if err != nil {
		return err
	}
	err = c.service.NormalizeUser(userId)
	if err != nil {
		return err
	}
	return nil
}

func (c *userController) ChangeUserRole(ctx *gin.Context) error {
	var userId entity.UserChangeRoleRequest
	err := ctx.ShouldBindJSON(&userId)
	if err != nil {
		return err
	}
	err = validate.Struct(userId)
	if err != nil {
		return err
	}
	err = c.service.ChangeUserRole(userId)
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

func (c *userController) UserAddressesList(ctx *gin.Context, sentId int) ([]entity.Address, error) {

	addressUserID := entity.UserAddressesRequest{
		UserId: sentId,
	}
	var address []entity.Address
	// err := ctx.ShouldBindJSON(&addressUserID)
	// if err != nil {
	// 	return address, err
	// }
	address, err := c.service.FindUserAddresses(addressUserID)
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

func (c *userController) UserDeleteAddress(ctx *gin.Context) error {

	var addressInfo entity.UserDeleteAddressRequest
	err := ctx.ShouldBindJSON(&addressInfo)
	if err != nil {
		return err
	}
	err = c.service.UserDeleteAddress(addressInfo)
	/*if err != nil {
		return err
	}*/
	return err
}

func (c *userController) UserCircles(ctx *gin.Context, sentId int) (entity.UserCirclesResponse, error) {

	addressUserID := entity.UserInfoRequest{
		ID: sentId,
	}
	// err := ctx.ShouldBindJSON(&addressUserID)
	// if err != nil {
	// 	return entity.UserCirclesResponse{}, err
	// }
	circles, err := c.service.UserCircles(addressUserID)
	return circles, err
}
