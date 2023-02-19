package service

import (
	"errors"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

type UserService interface {
	Save(entity.User) entity.User
	FindAll() []entity.User
	FindUser(id entity.UserInfoRequest) (entity.User, error)
}

type userService struct {
	users []entity.User
}

func NewUserService() UserService {
	return &userService{}
}

func (service *userService) Save(user entity.User) entity.User {
	var failUser entity.User
	for i := 0; i < len(service.users); i++ {
		if service.users[i].Phone == user.Phone {
			return failUser
		}
	}
	service.users = append(service.users, user)
	return user
}

func (service *userService) FindAll() []entity.User {
	return service.users
}

func (service *userService) FindUser(id entity.UserInfoRequest) (entity.User, error) {
	users := service.FindAll()
	var user entity.User
	for i := 0; i < len(users) && len(users) != 0; i++ {
		if id.ID != 0 {
			if users[i].ID == id.ID {
				user = users[i]
			}
		} else {
			return user, errors.New("user id cannot be zero")
		}
	}
	if id.ID == 0 {
		return user, errors.New("user id cannot be zero")
	}
	if /*user.ID == 0 || */ user.Name == "" {
		return user, errors.New("the user couldn't be found")
	}
	return user, nil
}
