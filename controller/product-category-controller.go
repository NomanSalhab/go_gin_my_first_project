package controller

import (
	"errors"
	"fmt"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
	"github.com/NomanSalhab/go_gin_my_first_project/service"
	"github.com/gin-gonic/gin"
)

type ProductCategoryController interface {
	FindAllProductCategories() []entity.ProductCategory
	AddProductCategory(ctx *gin.Context, cst StoreController) error
	FindActiveProductCategories() []entity.ProductCategory
	FindNotActiveProductCategories() []entity.ProductCategory
	GetProductCategoryById(ctx *gin.Context) (entity.ProductCategory, error)
	GetProductCategoriesByStore(ctx *gin.Context, sc StoreController) ([]entity.ProductCategory, error)
	EditProductCategory(ctx *gin.Context, cst StoreController) error
	ActivateProductCategory(ctx *gin.Context) error
	DeactivateProductCategory(ctx *gin.Context) error
	DeleteProductCategory(ctx *gin.Context) error
}

type productCategoryController struct {
	service service.ProductCategoryService
}

func NewProductCategoryController(service service.ProductCategoryService) ProductCategoryController {
	return &productCategoryController{
		service: service,
	}
}

func (c *productCategoryController) FindAllProductCategories() []entity.ProductCategory {
	return c.service.FindAllProductCategories()
}

func (c *productCategoryController) AddProductCategory(ctx *gin.Context, sc StoreController) error {
	var productCategory entity.ProductCategory
	err := ctx.ShouldBindJSON(&productCategory)
	if err != nil {
		return err
	}

	err = validate.Struct(productCategory)
	if err != nil {
		return err
	}

	stores := sc.FindAllStores()
	for i := 0; i < len(stores); i++ {
		if stores[i].ID == productCategory.StoreId {
			err = c.service.AddProductCategory(productCategory)
			return err
		}
	}
	return errors.New("store does not exist")
}

func (c *productCategoryController) FindActiveProductCategories() []entity.ProductCategory {
	return c.service.FindActiveProductCategories()
}

func (c *productCategoryController) FindNotActiveProductCategories() []entity.ProductCategory {
	return c.service.FindNotActiveProductCategories()
}

func (c *productCategoryController) GetProductCategoryById(ctx *gin.Context) (entity.ProductCategory, error) {
	var productCategoryId entity.ProductCategoryInfoRequest
	var productCategory entity.ProductCategory
	err := ctx.ShouldBindJSON(&productCategoryId)
	if err != nil {
		return productCategory, err
	}
	productCategory, err = c.service.FindProductCategory(productCategoryId)
	if err != nil {
		return productCategory, err
	}
	return productCategory, nil
}

func (c *productCategoryController) EditProductCategory(ctx *gin.Context, sc StoreController) error {
	var productCategoryEditInfo entity.ProductCategoryEditRequest
	err := ctx.ShouldBindJSON(&productCategoryEditInfo)
	if err != nil {
		return err
	}
	err = validate.Struct(productCategoryEditInfo)
	if err != nil {
		return err
	}

	// stores := sc.FindAllStores()
	// fmt.Println("Store Id:", productCategoryEditInfo.StoreId, "Stores Length:", len(stores))
	// for i := 0; i < len(stores); i++ {
	// if stores[i].ID == productCategoryEditInfo.StoreId {
	err = c.service.EditProductCategory(productCategoryEditInfo)
	return err
	// }
	// }
	// return errors.New("store does not exist")
}

func (c *productCategoryController) ActivateProductCategory(ctx *gin.Context) error {
	var productCategoryInfo entity.ProductCategoryInfoRequest
	err := ctx.ShouldBindJSON(&productCategoryInfo)
	if err != nil {
		return err
	}
	err = validate.Struct(productCategoryInfo)
	if err != nil {
		return err
	}

	err = c.service.ActivateProductCategory(productCategoryInfo)
	return err
}

func (c *productCategoryController) DeactivateProductCategory(ctx *gin.Context) error {
	var productCategoryInfo entity.ProductCategoryInfoRequest
	err := ctx.ShouldBindJSON(&productCategoryInfo)
	if err != nil {
		return err
	}
	err = validate.Struct(productCategoryInfo)
	if err != nil {
		return err
	}

	err = c.service.DeactivateProductCategory(productCategoryInfo)
	return err
}

func (c *productCategoryController) DeleteProductCategory(ctx *gin.Context) error {
	var productCategoryId entity.ProductCategoryDeleteRequest
	err := ctx.ShouldBindJSON(&productCategoryId)
	if err != nil {
		return err
	}
	err = c.service.DeleteProductCategory(productCategoryId)
	if err != nil {
		return err
	}
	return nil
}

func (c *productCategoryController) GetProductCategoriesByStore(ctx *gin.Context, sc StoreController) ([]entity.ProductCategory, error) {
	var productCategoryStoreId entity.ProductCategoriesByStoreInfoRequest
	var storeProductCategories []entity.ProductCategory
	err := ctx.ShouldBindJSON(&productCategoryStoreId)
	if err != nil {
		return storeProductCategories, err
	}

	err = validate.Struct(productCategoryStoreId)
	if err != nil {
		return storeProductCategories, err
	}

	stores := sc.FindAllStores()

	fmt.Println("Store Id:", productCategoryStoreId.StoreId, "Stores Length:", len(stores))
	for i := 0; i < len(stores); i++ {
		if stores[i].ID == productCategoryStoreId.StoreId {
			storeProductCategories, err = c.service.FindProductCategoryByStore(productCategoryStoreId)
			if err != nil {
				return storeProductCategories, err
			}
			if len(storeProductCategories) > 0 {
				return storeProductCategories, nil
			} else {
				return make([]entity.ProductCategory, 0), nil
				// return storeProductCategories, errors.New("store does not have any product categories")
			}
		}
	}

	if err != nil {
		return storeProductCategories, err
	}

	return storeProductCategories, errors.New("store does not exist")
}
