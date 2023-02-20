package service

import (
	"errors"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

type UserService interface {
	Save(entity.User) entity.User
	FindAll() []entity.User
	FindUser(id entity.UserInfoRequest) (entity.User, error)
	LoginUser(userAuth entity.UserLoginRequest) (entity.User, error)
	EditUser(user entity.UserEditRequest) error
	DeleteUser(user entity.UserInfoRequest) error
}

type userService struct {
	users []entity.User
}

func NewUserService() UserService {
	return &userService{}
}

func (service *userService) Save(user entity.User) entity.User {
	var failUser entity.User
	/*if user.ID == 0 {
		return failUser
	}*/
	for i := 0; i < len(service.users); i++ {
		if service.users[i].Phone == user.Phone {
			return failUser
		}
	}
	seccessUser := user
	if len(service.users) > 0 {
		seccessUser.ID = service.users[len(service.users)-1].ID + 1
	} else {
		seccessUser.ID = 1
	}
	service.users = append(service.users, seccessUser)
	return seccessUser
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

func (service *userService) LoginUser(userAuth entity.UserLoginRequest) (entity.User, error) {
	users := service.FindAll()
	var user entity.User
	for i := 0; i < len(users) && len(users) != 0; i++ {
		if users[i].Phone == userAuth.Phone {
			if users[i].Password == userAuth.Password {
				user = users[i]
				return users[i], nil
			} else {
				return user, errors.New("password is wrong")
			}
		}
	}
	if user.Name == "" {
		return user, errors.New("phone number is wrong")
	}
	return user, errors.New("failed")
}

func (service *userService) EditUser(user entity.UserEditRequest) error {
	users := service.FindAll()
	if user.ID == 0 {
		return errors.New("user id cannot be zero")
	}
	for i := 0; i < len(users) && len(users) != 0; i++ {
		if user.ID != 0 {
			if users[i].ID == user.ID {
				if user.Name != "" {
					users[i].Name = user.Name
				}
				if user.Balance != 0 {
					users[i].Balance = user.Balance
				}

				if user.Password != "" {
					users[i].Password = user.Password
				}
				/*users[i] = entity.User{
					ID:        user.ID,
					Name:      user.Name,
					Phone:     user.Phone,
					Password:  user.Password,
					Addresses: user.Addresses,
					Balance:   user.Balance,
				}*/
				return nil
			}
		}
	}
	// if /*user.ID == 0 || */ user.Name == "" {
	// 	return errors.New("the user couldn't be found")
	// }
	return errors.New("user could not be found")
}

func (service *userService) DeleteUser(user entity.UserInfoRequest) error {
	users := service.FindAll()
	var tempUsers []entity.User
	if user.ID == 0 {
		return errors.New("user id cannot be zero")
	}
	for i := 0; i < len(users) && len(users) != 0; i++ {
		if user.ID != 0 {
			if users[i].ID != user.ID {
				tempUsers = append(tempUsers, users[i])
			}
		}
	}
	if len(users) != len(tempUsers)+1 {
		return errors.New("user could not be found")
	}
	service.users = tempUsers
	return nil
}
