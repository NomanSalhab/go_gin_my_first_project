package service

import (
	"errors"

	"github.com/NomanSalhab/go_gin_my_first_project/driver"
	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

type AreaService interface {
	AddArea(entity.Area) error
	EditArea(areaEditInfo entity.AreaEditRequest) error
	ActivateArea(areaEditInfo entity.AreaActivateRequest) error
	DeactivateArea(areaEditInfo entity.AreaDeactivateRequest) error
	DeleteArea(areaDeleteInfo entity.AreaEditRequest) error

	FindAllAreas() []entity.Area
	FindActiveAreas() []entity.Area
	FindNotActiveAreas() []entity.Area
}

type areaService struct {
	driver driver.AreaDriver
}

func NewAreaService(driver driver.AreaDriver) AreaService {
	return &areaService{
		driver: driver}
}

func (service *areaService) AddArea(area entity.Area) error {
	areasList, err := service.driver.FindAllAreas()
	if err != nil {
		return err
	}
	for i := 0; i < len(areasList); i++ {
		if areasList[i].Name == area.Name {
			return errors.New("area name already exists")
		}
	}
	err = service.driver.AddArea(area)
	if err != nil {
		return err
	}
	return nil
}

func (service *areaService) EditArea(areaEditInfo entity.AreaEditRequest) error {
	_, err := service.driver.EditArea(areaEditInfo)
	if err != nil {
		return err
	}
	return nil
}

func (service *areaService) ActivateArea(areaEditInfo entity.AreaActivateRequest) error {
	err := service.driver.ActivateArea(areaEditInfo)
	if err != nil {
		return err
	}
	return nil
}

func (service *areaService) DeactivateArea(areaEditInfo entity.AreaDeactivateRequest) error {
	err := service.driver.DeactivateArea(areaEditInfo)
	if err != nil {
		return err
	}
	return nil
}

func (service *areaService) DeleteArea(areaDeleteInfo entity.AreaEditRequest) error {
	err := service.driver.DeleteArea(areaDeleteInfo.ID)
	if err != nil {
		return err
	}
	return nil
}

func (service *areaService) FindAllAreas() []entity.Area {
	areas, err := service.driver.FindAllAreas()
	if err != nil {
		return make([]entity.Area, 0)
	}
	return areas
}

func (service *areaService) FindActiveAreas() []entity.Area {
	activeAreas, err := service.driver.FindActiveAreas()
	if err != nil {
		return make([]entity.Area, 0)
	}
	return activeAreas
}

func (service *areaService) FindNotActiveAreas() []entity.Area {
	notActiveAreas, err := service.driver.FindNotActiveAreas()
	if err != nil {
		return make([]entity.Area, 0)
	}
	return notActiveAreas
}
