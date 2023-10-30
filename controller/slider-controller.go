package controller

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
	"github.com/NomanSalhab/go_gin_my_first_project/service"
	"github.com/gin-gonic/gin"
)

type SliderController interface {
	AddSlider(ctx *gin.Context, fc FileController) error
	FindAllSliders() []entity.Slider
	FindActiveSliders() []entity.Slider
	FindNotActiveSliders() []entity.Slider
	EditSlider(ctx *gin.Context) error
	DeleteSlider(ctx *gin.Context, fc FileController) error
	GetSlidersByStore(ctx *gin.Context, sc StoreController) ([]entity.Slider, error)
}

// var (
// 	fs FileService    = NewFileService()
// 	fc FileController = NewFileController(fs)
// )

type sliderController struct {
	service service.SliderService
}

func NewSliderController(service service.SliderService) SliderController {
	/*validate = validator.New()
	validate.RegisterValidation("is-full-name", validators.ValidateFullUserName)*/
	return &sliderController{
		service: service,
	}
}

func (c *sliderController) FindAllSliders() []entity.Slider {
	return c.service.FindAllSliders()
}

func (c *sliderController) AddSlider(ctx *gin.Context, fc FileController) error /*entity.Video*/ {
	var slider entity.Slider

	// err = ctx.ShouldBind(&slider)
	// if err != nil {
	// 	return err
	// }
	// err = validate.Struct(slider)
	// if err != nil {
	// 	return err
	// }
	// slider.Image = fileMetadata

	slider, err := SetSlider(ctx, fc)
	if err != nil {
		return err
	}

	err = c.service.AddSlider(slider)
	return err /*video*/
}

func (c *sliderController) FindActiveSliders() []entity.Slider {
	return c.service.FindActiveSliders()
}

func (c *sliderController) FindNotActiveSliders() []entity.Slider {
	return c.service.FindNotActiveSliders()
}

// func (c *sliderController) GetSliderById(ctx *gin.Context) (entity.Slider, error) {
// 	var sliderId entity.SliderEditRequest
// 	var slider entity.Slider
// 	err := ctx.ShouldBindJSON(&sliderId)
// 	if err != nil {
// 		return slider, err
// 	}
// 	slider, err = c.service.FindSlider(sliderId)
// 	if err != nil {
// 		return slider, err
// 	}
// 	return slider, nil
// }

func (c *sliderController) EditSlider(ctx *gin.Context) error {
	var sliderEditInfo entity.SliderEditRequest
	_, err1 := GetImage(ctx)
	if err1 != nil {
		id, idBool := ctx.GetPostForm("id")
		if !idBool || id == "0" {
			return errors.New("id is required")
		}
		idValue, idError := strconv.Atoi(id)
		if idError != nil {
			return errors.New("an unacceptable store id value")
		}
		storeId, storeIdBool := ctx.GetPostForm("store_id")
		if !storeIdBool {
			storeId = "0"
			// return errors.New("store id is required")
		}
		storeIdValue, storeIdError := strconv.Atoi(storeId)
		if storeIdError != nil {
			return errors.New("an unacceptable store id value")
		}
		productId, productIdBool := ctx.GetPostForm("product_id")
		if !productIdBool {
			productId = "0"
			// return errors.New("product id is required")
		}
		productIdValue, productIdError := strconv.Atoi(productId)
		if productIdError != nil {
			return errors.New("an unacceptable product id value")
		}
		active, _ := ctx.GetPostForm("active")
		var activeValue bool
		if active == "true" {
			activeValue = true
		} else {
			activeValue = false
		}
		sliderEditInfo = entity.SliderEditRequest{
			ID:        idValue,
			StoreId:   storeIdValue,
			ProductId: productIdValue,
			Active:    activeValue,
		}
		// return nil
	} /* else {
		err := ctx.ShouldBindJSON(&sliderEditInfo)
		if err != nil {
			return err
		}
		err = validate.Struct(sliderEditInfo)
		if err != nil {
			return err
		}
	}*/

	err := c.service.EditSlider(sliderEditInfo)
	if err != nil {
		return err
	}
	return nil
}

// func (c *sliderController) ActivateSlider(ctx *gin.Context) error {
// 	var sliderId entity.SliderEditRequest
// 	err := ctx.ShouldBindJSON(&sliderId)
// 	if err != nil {
// 		return err
// 	}
// 	err = validate.Struct(sliderId)
// 	if err != nil {
// 		return err
// 	}
// 	err = c.service.ActivateSlider(sliderId)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (c *sliderController) DeactivateSlider(ctx *gin.Context) error {
// 	var sliderId entity.SliderEditRequest
// 	err := ctx.ShouldBindJSON(&sliderId)
// 	if err != nil {
// 		return err
// 	}
// 	err = validate.Struct(sliderId)
// 	if err != nil {
// 		return err
// 	}
// 	err = c.service.DeactivateSlider(sliderId)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (c *sliderController) DeleteSlider(ctx *gin.Context, fc FileController) error {
	var sliderId entity.SliderEditRequest
	err := ctx.ShouldBindJSON(&sliderId)
	if err != nil {
		return err
	}
	allSilders := c.service.FindAllSliders()
	var oldSlider entity.Slider
	for i := 0; i < len(allSilders); i++ {
		if allSilders[i].ID == sliderId.ID {
			oldSlider = allSilders[i]
			break
		}
	}

	err = c.service.DeleteSlider(sliderId)
	if err != nil {
		return err
	}
	err = fc.DeleteFile(oldSlider.Image)
	if err != nil {
		return err
	}
	return nil
}

func (c *sliderController) GetSlidersByStore(ctx *gin.Context, sc StoreController) ([]entity.Slider, error) {
	var sliderStoreId entity.StoreSliders
	var storeSliders []entity.Slider
	err := ctx.ShouldBindJSON(&sliderStoreId)
	if err != nil {
		return storeSliders, err
	}

	err = validate.Struct(sliderStoreId)
	if err != nil {
		return storeSliders, err
	}

	stores := sc.FindAllStores()

	fmt.Println("Store Id:", sliderStoreId.StoreId, "Stores Length:", len(stores))
	for i := 0; i < len(stores); i++ {
		if stores[i].ID == sliderStoreId.StoreId {
			storeSliders, err = c.service.FindSlidersByStore(sliderStoreId)
			if err != nil {
				return storeSliders, err
			}
			if len(storeSliders) > 0 {
				return storeSliders, nil
			} else {
				return storeSliders, errors.New("store does not have any sliders")
			}
		}
	}

	if err != nil {
		return storeSliders, err
	}

	return storeSliders, errors.New("store does not exist")
}
func SetSlider(ctx *gin.Context, fc FileController) (entity.Slider, error) {
	fileMetadata, err := GetImage(ctx)
	if err != nil {
		return entity.Slider{}, err
	}

	storeId, storeIdBool := ctx.GetPostForm("store_id")
	if !storeIdBool {
		storeId = "0"
		// return errors.New("store id is required")
	}
	storeIdValue, storeIdError := strconv.Atoi(storeId)
	if storeIdError != nil {
		return entity.Slider{}, errors.New("an unacceptable store id value")
	}
	productId, productIdBool := ctx.GetPostForm("product_id")
	if !productIdBool && storeIdBool {
		productId = "0"
		// return errors.New("product id is required")
	}
	if !productIdBool && !storeIdBool {
		// productId = "0"
		return entity.Slider{}, errors.New("product id or store id is required")
	}
	productIdValue, productIdError := strconv.Atoi(productId)
	if productIdError != nil {
		return entity.Slider{}, errors.New("an unacceptable product id value")
	}
	active, _ := ctx.GetPostForm("active")
	var activeValue bool
	if active == "true" {
		activeValue = true
	} else {
		activeValue = false
	}
	err = fc.AddFile(fileMetadata)
	if err != nil {
		return entity.Slider{}, err
	}
	slider := entity.Slider{
		StoreId:   storeIdValue,
		ProductId: productIdValue,
		Active:    activeValue,
		Image:     fileMetadata.UUID,
	}
	return slider, nil
}

// func GetImage(ctx *gin.Context) (entity.File, error) {
// 	file, err := ctx.FormFile("image")
// 	if err != nil {
// 		return entity.File{}, err
// 	}
// 	filePath := filepath.Join("images", file.Filename)
// 	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
// 		// ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
// 		return entity.File{}, errors.New("failed to save file")
// 	}
// 	uuid := uuid.New().String()
// 	fileMetadata := entity.File{
// 		Filename: file.Filename,
// 		UUID:     uuid,
// 	}
// 	return fileMetadata, nil
// }
