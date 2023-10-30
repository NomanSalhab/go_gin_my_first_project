package controller

import (
	"errors"
	"strconv"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
	"github.com/NomanSalhab/go_gin_my_first_project/service"
	"github.com/gin-gonic/gin"
)

type StoreController interface {
	FindAllStores() []entity.Store
	AddStore(ctx *gin.Context, cst StoreCategoryController, fc FileController) error
	FindActiveStores(ctx *gin.Context) []entity.Store
	FindNotActiveStores() []entity.Store
	GetStoreById(ctx *gin.Context) (entity.Store, error)
	EditStore(ctx *gin.Context, cst StoreCategoryController, fc FileController) error
	ActivateStore(ctx *gin.Context) error
	DeactivateStore(ctx *gin.Context) error
	DeleteStore(ctx *gin.Context, fc FileController) error
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

func (c *storeController) AddStore(ctx *gin.Context, cst StoreCategoryController, fc FileController) error {
	var store entity.Store
	// err := ctx.ShouldBindJSON(&store)
	// if err != nil {
	// 	return err
	// }
	// err = validate.Struct(store)
	// if err != nil {
	// 	return err
	// }

	// id, _ := ctx.GetPostForm("id")
	// idValue, idError := strconv.Atoi(id)
	// if idError != nil {
	// 	return errors.New("an unacceptable store id value")
	// }
	store, image, err := SetStore(ctx)
	if err != nil {
		return err
	}

	storeCategories := cst.FindAllStoreCategories()
	// fmt.Println("Store Category Id:", store.StoreCategoryId, "Store Categories Length:", len(storeCategories))
	for i := 0; i < len(storeCategories); i++ {
		// fmt.Println(storeCategories[i].ID, store.StoreCategoryId)
		if storeCategories[i].ID == store.StoreCategoryId {
			err = c.service.AddStore(store)
			if err != nil {
				return err
			}
			err = fc.AddFile(image)
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

func (c *storeController) EditStore(ctx *gin.Context, cst StoreCategoryController, fc FileController) error {
	var storeEditInfo entity.StoreEditRequest
	// err := ctx.ShouldBindJSON(&storeEditInfo)
	// if err != nil {
	// 	return err
	// }
	// err = validate.Struct(storeEditInfo)
	// if err != nil {
	// 	return err
	// }

	storeEditInfo, err := SetStoreForEditing(ctx)
	if err != nil {
		return err
	}

	storeCategories := cst.FindAllStoreCategories()
	for i := 0; i < len(storeCategories); i++ {
		if storeCategories[i].ID == storeEditInfo.StoreCategoryId {
			if storeEditInfo.Image != "" {
				oldStoreInfo, err := c.service.FindStore(entity.StoreInfoRequest{ID: storeEditInfo.ID})
				if err != nil {
					return err
				}
				err = fc.DeleteFile(oldStoreInfo.Image)
				if err != nil {
					return err
				}
				file, err := GetImage(ctx)
				if err != nil {
					return err
				}
				err = fc.AddFile(file)
				if err != nil {
					return err
				}
				storeEditInfo.Image = file.UUID
			}
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

func (c *storeController) DeleteStore(ctx *gin.Context, fc FileController) error {
	var storeId entity.StoreInfoRequest
	err := ctx.ShouldBindJSON(&storeId)
	if err != nil {
		return err
	}
	oldStoreInfo, err := c.service.FindStore(storeId)
	if err != nil {
		return err
	}
	err = fc.DeleteFile(oldStoreInfo.Image)
	if err != nil {
		return err
	}
	err = c.service.DeleteStore(storeId)
	if err != nil {
		return err
	}
	return nil
}

func SetStore(ctx *gin.Context) (entity.Store, entity.File, error) {
	areaId, areaIdBool := ctx.GetPostForm("area_id")
	if !areaIdBool {
		return entity.Store{}, entity.File{}, errors.New("area id is required")
	}
	areaIdValue, areaIdError := strconv.Atoi(areaId)
	if areaIdError != nil {
		return entity.Store{}, entity.File{}, errors.New("an unacceptable area id value")
	}
	balanceId, balanceBool := ctx.GetPostForm("balance")
	if !balanceBool {
		balanceId = "0"
	}
	balanceIdValue, balanceIdError := strconv.Atoi(balanceId)
	if balanceIdError != nil {
		return entity.Store{}, entity.File{}, errors.New("an unacceptable balance value")
	}
	deliveryRentId, deliveryRentIdBool := ctx.GetPostForm("delivery_rent")
	if !deliveryRentIdBool {
		return entity.Store{}, entity.File{}, errors.New("delivery rent is required")
	}
	deliveryRentIdValue, deliveryRentIdError := strconv.Atoi(deliveryRentId)
	if deliveryRentIdError != nil {
		return entity.Store{}, entity.File{}, errors.New("an unacceptable delivery rent value")
	}
	discount, discountBool := ctx.GetPostForm("discount")
	if !discountBool {
		discount = "0.0"
	}
	discount64Value, discountError := strconv.ParseFloat(discount, 32)
	if discountError != nil {
		return entity.Store{}, entity.File{}, errors.New("an unacceptable discount value")
	}
	discountValue := float32(discount64Value)
	active, _ := ctx.GetPostForm("active")
	var activeValue bool
	if active == "true" {
		activeValue = true
	} else {
		activeValue = false
	}
	name, _ := ctx.GetPostForm("name")
	if len(name) == 0 {
		return entity.Store{}, entity.File{}, errors.New("store name is required")
	}
	storeCategoryId, storeCategoryIdBool := ctx.GetPostForm("store_category_id")
	if !storeCategoryIdBool {
		return entity.Store{}, entity.File{}, errors.New("store category id is required")
	}
	storeCategoryIdValue, storeCategoryIdError := strconv.Atoi(storeCategoryId)
	if storeCategoryIdError != nil {
		return entity.Store{}, entity.File{}, errors.New("an unacceptable store categoty id value")
	}
	fileMetadata, err := GetImage(ctx)
	if err != nil {
		return entity.Store{}, entity.File{}, err
	}
	store := entity.Store{
		AreaID:          areaIdValue,
		Active:          activeValue,
		Name:            name,
		StoreCategoryId: storeCategoryIdValue,
		Image:           fileMetadata.UUID,
		Balance:         balanceIdValue,
		DeliveryRent:    deliveryRentIdValue,
		Discount:        discountValue,
	}
	return store, fileMetadata, nil
}

func SetStoreForEditing(ctx *gin.Context) (entity.StoreEditRequest, error) {
	id, idBool := ctx.GetPostForm("id")
	if !idBool {
		return entity.StoreEditRequest{}, errors.New("id is required")
	}
	idValue, idError := strconv.Atoi(id)
	if idError != nil {
		return entity.StoreEditRequest{}, errors.New("an unacceptable id value")
	}
	areaId, areaIdBool := ctx.GetPostForm("area_id")
	if !areaIdBool {
		areaId = "0"
	}
	areaIdValue, areaIdError := strconv.Atoi(areaId)
	if areaIdError != nil {
		return entity.StoreEditRequest{}, errors.New("an unacceptable area id value")
	}
	balanceId, balanceBool := ctx.GetPostForm("balance")
	if !balanceBool {
		balanceId = "0"
	}
	balanceIdValue, balanceIdError := strconv.Atoi(balanceId)
	if balanceIdError != nil {
		return entity.StoreEditRequest{}, errors.New("an unacceptable balance value")
	}
	deliveryRentId, deliveryRentIdBool := ctx.GetPostForm("delivery_rent")
	if !deliveryRentIdBool {
		deliveryRentId = "0"
	}
	deliveryRentIdValue, deliveryRentIdError := strconv.Atoi(deliveryRentId)
	if deliveryRentIdError != nil {
		return entity.StoreEditRequest{}, errors.New("an unacceptable delivery rent value")
	}
	discount, discountBool := ctx.GetPostForm("discount")
	if !discountBool {
		discount = "0.0"
	}
	discount64Value, discountError := strconv.ParseFloat(discount, 32)
	if discountError != nil {
		return entity.StoreEditRequest{}, errors.New("an unacceptable discount value")
	}
	discountValue := float32(discount64Value)
	active, _ := ctx.GetPostForm("active")
	var activeValue bool
	if active == "true" {
		activeValue = true
	} else {
		activeValue = false
	}
	name, _ := ctx.GetPostForm("name")
	if len(name) == 0 {
		name = ""
	}
	storeCategoryId, storeCategoryIdBool := ctx.GetPostForm("store_category_id")
	if !storeCategoryIdBool {
		storeCategoryId = "0"
	}
	storeCategoryIdValue, storeCategoryIdError := strconv.Atoi(storeCategoryId)
	if storeCategoryIdError != nil {
		return entity.StoreEditRequest{}, errors.New("an unacceptable store categoty id value")
	}
	var store entity.StoreEditRequest
	fileMetadata, err := GetImage(ctx)
	if err != nil {
		store = entity.StoreEditRequest{
			ID:              idValue,
			AreaID:          areaIdValue,
			Active:          activeValue,
			Name:            name,
			StoreCategoryId: storeCategoryIdValue,
			Balance:         balanceIdValue,
			DeliveryRent:    deliveryRentIdValue,
			Discount:        discountValue,
		}
	} else {
		store = entity.StoreEditRequest{
			ID:              idValue,
			AreaID:          areaIdValue,
			Active:          activeValue,
			Name:            name,
			StoreCategoryId: storeCategoryIdValue,
			Image:           fileMetadata.UUID,
			Balance:         balanceIdValue,
			DeliveryRent:    deliveryRentIdValue,
			Discount:        discountValue,
		}
	}
	return store, nil
}
