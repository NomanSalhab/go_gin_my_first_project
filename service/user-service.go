package service

import (
	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

type UserService interface {
	Save(entity.User) entity.User
	FindAll() []entity.User
}

type userService struct {
	users []entity.User
}

func New() UserService {
	return &userService{}
}

func (service *userService) Save(video entity.User) entity.User {
	service.users = append(service.users, video)
	return video
}

func (service *userService) FindAll() []entity.User {
	return service.users
}
