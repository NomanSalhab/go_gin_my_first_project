package service

import (
	"errors"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

type StoreCategoryService interface {
	AddStoreCategory(entity.StoreCategory) error
	FindAllStoreCategories() []entity.StoreCategory
	FindActiveStoreCategories() []entity.StoreCategory
	FindNotActiveStoreCategories() []entity.StoreCategory
	FindStoreCategory(id entity.StoreCategoryInfoRequest) (entity.StoreCategory, error)
	EditStoreCategory(storeCategoryEditInfo entity.StoreCategoryEditRequest) error
	DeleteStoreCategory(user entity.StoreCategoryDeleteRequest) error
}

type storeCategoryService struct {
	storeCategories []entity.StoreCategory
}

func NewStoreCategoryService() StoreCategoryService {
	return &storeCategoryService{}
}

func (service *storeCategoryService) AddStoreCategory(storeCategory entity.StoreCategory) error {
	successStoreCategory := storeCategory
	for i := 0; i < len(service.storeCategories); i++ {
		if service.storeCategories[i].Name == storeCategory.Name {
			return errors.New("store category already exisists")
		}
	}
	if len(service.storeCategories) > 0 {
		successStoreCategory.ID = service.storeCategories[len(service.storeCategories)-1].ID + 1
	} else {
		successStoreCategory.ID = 1
	}

	if !storeCategory.Active {
		storeCategory.Active = false
	}
	service.storeCategories = append(service.storeCategories, successStoreCategory)
	return nil
}

func (service *storeCategoryService) FindAllStoreCategories() []entity.StoreCategory {
	return service.storeCategories
}

func (service *storeCategoryService) FindActiveStoreCategories() []entity.StoreCategory {
	var activeStoreCategories []entity.StoreCategory
	for i := 0; i < len(service.storeCategories); i++ {
		if service.storeCategories[i].Active {
			activeStoreCategories = append(activeStoreCategories, service.storeCategories[i])
		}
	}
	return activeStoreCategories
}

func (service *storeCategoryService) FindNotActiveStoreCategories() []entity.StoreCategory {
	var notActiveStoreCategories []entity.StoreCategory
	for i := 0; i < len(service.storeCategories); i++ {
		if !service.storeCategories[i].Active {
			notActiveStoreCategories = append(notActiveStoreCategories, service.storeCategories[i])
		}
	}
	return notActiveStoreCategories
}

func (service *storeCategoryService) FindStoreCategory(id entity.StoreCategoryInfoRequest) (entity.StoreCategory, error) {
	storeCategories := service.FindAllStoreCategories()
	var storeCategory entity.StoreCategory
	for i := 0; i < len(storeCategories) && len(storeCategories) != 0; i++ {
		if id.ID != 0 {
			if storeCategories[i].ID == id.ID {
				storeCategory = storeCategories[i]
			}
		} else {
			return storeCategory, errors.New("store category id cannot be zero")
		}
	}
	if id.ID == 0 {
		return storeCategory, errors.New("store category id cannot be zero")
	}
	if storeCategory.Name == "" {
		return storeCategory, errors.New("the store category couldn't be found")
	}
	return storeCategory, nil
}

func (service *storeCategoryService) EditStoreCategory(id entity.StoreCategoryEditRequest) error {
	storeCategories := service.FindAllStoreCategories()
	var storeCategory entity.StoreCategory
	for i := 0; i < len(storeCategories) && len(storeCategories) != 0; i++ {
		if id.ID != 0 {
			if storeCategories[i].ID == id.ID {
				storeCategory = entity.StoreCategory{
					ID:     id.ID,
					Name:   id.Name,
					Active: id.Active,
				}
				if id.Name != "" {
					storeCategories[i].Name = id.Name
				}
				if id.Active {
					storeCategories[i].Active = id.Active
				}
			}
		} else {
			return errors.New("store category id cannot be zero")
		}
	}
	if id.ID == 0 {
		return errors.New("store category id cannot be zero")
	}
	if storeCategory.ID == 0 {
		return errors.New("the store category couldn't be found")
	}
	return nil
}

func (service *storeCategoryService) DeleteStoreCategory(user entity.StoreCategoryDeleteRequest) error {
	storeCategories := service.FindAllStoreCategories()
	var tempStoreCategory []entity.StoreCategory
	if user.ID == 0 {
		return errors.New("user id cannot be zero")
	}
	for i := 0; i < len(storeCategories) && len(storeCategories) != 0; i++ {
		if user.ID != 0 {
			if storeCategories[i].ID != user.ID {
				tempStoreCategory = append(tempStoreCategory, storeCategories[i])
			}
		}
	}
	if len(storeCategories) != len(tempStoreCategory)+1 {
		return errors.New("user could not be found")
	}
	service.storeCategories = tempStoreCategory
	return nil
}
