package service

import (
	"errors"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

type ProductCategoryService interface {
	AddProductCategory(entity.ProductCategory) error
	FindAllProductCategories() []entity.ProductCategory
	FindActiveProductCategories() []entity.ProductCategory
	FindNotActiveProductCategories() []entity.ProductCategory
	FindProductCategory(productCategoryId entity.ProductCategoryInfoRequest) (entity.ProductCategory, error)
	FindProductCategoryByStore(productCategoryId entity.ProductCategoriesByStoreInfoRequest) ([]entity.ProductCategory, error)
	EditProductCategory(productCategoryEditInfo entity.ProductCategoryEditRequest) error
	ActivateProductCategory(productCategoryEditInfo entity.ProductCategoryInfoRequest) error
	DeactivateProductCategory(productCategoryEditInfo entity.ProductCategoryInfoRequest) error
	DeleteProductCategory(productCategoryDeleteInfo entity.ProductCategoryDeleteRequest) error

	AddMockProductCategories(productCategories []entity.ProductCategory)
}

type productCategoryService struct {
	productCategories []entity.ProductCategory
}

func NewProductCategoryService() ProductCategoryService {
	return &productCategoryService{}
}

func (service *productCategoryService) AddProductCategory(productCategory entity.ProductCategory) error {
	successProductCategory := productCategory
	for i := 0; i < len(service.productCategories); i++ {
		if (service.productCategories[i].Name == productCategory.Name) && (service.productCategories[i].StoreId == productCategory.StoreId) {
			return errors.New("store category already exisists")
		}
	}
	if len(service.productCategories) > 0 {
		successProductCategory.ID = service.productCategories[len(service.productCategories)-1].ID + 1
	} else {
		successProductCategory.ID = 1
	}
	service.productCategories = append(service.productCategories, successProductCategory)
	return nil
}

func (service *productCategoryService) FindAllProductCategories() []entity.ProductCategory {
	return service.productCategories
}

func (service *productCategoryService) FindActiveProductCategories() []entity.ProductCategory {
	var activeProductCategoris []entity.ProductCategory
	for i := 0; i < len(service.productCategories); i++ {
		if service.productCategories[i].Active {
			activeProductCategoris = append(activeProductCategoris, service.productCategories[i])
		}
	}
	return activeProductCategoris
}

func (service *productCategoryService) FindNotActiveProductCategories() []entity.ProductCategory {
	var notActiveProductCategories []entity.ProductCategory
	for i := 0; i < len(service.productCategories); i++ {
		if !service.productCategories[i].Active {
			notActiveProductCategories = append(notActiveProductCategories, service.productCategories[i])
		}
	}
	return notActiveProductCategories
}

func (service *productCategoryService) FindProductCategory(id entity.ProductCategoryInfoRequest) (entity.ProductCategory, error) {
	productCategories := service.FindAllProductCategories()
	var productCategory entity.ProductCategory
	if id.ID != 0 {
		for i := 0; i < len(productCategories) && len(productCategories) != 0; i++ {
			if productCategories[i].ID == id.ID {
				productCategory = productCategories[i]
				return productCategory, nil
			}
		}
	} else {
		return productCategory, errors.New("store id cannot be zero")
	}
	return productCategory, errors.New("the store couldn't be found")
}

func (service *productCategoryService) FindProductCategoryByStore(productCategoryId entity.ProductCategoriesByStoreInfoRequest) ([]entity.ProductCategory, error) {
	productCategories := service.productCategories
	var storeProductCategories []entity.ProductCategory
	if productCategoryId.StoreId != 0 {
		for i := 0; i < len(productCategories) && len(productCategories) != 0; i++ {
			if productCategories[i].StoreId == productCategoryId.StoreId {
				storeProductCategories = append(storeProductCategories, productCategories[i])
			}
		}
	} else {
		return storeProductCategories, errors.New("store id cannot be zero")
	}
	return storeProductCategories, nil
}

func (service *productCategoryService) EditProductCategory(productCategoryEditInfo entity.ProductCategoryEditRequest) error {
	productCategories := service.FindAllProductCategories()
	if productCategoryEditInfo.ID != 0 {
		for i := 0; i < len(productCategories) && len(productCategories) != 0; i++ {
			if productCategories[i].ID == productCategoryEditInfo.ID {
				if productCategoryEditInfo.Name != "" {
					productCategories[i].Name = productCategoryEditInfo.Name
					productCategories[i].StoreId = productCategoryEditInfo.StoreId
					return nil
				}
			}
		}
	} else {
		return errors.New("product category id cannot be zero")
	}
	return errors.New("the product category couldn't be found")
}

func (service *productCategoryService) ActivateProductCategory(productCategoryEditInfo entity.ProductCategoryInfoRequest) error {
	productCategories := service.FindAllProductCategories()
	if productCategoryEditInfo.ID != 0 {
		for i := 0; i < len(productCategories) && len(productCategories) != 0; i++ {
			if productCategories[i].ID == productCategoryEditInfo.ID {
				productCategories[i].Active = true
				return nil
			}
		}
	} else {
		return errors.New("product category id cannot be zero")
	}
	return errors.New("the product category couldn't be found")
}

func (service *productCategoryService) DeactivateProductCategory(productCategoryEditInfo entity.ProductCategoryInfoRequest) error {
	productCategories := service.FindAllProductCategories()
	if productCategoryEditInfo.ID != 0 {
		for i := 0; i < len(productCategories) && len(productCategories) != 0; i++ {
			if productCategories[i].ID == productCategoryEditInfo.ID {
				productCategories[i].Active = false
				return nil
			}
		}
	} else {
		return errors.New("product category id cannot be zero")
	}
	return errors.New("the product category couldn't be found")
}

func (service *productCategoryService) DeleteProductCategory(productCategoryId entity.ProductCategoryDeleteRequest) error {
	productCategories := service.FindAllProductCategories()
	var tempProductCategory []entity.ProductCategory
	if productCategoryId.ID != 0 {
		for i := 0; i < len(productCategories) && len(productCategories) != 0; i++ {
			if productCategories[i].ID != productCategoryId.ID {
				tempProductCategory = append(tempProductCategory, productCategories[i])
			}
		}
	}
	if len(productCategories) != len(tempProductCategory)+1 {
		return errors.New("store could not be found")
	}
	service.productCategories = tempProductCategory
	return nil
}

func (service *productCategoryService) AddMockProductCategories(productCategories []entity.ProductCategory) {
	service.productCategories = append(service.productCategories, productCategories...)
}
