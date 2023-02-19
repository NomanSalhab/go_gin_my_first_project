package controller

import (
	"github.com/NomanSalhab/go_gin_my_first_project/entity"
	"github.com/NomanSalhab/go_gin_my_first_project/service"
	"github.com/gin-gonic/gin"
)

type StoreCategoryController interface {
	FindAllStoreCategories() []entity.StoreCategory
	SaveStoreCategory(ctx *gin.Context) error
}

type storeCategoryController struct {
	service service.StoreCategoryService
}

func NewStoreCategoryController(service service.StoreCategoryService) StoreCategoryController {
	/*validate = validator.New()
	validate.RegisterValidation("is-full-name", validators.ValidateFullUserName)*/
	return &storeCategoryController{
		service: service,
	}
}

func (c *storeCategoryController) FindAllStoreCategories() []entity.StoreCategory {
	return c.service.FindAllStoreCategories()
}

func (c *storeCategoryController) SaveStoreCategory(ctx *gin.Context) error /*entity.Video*/ {
	var storeCategory entity.StoreCategory
	err := ctx.ShouldBindJSON(&storeCategory)
	if err != nil {
		return err
	}

	err = validate.Struct(storeCategory)
	if err != nil {
		return err
	}

	c.service.SaveStoreCategory(storeCategory)
	return nil /*video*/
}
