package service

import (
	"errors"

	"github.com/NomanSalhab/go_gin_my_first_project/driver"
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
	driver            driver.ProductCategoryDriver
}

func NewProductCategoryService(driver driver.ProductCategoryDriver) ProductCategoryService {
	return &productCategoryService{
		driver: driver,
	}
}

func (service *productCategoryService) AddProductCategory(productCategory entity.ProductCategory) error {
	productCategoriesList, err := service.driver.FindAllProductCategories()
	if err != nil {
		return err
	}
	for i := 0; i < len(productCategoriesList); i++ {
		if productCategoriesList[i].Name == productCategory.Name && productCategoriesList[i].StoreId == productCategory.StoreId {
			return errors.New("product category name with store_id already exists")
		}
	}
	err = service.driver.AddProductCategory(productCategory)
	if err != nil {
		return err
	}
	return nil
	// successProductCategory := productCategory
	// for i := 0; i < len(service.productCategories); i++ {
	// 	if (service.productCategories[i].Name == productCategory.Name) && (service.productCategories[i].StoreId == productCategory.StoreId) {
	// 		return errors.New("store category already exisists")
	// 	}
	// }
	// if len(service.productCategories) > 0 {
	// 	successProductCategory.ID = service.productCategories[len(service.productCategories)-1].ID + 1
	// } else {
	// 	successProductCategory.ID = 1
	// }
	// service.productCategories = append(service.productCategories, successProductCategory)
	// return nil
}

func (service *productCategoryService) FindAllProductCategories() []entity.ProductCategory {
	allProductCategories, err := service.driver.FindAllProductCategories()
	if err != nil {
		return make([]entity.ProductCategory, 0)
	}
	return allProductCategories
	// return service.productCategories
}

func (service *productCategoryService) FindActiveProductCategories() []entity.ProductCategory {
	activeProductCategories, err := service.driver.FindActiveProductCategories()
	if err != nil {
		return make([]entity.ProductCategory, 0)
	}
	return activeProductCategories
	// var activeProductCategoris []entity.ProductCategory
	// for i := 0; i < len(service.productCategories); i++ {
	// 	if service.productCategories[i].Active {
	// 		activeProductCategoris = append(activeProductCategoris, service.productCategories[i])
	// 	}
	// }
	// return activeProductCategoris
}

func (service *productCategoryService) FindNotActiveProductCategories() []entity.ProductCategory {
	notActiveProductCategories, err := service.driver.FindNotActiveProductCategories()
	if err != nil {
		return make([]entity.ProductCategory, 0)
	}
	return notActiveProductCategories
	// var notActiveProductCategories []entity.ProductCategory
	// for i := 0; i < len(service.productCategories); i++ {
	// 	if !service.productCategories[i].Active {
	// 		notActiveProductCategories = append(notActiveProductCategories, service.productCategories[i])
	// 	}
	// }
	// return notActiveProductCategories
}

func (service *productCategoryService) FindProductCategory(id entity.ProductCategoryInfoRequest) (entity.ProductCategory, error) {
	productCategory, _ := service.driver.FindProductCategory(id.ID)
	if productCategory.Name == "" {
		return productCategory, errors.New("the product category couldn't be found")
	}
	return productCategory, nil
	// productCategories := service.FindAllProductCategories()
	// var productCategory entity.ProductCategory
	// if id.ID != 0 {
	// 	for i := 0; i < len(productCategories) && len(productCategories) != 0; i++ {
	// 		if productCategories[i].ID == id.ID {
	// 			productCategory = productCategories[i]
	// 			return productCategory, nil
	// 		}
	// 	}
	// } else {
	// 	return productCategory, errors.New("store id cannot be zero")
	// }
	// return productCategory, errors.New("the store couldn't be found")
}

func (service *productCategoryService) FindProductCategoryByStore(productCategoryId entity.ProductCategoriesByStoreInfoRequest) ([]entity.ProductCategory, error) {
	storeProductCategories, err := service.driver.FindProductCategoryByStore(productCategoryId.StoreId)
	if err != nil {
		return make([]entity.ProductCategory, 0), err
	}
	return storeProductCategories, nil
	// productCategories := service.productCategories
	// var storeProductCategories []entity.ProductCategory
	// if productCategoryId.StoreId != 0 {
	// 	for i := 0; i < len(productCategories) && len(productCategories) != 0; i++ {
	// 		if productCategories[i].StoreId == productCategoryId.StoreId {
	// 			storeProductCategories = append(storeProductCategories, productCategories[i])
	// 		}
	// 	}
	// } else {
	// 	return storeProductCategories, errors.New("store id cannot be zero")
	// }
	// return storeProductCategories, nil
}

func (service *productCategoryService) EditProductCategory(productCategoryEditInfo entity.ProductCategoryEditRequest) error {
	_, err := service.driver.EditProductCategory(productCategoryEditInfo)
	if err != nil {
		return err
	}
	return nil
	// productCategories := service.FindAllProductCategories()
	// if productCategoryEditInfo.ID != 0 {
	// 	for i := 0; i < len(productCategories) && len(productCategories) != 0; i++ {
	// 		if productCategories[i].ID == productCategoryEditInfo.ID {
	// 			if productCategoryEditInfo.Name != "" {
	// 				productCategories[i].Name = productCategoryEditInfo.Name
	// 				if productCategoryEditInfo.StoreId != 0 {
	// 					productCategories[i].StoreId = productCategoryEditInfo.StoreId
	// 				}
	// 				productCategories[i].Active = productCategoryEditInfo.Active
	// 				return nil
	// 			}
	// 		}
	// 	}
	// } else {
	// 	return errors.New("product category id cannot be zero")
	// }
	// return errors.New("the product category couldn't be found")
}

func (service *productCategoryService) ActivateProductCategory(productCategoryEditInfo entity.ProductCategoryInfoRequest) error {
	err := service.driver.ActivateProductCategory(productCategoryEditInfo)
	if err != nil {
		return err
	}
	return nil
	// productCategories := service.FindAllProductCategories()
	// if productCategoryEditInfo.ID != 0 {
	// 	for i := 0; i < len(productCategories) && len(productCategories) != 0; i++ {
	// 		if productCategories[i].ID == productCategoryEditInfo.ID {
	// 			productCategories[i].Active = true
	// 			return nil
	// 		}
	// 	}
	// } else {
	// 	return errors.New("product category id cannot be zero")
	// }
	// return errors.New("the product category couldn't be found")
}

func (service *productCategoryService) DeactivateProductCategory(productCategoryEditInfo entity.ProductCategoryInfoRequest) error {
	err := service.driver.DeactivateProductCategory(productCategoryEditInfo)
	if err != nil {
		return err
	}
	return nil
	// productCategories := service.FindAllProductCategories()
	// if productCategoryEditInfo.ID != 0 {
	// 	for i := 0; i < len(productCategories) && len(productCategories) != 0; i++ {
	// 		if productCategories[i].ID == productCategoryEditInfo.ID {
	// 			productCategories[i].Active = false
	// 			return nil
	// 		}
	// 	}
	// } else {
	// 	return errors.New("product category id cannot be zero")
	// }
	// return errors.New("the product category couldn't be found")
}

func (service *productCategoryService) DeleteProductCategory(productCategoryId entity.ProductCategoryDeleteRequest) error {
	err := service.driver.DeleteProductCategory(productCategoryId.ID)
	if err != nil {
		return err
	}
	return nil
	// productCategories := service.FindAllProductCategories()
	// var tempProductCategory []entity.ProductCategory
	// if productCategoryId.ID != 0 {
	// 	for i := 0; i < len(productCategories) && len(productCategories) != 0; i++ {
	// 		if productCategories[i].ID != productCategoryId.ID {
	// 			tempProductCategory = append(tempProductCategory, productCategories[i])
	// 		}
	// 	}
	// }
	// if len(productCategories) != len(tempProductCategory)+1 {
	// 	return errors.New("store could not be found")
	// }
	// service.productCategories = tempProductCategory
	// return nil
}

func (service *productCategoryService) AddMockProductCategories(productCategories []entity.ProductCategory) {
	service.productCategories = append(service.productCategories, productCategories...)
}
