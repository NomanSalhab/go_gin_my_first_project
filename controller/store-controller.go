package controller

import (
	"errors"
	"fmt"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
	"github.com/NomanSalhab/go_gin_my_first_project/service"
	"github.com/gin-gonic/gin"
)

type StoreController interface {
	FindAllStores() []entity.Store
	AddStore(ctx *gin.Context, cst StoreCategoryController) error
	FindActiveStores(ctx *gin.Context) []entity.Store
	FindNotActiveStores() []entity.Store
	GetStoreById(ctx *gin.Context) (entity.Store, error)
	EditStore(ctx *gin.Context, cst StoreCategoryController) error
	ActivateStore(ctx *gin.Context) error
	DeactivateStore(ctx *gin.Context) error
	DeleteStore(ctx *gin.Context) error
}

type storeController struct {
	service service.StoreService
}

func NewStoreController(service service.StoreService) StoreController {
	/*validate = validator.New()
	validate.RegisterValidation("is-full-name", validators.ValidateFullUserName)*/
	return &storeController{
		service: service,
	}
}

func (c *storeController) FindAllStores() []entity.Store {
	return c.service.FindAllStores()
}

func (c *storeController) AddStore(ctx *gin.Context, cst StoreCategoryController) error {
	var store entity.Store
	err := ctx.ShouldBindJSON(&store)
	if err != nil {
		return err
	}

	err = validate.Struct(store)
	if err != nil {
		return err
	}

	storeCategories := cst.FindAllStoreCategories()
	fmt.Println("Store Category Id:", store.StoreCategoryId, "Store Categories Length:", len(storeCategories))
	for i := 0; i < len(storeCategories); i++ {
		// fmt.Println(storeCategories[i].ID, store.StoreCategoryId)
		if storeCategories[i].ID == store.StoreCategoryId {
			err = c.service.AddStore(store)
			return err
		}
	}
	return errors.New("store category does not exist")
}

func (c *storeController) FindActiveStores(ctx *gin.Context) []entity.Store {
	var govId entity.AreaEditRequest
	err := ctx.ShouldBindJSON(&govId)
	if err != nil {
		return make([]entity.Store, 0)
	}
	return c.service.FindActiveStores(govId.ID)
}

func (c *storeController) FindNotActiveStores() []entity.Store {
	return c.service.FindNotActiveStores()
}

func (c *storeController) GetStoreById(ctx *gin.Context) (entity.Store, error) {
	var storeId entity.StoreInfoRequest
	var store entity.Store
	err := ctx.ShouldBindJSON(&storeId)
	if err != nil {
		return store, err
	}
	store, err = c.service.FindStore(storeId)
	if err != nil {
		return store, err
	}
	return store, nil
}

func (c *storeController) EditStore(ctx *gin.Context, cst StoreCategoryController) error {
	var storeEditInfo entity.StoreEditRequest
	err := ctx.ShouldBindJSON(&storeEditInfo)
	if err != nil {
		return err
	}
	err = validate.Struct(storeEditInfo)
	if err != nil {
		return err
	}

	storeCategories := cst.FindAllStoreCategories()
	for i := 0; i < len(storeCategories); i++ {
		if storeCategories[i].ID == storeEditInfo.StoreCategoryId {
			err = c.service.EditStore(storeEditInfo)
			return err
		}
	}
	return errors.New("store category does not exist")
}

func (c *storeController) ActivateStore(ctx *gin.Context) error {
	var storeEditInfo entity.StoreInfoRequest
	err := ctx.ShouldBindJSON(&storeEditInfo)
	if err != nil {
		return err
	}
	err = validate.Struct(storeEditInfo)
	if err != nil {
		return err
	}

	err = c.service.ActivateStore(storeEditInfo)
	return err
}

func (c *storeController) DeactivateStore(ctx *gin.Context) error {
	var storeEditInfo entity.StoreInfoRequest
	err := ctx.ShouldBindJSON(&storeEditInfo)
	if err != nil {
		return err
	}
	err = validate.Struct(storeEditInfo)
	if err != nil {
		return err
	}

	err = c.service.DeactivateStore(storeEditInfo)
	return err
}

func (c *storeController) DeleteStore(ctx *gin.Context) error {
	var storeId entity.StoreDeleteRequest
	err := ctx.ShouldBindJSON(&storeId)
	if err != nil {
		return err
	}
	err = c.service.DeleteStore(storeId)
	if err != nil {
		return err
	}
	return nil
}
