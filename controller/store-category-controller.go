package controller

import (
	"github.com/NomanSalhab/go_gin_my_first_project/entity"
	"github.com/NomanSalhab/go_gin_my_first_project/service"
	"github.com/gin-gonic/gin"
)

type StoreCategoryController interface {
	FindAllStoreCategories() []entity.StoreCategory
	AddStoreCategory(ctx *gin.Context) error
	FindActiveStoreCategories() []entity.StoreCategory
	FindNotActiveStoreCategories() []entity.StoreCategory
	GetStoreCategoryById(ctx *gin.Context) (entity.StoreCategory, error)
	EditStoreCategory(ctx *gin.Context) error
	DeleteStoreCategory(ctx *gin.Context) error
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

func (c *storeCategoryController) AddStoreCategory(ctx *gin.Context) error /*entity.Video*/ {
	var storeCategory entity.StoreCategory
	err := ctx.ShouldBindJSON(&storeCategory)
	if err != nil {
		return err
	}

	err = validate.Struct(storeCategory)
	if err != nil {
		return err
	}

	err = c.service.AddStoreCategory(storeCategory)
	return err /*video*/
}

func (c *storeCategoryController) FindActiveStoreCategories() []entity.StoreCategory {
	return c.service.FindActiveStoreCategories()
}

func (c *storeCategoryController) FindNotActiveStoreCategories() []entity.StoreCategory {
	return c.service.FindNotActiveStoreCategories()
}

func (c *storeCategoryController) GetStoreCategoryById(ctx *gin.Context) (entity.StoreCategory, error) {
	var storeCategoryId entity.StoreCategoryInfoRequest
	var storeCategory entity.StoreCategory
	err := ctx.ShouldBindJSON(&storeCategoryId)
	if err != nil {
		return storeCategory, err
	}
	storeCategory, err = c.service.FindStoreCategory(storeCategoryId)
	if err != nil {
		return storeCategory, err
	}
	return storeCategory, nil
}

func (c *storeCategoryController) EditStoreCategory(ctx *gin.Context) error {
	var storeCategoryEditInfo entity.StoreCategoryEditRequest
	err := ctx.ShouldBindJSON(&storeCategoryEditInfo)
	if err != nil {
		return err
	}
	err = c.service.EditStoreCategory(storeCategoryEditInfo)
	if err != nil {
		return err
	}
	return nil
}

func (c *storeCategoryController) DeleteStoreCategory(ctx *gin.Context) error {
	var storeCategoryId entity.StoreCategoryDeleteRequest
	err := ctx.ShouldBindJSON(&storeCategoryId)
	if err != nil {
		return err
	}
	err = c.service.DeleteStoreCategory(storeCategoryId)
	if err != nil {
		return err
	}
	return nil
}
