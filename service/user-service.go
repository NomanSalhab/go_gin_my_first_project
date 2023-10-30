package service

import (
	"errors"

	"github.com/NomanSalhab/go_gin_my_first_project/driver"
	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

type UserService interface {
	Save(entity.User) error
	FindAll() []entity.User
	FindActiveUsers() []entity.User
	FindNotActiveUsers() []entity.User
	FindUser(id entity.UserInfoRequest) (entity.User, error)
	LoginUser(userAuth entity.UserLoginRequest) (entity.User, error)
	EditUser(user entity.UserEditRequest) error
	DeleteUser(user entity.UserInfoRequest) error
	FindUserAddresses(addressUserId entity.UserAddressesRequest) ([]entity.Address, error)
	UserAddAddress(addedAddress entity.AddAddressRequest) error
	UserDeleteAddress(addedAddress entity.UserDeleteAddressRequest) error
	UserCircles(addressUserId entity.UserInfoRequest) (entity.UserCirclesResponse, error)
	ActivateUser(userInfo entity.UserInfoRequest) error
	DeactivateUser(userInfo entity.UserInfoRequest) error

	SpecializeUser(userInfo entity.UserInfoRequest) error
	NormalizeUser(userInfo entity.UserInfoRequest) error
	ChangeUserRole(userInfo entity.UserChangeRoleRequest) error
}

type userService struct {
	driver driver.UserDriver
	// addresses []entity.Address
}

func NewUserService(driver driver.UserDriver) UserService {
	return &userService{
		driver: driver,
	}
}

func (service *userService) Save(user entity.User) /*entity.User*/ error {
	usersList, err := service.driver.FindAllUsers()
	if err != nil {
		return err
	}
	for i := 0; i < len(usersList); i++ {
		if usersList[i].Phone == user.Phone {
			return errors.New("user phone number is already registered")
		}
	}
	err = service.driver.AddUser(user)
	if err != nil {
		return err
	}
	return nil
	// 	var failUser entity.User
	// 	for i := 0; i < len(service.users); i++ {
	// 		if service.users[i].Phone == user.Phone {
	// 			return failUser
	// 		}
	// 	}
	// 	seccessUser := user
	// 	if len(service.users) > 0 {
	// 		seccessUser.ID = service.users[len(service.users)-1].ID + 1
	// 	} else {
	// 		seccessUser.ID = 1
	// 	}
	// 	service.users = append(service.users, seccessUser)
	// 	return seccessUser
}

func (service *userService) FindAll() []entity.User {
	allUsers, err := service.driver.FindAllUsers()
	if err != nil {
		return make([]entity.User, 0)
	}
	return allUsers
}

func (service *userService) FindActiveUsers() []entity.User {
	activeUsers, err := service.driver.FindActiveUsers()
	if err != nil {
		return make([]entity.User, 0)
	}
	return activeUsers
	// var activeUsers []entity.User
	// for i := 0; i < len(service.users); i++ {
	// 	if service.users[i].Active {
	// 		activeUsers = append(activeUsers, service.users[i])
	// 	}
	// }
	// return activeUsers
}

func (service *userService) FindNotActiveUsers() []entity.User {
	notActiveUsers, err := service.driver.FindNotActiveUsers()
	if err != nil {
		return make([]entity.User, 0)
	}
	return notActiveUsers
	// var notActiveUsers []entity.User
	// for i := 0; i < len(service.users); i++ {
	// 	if !service.users[i].Active {
	// 		notActiveUsers = append(notActiveUsers, service.users[i])
	// 	}
	// }
	// return notActiveUsers
}

func (service *userService) FindUser(id entity.UserInfoRequest) (entity.User, error) {
	// users := service.FindAll()
	// var user entity.User
	// if id.ID != 0 {
	// 	for i := 0; i < len(users) && len(users) != 0; i++ {
	// 		if users[i].ID == id.ID {
	// 			user = users[i]
	// 		}
	// 	}
	// } else {
	// 	return user, errors.New("user id cannot be zero")
	// }
	user, err := service.driver.FindUser(id.ID)
	if err != nil {
		return entity.User{}, err
	}
	if user.Name == "" {
		return user, errors.New("the user couldn't be found")
	}
	return user, nil
}

func (service *userService) FindUserAddresses(addressUserId entity.UserAddressesRequest) ([]entity.Address, error) {
	// users := service.FindAll()
	addresses, err := service.driver.FindUserAddresses(addressUserId.UserId)
	if err != nil {
		return make([]entity.Address, 0), err
	}
	// if addressUserId.UserId != 0 {
	// 	for i := 0; i < len(service.addresses) && len(service.addresses) != 0; i++ {
	// 		if service.addresses[i].UserId == addressUserId.UserId {
	// 			addresses = append(addresses, service.addresses[i])
	// 		}
	// 	}
	// } else {
	// 	return addresses, errors.New("user id cannot be zero")
	// }
	if len(addresses) == 0 {
		return make([]entity.Address, 0), nil
	}
	return addresses, nil
}

func (service *userService) UserAddAddress(addedAddress entity.AddAddressRequest) error {
	err := service.driver.UserAddAddress(addedAddress)
	if err != nil {
		return err
	}
	return nil
	// users := service.FindAll()
	// address := entity.Address{
	// 	UserId:    addedAddress.UserId,
	// 	Name:      addedAddress.Name,
	// 	Latitude:  addedAddress.Latitude,
	// 	Longitude: addedAddress.Longitude,
	// }
	// if addedAddress.UserId != 0 {
	// 	for i := 0; i < len(users) && len(users) != 0; i++ {
	// 		if users[i].ID == addedAddress.UserId {
	// 			if len(service.addresses) > 0 {
	// 				address.ID = service.addresses[len(service.addresses)-1].ID + 1
	// 			} else {
	// 				address.ID = 1
	// 			}
	// 			// if len(users[i].Addresses) > 0 {
	// 			// 	address.ID = users[i].Addresses[len(users[i].Addresses)-1].ID + 1
	// 			// } else {
	// 			// 	address.ID = 1
	// 			// }
	// 			if address.ID != 0 {
	// 				// users[i].Addresses = append(users[i].Addresses, address)
	// 				service.addresses = append(service.addresses, address)
	// 				return nil
	// 			} else {
	// 				return errors.New("failed to add the address")
	// 			}
	// 		}
	// 	}
	// } else {
	// 	return errors.New("user id cannot be zero")
	// }
	// return errors.New("user is not found")
}

func (service *userService) UserDeleteAddress(deletedAddress entity.UserDeleteAddressRequest) error {
	err := service.driver.UserDeleteAddress(deletedAddress.ID)
	if err != nil {
		return err
	}
	return nil
	// // users := service.FindAll()
	// var tempAddresses []entity.Address
	// if deletedAddress.ID != 0 {
	// 	for i := 0; i < len(service.addresses) && len(service.addresses) != 0; i++ {
	// 		if service.addresses[i].ID != deletedAddress.ID {
	// 			tempAddresses = append(tempAddresses, service.addresses[i])
	// 		}
	// 	}
	// 	if len(service.addresses) != len(tempAddresses)+1 {
	// 		return errors.New("address could not be found")
	// 	} else {
	// 		service.addresses = tempAddresses
	// 		return nil
	// 	}
	// } else {
	// 	return errors.New("address id cannot be zero")
	// }
	// // return nil
}

func (service *userService) UserCircles(userId entity.UserInfoRequest) (entity.UserCirclesResponse, error) {
	// users := service.FindAll()
	// rate := len(users)
	// index := 0
	// if userId.ID != 0 {
	// 	for i := 0; i < len(users) && len(users) != 0; i++ {
	// 		if users[i].ID == userId.ID {
	// 			index = i
	// 		}
	// 	}
	// 	for i := 0; i < len(users) && len(users) != 0; i++ {
	// 		if users[i].Circles <= users[index].Circles && users[i].ID != users[index].ID {
	// 			rate--
	// 		}
	// 	}
	// 	if index < len(users) {
	// 		return entity.UserCirclesResponse{
	// 			Circles: users[index].Circles,
	// 			Rate:    rate,
	// 		}, nil
	// 	}
	// } else {
	// 	return entity.UserCirclesResponse{}, errors.New("user id cannot be zero")
	// }
	// return entity.UserCirclesResponse{}, errors.New("user could not be found")
	userCircles, err := service.driver.FindUserCircles(userId.ID)
	if err != nil {
		return entity.UserCirclesResponse{}, err
	}
	return userCircles, nil
}

func (service *userService) LoginUser(userAuth entity.UserLoginRequest) (entity.User, error) {
	user, err := service.driver.LoginUser(userAuth)
	if err != nil {
		return user, err
	}
	if !user.Active {
		return user, errors.New("user is not activated")
	}
	return user, nil
}

func (service *userService) EditUser(userEditInfo entity.UserEditRequest) error {
	_, err := service.driver.EditUser(userEditInfo)
	if err != nil {
		return err
	}
	return nil
}

func (service *userService) ActivateUser(userInfo entity.UserInfoRequest) error {
	err := service.driver.ActivateUser(userInfo)
	if err != nil {
		return err
	}
	return nil
	// storeCategories := service.FindAllStoreCategories()
	// var storeCategory entity.StoreCategory
	// for i := 0; i < len(storeCategories) && len(storeCategories) != 0; i++ {
	// 	if storeCategoryEditInfo.ID != 0 {
	// 		if storeCategories[i].ID == storeCategoryEditInfo.ID {
	// 			storeCategory.ID = storeCategoryEditInfo.ID
	// 			storeCategories[i].Active = true
	// 			return nil
	// 		}
	// 	} else {
	// 		return errors.New("store category id cannot be zero")
	// 	}
	// }
	// return errors.New("the store category couldn't be found")
}

func (service *userService) DeactivateUser(userInfo entity.UserInfoRequest) error {
	err := service.driver.DeactivateUser(userInfo)
	if err != nil {
		return err
	}
	return nil
	// storeCategories := service.FindAllStoreCategories()
	// var storeCategory entity.StoreCategory
	// for i := 0; i < len(storeCategories) && len(storeCategories) != 0; i++ {
	// 	if storeCategoryEditInfo.ID != 0 {
	// 		if storeCategories[i].ID == storeCategoryEditInfo.ID {
	// 			storeCategory.ID = storeCategoryEditInfo.ID
	// 			storeCategories[i].Active = false
	// 			return nil
	// 		}
	// 	} else {
	// 		return errors.New("store category id cannot be zero")
	// 	}
	// }
	// return errors.New("the store category couldn't be found")
}

func (service *userService) SpecializeUser(userInfo entity.UserInfoRequest) error {
	err := service.driver.SpecializeUser(userInfo)
	if err != nil {
		return err
	}
	return nil
}

func (service *userService) NormalizeUser(userInfo entity.UserInfoRequest) error {
	err := service.driver.NormalizeUser(userInfo)
	if err != nil {
		return err
	}
	return nil
}

func (service *userService) ChangeUserRole(userInfo entity.UserChangeRoleRequest) error {
	if userInfo.Role == 0 || userInfo.Role == 1 || userInfo.Role == 2 {
		err := service.driver.ChangeUserRole(userInfo)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("specified role is invalid")
	}
}

func (service *userService) DeleteUser(user entity.UserInfoRequest) error {
	err := service.driver.DeleteUser(user.ID)
	if err != nil {
		return err
	}
	return nil
	// users := service.FindAll()
	// var tempUsers []entity.User
	// if user.ID != 0 {
	// 	for i := 0; i < len(users) && len(users) != 0; i++ {
	// 		if users[i].ID != user.ID {
	// 			tempUsers = append(tempUsers, users[i])
	// 		}
	// 	}
	// } else {
	// 	return errors.New("user id cannot be zero")
	// }
	// if len(users) != len(tempUsers)+1 {
	// 	return errors.New("user could not be found")
	// }
	// service.users = tempUsers
	// return nil
}
