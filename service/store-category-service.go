package service

import (
	"errors"

	"github.com/NomanSalhab/go_gin_my_first_project/driver"
	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

type StoreCategoryService interface {
	AddStoreCategory(entity.StoreCategory) error
	FindAllStoreCategories() []entity.StoreCategory
	FindActiveStoreCategories() []entity.StoreCategory
	FindNotActiveStoreCategories() []entity.StoreCategory
	FindStoreCategory(id entity.StoreCategoryInfoRequest) (entity.StoreCategory, error)
	EditStoreCategory(storeCategoryEditInfo entity.StoreCategoryEditRequest) error
	DeactivateStoreCategory(storeCategoryEditInfo entity.StoreCategoryInfoRequest) error
	ActivateStoreCategory(storeCategoryEditInfo entity.StoreCategoryInfoRequest) error
	DeleteStoreCategory(user entity.StoreCategoryDeleteRequest) error

	AddMockStoreCategories(storeCategories []entity.StoreCategory)
}

type storeCategoryService struct {
	storeCategories []entity.StoreCategory
	driver          driver.StoreCategoryDriver
}

func NewStoreCategoryService(driver driver.StoreCategoryDriver) StoreCategoryService {
	return &storeCategoryService{
		driver: driver,
	}
}

func (service *storeCategoryService) AddStoreCategory(storeCategory entity.StoreCategory) error {
	storeCategoriesList, err := service.driver.FindAllStoreCategories()
	if err != nil {
		return err
	}
	for i := 0; i < len(storeCategoriesList); i++ {
		if storeCategoriesList[i].Name == storeCategory.Name {
			return errors.New("store category name already exists")
		}
	}
	err = service.driver.AddStoreCategory(storeCategory)
	if err != nil {
		return err
	}
	return nil
	// successStoreCategory := storeCategory
	// for i := 0; i < len(service.storeCategories); i++ {
	// 	if service.storeCategories[i].Name == storeCategory.Name {
	// 		return errors.New("store category already exisists")
	// 	}
	// }
	// if len(service.storeCategories) > 0 {
	// 	successStoreCategory.ID = service.storeCategories[len(service.storeCategories)-1].ID + 1
	// } else {
	// 	successStoreCategory.ID = 1
	// }
	// service.storeCategories = append(service.storeCategories, successStoreCategory)
	// return nil
}

func (service *storeCategoryService) FindAllStoreCategories() []entity.StoreCategory {
	allStoreCategories, err := service.driver.FindAllStoreCategories()
	if err != nil {
		return make([]entity.StoreCategory, 0)
	}
	return allStoreCategories
	// return service.storeCategories
}

func (service *storeCategoryService) FindActiveStoreCategories() []entity.StoreCategory {
	activeStoreCategories, err := service.driver.FindActiveStoreCategories()
	if err != nil {
		return make([]entity.StoreCategory, 0)
	}
	return activeStoreCategories
	// var activeStoreCategories []entity.StoreCategory
	// for i := 0; i < len(service.storeCategories); i++ {
	// 	if service.storeCategories[i].Active {
	// 		activeStoreCategories = append(activeStoreCategories, service.storeCategories[i])
	// 	}
	// }
	// return activeStoreCategories
}

func (service *storeCategoryService) FindNotActiveStoreCategories() []entity.StoreCategory {
	notActiveStoreCategories, err := service.driver.FindNotActiveStoreCategories()
	if err != nil {
		return make([]entity.StoreCategory, 0)
	}
	return notActiveStoreCategories
	// var notActiveStoreCategories []entity.StoreCategory
	// for i := 0; i < len(service.storeCategories); i++ {
	// 	if !service.storeCategories[i].Active {
	// 		notActiveStoreCategories = append(notActiveStoreCategories, service.storeCategories[i])
	// 	}
	// }
	// return notActiveStoreCategories
}

func (service *storeCategoryService) FindStoreCategory(id entity.StoreCategoryInfoRequest) (entity.StoreCategory, error) {
	storeCategory, _ := service.driver.FindStoreCategory(id.ID)
	if storeCategory.Name == "" {
		return storeCategory, errors.New("the store category couldn't be found")
	}
	return storeCategory, nil
	// storeCategories := service.FindAllStoreCategories()
	// var storeCategory entity.StoreCategory
	// for i := 0; i < len(storeCategories) && len(storeCategories) != 0; i++ {
	// 	if id.ID != 0 {
	// 		if storeCategories[i].ID == id.ID {
	// 			storeCategory = storeCategories[i]
	// 			return storeCategory, nil
	// 		}
	// 	} else {
	// 		return storeCategory, errors.New("store category id cannot be zero")
	// 	}
	// }
	// return storeCategory, errors.New("the store category couldn't be found")
}

func (service *storeCategoryService) EditStoreCategory(storeCategoryEditInfo entity.StoreCategoryEditRequest) error {
	_, err := service.driver.EditStoreCategory(storeCategoryEditInfo)
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
	// 			storeCategories[i].Name = storeCategoryEditInfo.Name
	// 			storeCategories[i].Active = storeCategoryEditInfo.Active
	// 			return nil
	// 		}
	// 	} else {
	// 		return errors.New("store category id cannot be zero")
	// 	}
	// }
	// return errors.New("the store category couldn't be found")
}

func (service *storeCategoryService) ActivateStoreCategory(storeCategoryEditInfo entity.StoreCategoryInfoRequest) error {
	err := service.driver.ActivateStoreCategory(storeCategoryEditInfo)
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

func (service *storeCategoryService) DeactivateStoreCategory(storeCategoryEditInfo entity.StoreCategoryInfoRequest) error {
	err := service.driver.DeactivateStoreCategory(storeCategoryEditInfo)
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

func (service *storeCategoryService) DeleteStoreCategory(storeCategoryDeleteInfo entity.StoreCategoryDeleteRequest) error {
	err := service.driver.DeleteStoreCategory(storeCategoryDeleteInfo.ID)
	if err != nil {
		return err
	}
	return nil
	// storeCategories := service.FindAllStoreCategories()
	// var tempStoreCategory []entity.StoreCategory
	// for i := 0; i < len(storeCategories) && len(storeCategories) != 0; i++ {
	// 	if storeCategoryDeleteInfo.ID != 0 {
	// 		if storeCategories[i].ID != storeCategoryDeleteInfo.ID {
	// 			tempStoreCategory = append(tempStoreCategory, storeCategories[i])
	// 		}
	// 	} else {
	// 		return errors.New("id cannot be zero")
	// 	}
	// }
	// if len(storeCategories) != (len(tempStoreCategory) + 1) {
	// 	return errors.New("category could not be found")
	// }
	// service.storeCategories = tempStoreCategory
	// return nil
}

func (service *storeCategoryService) AddMockStoreCategories(storeCategories []entity.StoreCategory) {
	service.storeCategories = append(service.storeCategories, storeCategories...)
}
