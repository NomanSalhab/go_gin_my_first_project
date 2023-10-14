package service

import (
	"github.com/NomanSalhab/go_gin_my_first_project/driver"
	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

type SliderService interface {
	AddSlider(slider entity.Slider) error
	FindAllSliders() []entity.Slider
	FindActiveSliders() []entity.Slider
	FindNotActiveSliders() []entity.Slider
	EditSlider(sliderEditInfo entity.SliderEditRequest) error
	FindSlidersByStore(sliderId entity.StoreSliders) ([]entity.Slider, error)
	DeleteSlider(sliderDeleteInfo entity.SliderEditRequest) error
}

type sliderService struct {
	driver driver.SliderDriver
}

func NewSliderService(driver driver.SliderDriver) SliderService {
	return &sliderService{
		driver: driver,
	}
}

func (service *sliderService) AddSlider(slider entity.Slider) error {
	err := service.driver.AddSlider(slider)
	if err != nil {
		return err
	}
	return nil
	// successSlider := slider
	// if len(service.sliders) > 0 {
	// 	successSlider.ID = service.sliders[len(service.sliders)-1].ID + 1
	// } else {
	// 	successSlider.ID = 1
	// }
	// service.sliders = append(service.sliders, successSlider)
	// return nil
}

func (service *sliderService) FindAllSliders() []entity.Slider {
	allSliders, err := service.driver.FindAllSliders()
	if err != nil {
		return make([]entity.Slider, 0)
	}
	return allSliders
	// return service.sliders
}

func (service *sliderService) FindSlidersByStore(sliderID entity.StoreSliders) ([]entity.Slider, error) {
	sliders, err := service.driver.FindSlidersByStore(sliderID.StoreId)
	if err != nil {
		return make([]entity.Slider, 0), err
	}
	return sliders, nil
	// sliders := service.sliders
	// var storeSliders []entity.Slider
	// if sliderID.StoreId != 0 {
	// 	for i := 0; i < len(sliders) && len(sliders) != 0; i++ {
	// 		if sliders[i].StoreId == sliderID.StoreId {
	// 			storeSliders = append(storeSliders, sliders[i])
	// 		}
	// 	}
	// } else {
	// 	return sliders, errors.New("store id cannot be zero")
	// }
	// return storeSliders, nil
}

func (service *sliderService) FindActiveSliders() []entity.Slider {
	activeSliders, err := service.driver.FindActiveSliders()
	if err != nil {
		return make([]entity.Slider, 0)
	}
	return activeSliders
}

func (service *sliderService) FindNotActiveSliders() []entity.Slider {
	notActiveSliders, err := service.driver.FindNotActiveSliders()
	if err != nil {
		return make([]entity.Slider, 0)
	}
	return notActiveSliders
	// var notActiveSliders []entity.Slider
	// for i := 0; i < len(service.sliders); i++ {
	// 	if !service.sliders[i].Active {
	// 		notActiveSliders = append(notActiveSliders, service.sliders[i])
	// 	}
	// }
	// return notActiveSliders
}

// func (service *sliderService) FindSlider(id entity.SliderEditRequest) (entity.Slider, error) {
// 	sliders := service.FindAllSliders()
// 	var slider entity.Slider
// 	for i := 0; i < len(sliders) && len(sliders) != 0; i++ {
// 		if id.ID != 0 {
// 			if sliders[i].ID == id.ID {
// 				slider = sliders[i]
// 				return slider, nil
// 			}
// 		} else {
// 			return slider, errors.New("store category id cannot be zero")
// 		}
// 	}
// 	return slider, errors.New("the store category couldn't be found")
// }

func (service *sliderService) EditSlider(sliderEditInfo entity.SliderEditRequest) error {
	_, err := service.driver.EditSlider(sliderEditInfo)
	if err != nil {
		return err
	}
	return nil
	// sliders := service.FindAllSliders()
	// var slider entity.Slider
	// for i := 0; i < len(sliders) && len(sliders) != 0; i++ {
	// 	if sliderEditInfo.ID != 0 {
	// 		if sliders[i].ID == sliderEditInfo.ID {
	// 			slider.ID = sliderEditInfo.ID
	// 			sliders[i].Active = sliderEditInfo.Active
	// 			sliders[i].StoreId = sliderEditInfo.StoreId
	// 			return nil
	// 		}
	// 	} else {
	// 		return errors.New("store category id cannot be zero")
	// 	}
	// }
	// return errors.New("the store category couldn't be found")
}

// func (service *sliderService) ActivateSlider(sliderEditInfo entity.SliderEditRequest) error {
// 	sliders := service.FindAllSliders()
// 	var slider entity.Slider
// 	for i := 0; i < len(sliders) && len(sliders) != 0; i++ {
// 		if sliderEditInfo.ID != 0 {
// 			if sliders[i].ID == sliderEditInfo.ID {
// 				slider.ID = sliderEditInfo.ID
// 				sliders[i].Active = true
// 				return nil
// 			}
// 		} else {
// 			return errors.New("store category id cannot be zero")
// 		}
// 	}
// 	return errors.New("the store category couldn't be found")
// }

// func (service *sliderService) DeactivateSlider(sliderEditInfo entity.SliderEditRequest) error {
// 	sliders := service.FindAllSliders()
// 	var slider entity.Slider
// 	for i := 0; i < len(sliders) && len(sliders) != 0; i++ {
// 		if sliderEditInfo.ID != 0 {
// 			if sliders[i].ID == sliderEditInfo.ID {
// 				slider.ID = sliderEditInfo.ID
// 				sliders[i].Active = false
// 				return nil
// 			}
// 		} else {
// 			return errors.New("store category id cannot be zero")
// 		}
// 	}
// 	return errors.New("the store category couldn't be found")
// }

func (service *sliderService) DeleteSlider(sliderDeleteInfo entity.SliderEditRequest) error {
	err := service.driver.DeleteSlider(sliderDeleteInfo.ID)
	if err != nil {
		return err
	}
	return nil
	// sliders := service.FindAllSliders()
	// var tempSlider []entity.Slider
	// for i := 0; i < len(sliders) && len(sliders) != 0; i++ {
	// 	if sliderDeleteInfo.ID != 0 {
	// 		if sliders[i].ID != sliderDeleteInfo.ID {
	// 			tempSlider = append(tempSlider, sliders[i])
	// 		}
	// 	} else {
	// 		return errors.New("id cannot be zero")
	// 	}
	// }
	// if len(sliders) != (len(tempSlider) + 1) {
	// 	return errors.New("slider could not be found")
	// }
	// service.sliders = tempSlider
	// return nil
}
