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
	DeactivateStoreCategory(storeCategoryEditInfo entity.StoreCategoryInfoRequest) error
	ActivateStoreCategory(storeCategoryEditInfo entity.StoreCategoryInfoRequest) error
	DeleteStoreCategory(user entity.StoreCategoryDeleteRequest) error

	AddMockStoreCategories(storeCategories []entity.StoreCategory)
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
				return storeCategory, nil
			}
		} else {
			return storeCategory, errors.New("store category id cannot be zero")
		}
	}
	return storeCategory, errors.New("the store category couldn't be found")
}

func (service *storeCategoryService) EditStoreCategory(storeCategoryEditInfo entity.StoreCategoryEditRequest) error {
	storeCategories := service.FindAllStoreCategories()
	var storeCategory entity.StoreCategory
	for i := 0; i < len(storeCategories) && len(storeCategories) != 0; i++ {
		if storeCategoryEditInfo.ID != 0 {
			if storeCategories[i].ID == storeCategoryEditInfo.ID {
				storeCategory.ID = storeCategoryEditInfo.ID
				storeCategories[i].Name = storeCategoryEditInfo.Name
				return nil
			}
		} else {
			return errors.New("store category id cannot be zero")
		}
	}
	return errors.New("the store category couldn't be found")
}

func (service *storeCategoryService) ActivateStoreCategory(storeCategoryEditInfo entity.StoreCategoryInfoRequest) error {
	storeCategories := service.FindAllStoreCategories()
	var storeCategory entity.StoreCategory
	for i := 0; i < len(storeCategories) && len(storeCategories) != 0; i++ {
		if storeCategoryEditInfo.ID != 0 {
			if storeCategories[i].ID == storeCategoryEditInfo.ID {
				storeCategory.ID = storeCategoryEditInfo.ID
				storeCategories[i].Active = true
				return nil
			}
		} else {
			return errors.New("store category id cannot be zero")
		}
	}
	return errors.New("the store category couldn't be found")
}

func (service *storeCategoryService) DeactivateStoreCategory(storeCategoryEditInfo entity.StoreCategoryInfoRequest) error {
	storeCategories := service.FindAllStoreCategories()
	var storeCategory entity.StoreCategory
	for i := 0; i < len(storeCategories) && len(storeCategories) != 0; i++ {
		if storeCategoryEditInfo.ID != 0 {
			if storeCategories[i].ID == storeCategoryEditInfo.ID {
				storeCategory.ID = storeCategoryEditInfo.ID
				storeCategories[i].Active = false
				return nil
			}
		} else {
			return errors.New("store category id cannot be zero")
		}
	}
	return errors.New("the store category couldn't be found")
}

func (service *storeCategoryService) DeleteStoreCategory(storeCategoryDeleteInfo entity.StoreCategoryDeleteRequest) error {
	storeCategories := service.FindAllStoreCategories()
	var tempStoreCategory []entity.StoreCategory
	for i := 0; i < len(storeCategories) && len(storeCategories) != 0; i++ {
		if storeCategoryDeleteInfo.ID != 0 {
			if storeCategories[i].ID != storeCategoryDeleteInfo.ID {
				tempStoreCategory = append(tempStoreCategory, storeCategories[i])
			}
		} else {
			return errors.New("id cannot be zero")
		}
	}
	if len(storeCategories) != (len(tempStoreCategory) + 1) {
		return errors.New("category could not be found")
	}
	service.storeCategories = tempStoreCategory
	return nil
}

func (service *storeCategoryService) AddMockStoreCategories(storeCategories []entity.StoreCategory) {
	service.storeCategories = append(service.storeCategories, storeCategories...)
}
