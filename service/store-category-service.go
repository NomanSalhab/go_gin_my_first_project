package service

import "github.com/NomanSalhab/go_gin_my_first_project/entity"

type StoreCategoryService interface {
	SaveStoreCategory(entity.StoreCategory) entity.StoreCategory
	FindAllStoreCategories() []entity.StoreCategory
}

type storeCategoryService struct {
	storeCategories []entity.StoreCategory
}

func NewStoreCategoryService() StoreCategoryService {
	return &storeCategoryService{}
}

func (service *storeCategoryService) SaveStoreCategory(storeCategory entity.StoreCategory) entity.StoreCategory {
	service.storeCategories = append(service.storeCategories, storeCategory)
	return storeCategory
}

func (service *storeCategoryService) FindAllStoreCategories() []entity.StoreCategory {
	return service.storeCategories
}
