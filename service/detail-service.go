package service

import (
	"errors"

	"github.com/NomanSalhab/go_gin_my_first_project/driver"
	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

type DetailService interface {
	AddDetail(entity.Detail) error
	FindAllDetails() []entity.DetailEditRequest
	FindAllAddons() []entity.Detail
	FindAllFlavors() []entity.Detail
	FindAllVolumes() []entity.Detail
	EditDetail(detailEditInfo entity.DetailEditRequest) error
	DeleteDetail(detailDeleteInfo entity.DetailEditRequest) error

	// AddMockDetails(details []entity.Detail)
}

type detailService struct {
	driver driver.DetailDriver
}

func NewDetailService(driver driver.DetailDriver) DetailService {
	return &detailService{
		driver: driver,
	}
}

func (service *detailService) AddDetail(detail entity.Detail) error {
	detailsList, err := service.driver.FindAllDetails()
	if err != nil {
		return err
	}
	for i := 0; i < len(detailsList); i++ {
		if detailsList[i].Name == detail.Name && detailsList[i].Price == detail.Price {
			return errors.New("detail name already exists")
		}
	}
	if detail.IsAddon || detail.IsFlavor || detail.IsVolume {
		err = service.driver.AddDetail(detail)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("detail type should be specified")
	}

	// successDetail := detail
	// if len(service.details) > 0 {
	// 	successDetail.ID = service.details[len(service.details)-1].ID + 1
	// } else {
	// 	successDetail.ID = 1
	// }
	// service.details = append(service.details, successDetail)
	// return nil
}

func (service *detailService) FindAllDetails() []entity.DetailEditRequest {
	allDetails, err := service.driver.FindAllDetails()
	if err != nil {
		return make([]entity.DetailEditRequest, 0)
	}
	return allDetails
	// return service.details
}

func (service *detailService) FindAllAddons() []entity.Detail {
	allDetails, err := service.driver.FindAllAddons()
	if err != nil {
		return make([]entity.Detail, 0)
	}
	return allDetails
	// return service.details
}

func (service *detailService) FindAllFlavors() []entity.Detail {
	allDetails, err := service.driver.FindAllFlavors()
	if err != nil {
		return make([]entity.Detail, 0)
	}
	return allDetails
	// return service.details
}

func (service *detailService) FindAllVolumes() []entity.Detail {
	allDetails, err := service.driver.FindAllVolumes()
	if err != nil {
		return make([]entity.Detail, 0)
	}
	return allDetails
	// return service.details
}

// func (service *detailService) FindDetail(id entity.DetailEditRequest) (entity.Detail, error) {
// 	details := service.FindAllDetails()
// 	var detail entity.Detail
// 	for i := 0; i < len(details) && len(details) != 0; i++ {
// 		if id.ID != 0 {
// 			if details[i].ID == id.ID {
// 				detail = details[i]
// 				return detail, nil
// 			}
// 		} else {
// 			return detail, errors.New("store category id cannot be zero")
// 		}
// 	}
// 	return detail, errors.New("the detail couldn't be found")
// }

func (service *detailService) EditDetail(detailEditInfo entity.DetailEditRequest) error {
	_, err := service.driver.EditDetail(detailEditInfo)
	if err != nil {
		return err
	}
	return nil
	// details := service.FindAllDetails()
	// var detail entity.Detail
	// for i := 0; i < len(details) && len(details) != 0; i++ {
	// 	if detailEditInfo.ID != 0 {
	// 		if details[i].ID == detailEditInfo.ID {
	// 			detail.ID = detailEditInfo.ID
	// 			return nil
	// 		}
	// 	} else {
	// 		return errors.New("detail id cannot be zero")
	// 	}
	// 	if len(detailEditInfo.Name) != 0 {
	// 		if details[i].Name == detailEditInfo.Name {
	// 			detail.Name = detailEditInfo.Name
	// 			return nil
	// 		}
	// 	}
	// }
	// return errors.New("the detail couldn't be found")
}

func (service *detailService) DeleteDetail(detailDeleteInfo entity.DetailEditRequest) error {
	err := service.driver.DeleteDetail(detailDeleteInfo.ID)
	if err != nil {
		return err
	}
	return nil
	// details := service.FindAllDetails()
	// var tempDetail []entity.Detail
	// for i := 0; i < len(details) && len(details) != 0; i++ {
	// 	if detailDeleteInfo.ID != 0 {
	// 		if details[i].ID != detailDeleteInfo.ID {
	// 			tempDetail = append(tempDetail, details[i])
	// 		}
	// 	} else {
	// 		return errors.New("id cannot be zero")
	// 	}
	// }
	// if len(details) != (len(tempDetail) + 1) {
	// 	return errors.New("detail could not be found")
	// }
	// service.details = tempDetail
	// return nil
}

// func (service *detailService) AddMockDetails(details []entity.Detail) {
// 	service.details = append(service.details, details...)
// }
