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

type areaService struct{}

func NewAreaService() AreaService {
	return &areaService{}
}

func (service *areaService) AddArea(area entity.Area) error {
	areasList, err := driver.FindAllAreas()
	if err != nil {
		return err
	}
	for i := 0; i < len(areasList); i++ {
		if areasList[i].Name == area.Name {
			return errors.New("area name already exists")
		}
	}
	err = driver.AddArea(area)
	if err != nil {
		return err
	}
	return nil
}

// func (service *areaService) FindArea(id entity.AreaEditRequest) (entity.Area, error) {
// 	area, _ := driver.FindArea(id.ID)
// 	if area.Name == "" {
// 		return area, errors.New("the area couldn't be found")
// 	}
// 	return area, nil
// }

func (service *areaService) EditArea(areaEditInfo entity.AreaEditRequest) error {
	_, err := driver.EditArea(areaEditInfo)
	if err != nil {
		return err
	}
	return nil
}

func (service *areaService) ActivateArea(areaEditInfo entity.AreaActivateRequest) error {
	err := driver.ActivateArea(areaEditInfo)
	if err != nil {
		return err
	}
	return nil
}

func (service *areaService) DeactivateArea(areaEditInfo entity.AreaDeactivateRequest) error {
	err := driver.DeactivateArea(areaEditInfo)
	if err != nil {
		return err
	}
	return nil
}

func (service *areaService) DeleteArea(areaDeleteInfo entity.AreaEditRequest) error {
	err := driver.DeleteArea(areaDeleteInfo.ID)
	if err != nil {
		return err
	}
	return nil
}

func (service *areaService) FindAllAreas() []entity.Area {
	areas, err := driver.FindAllAreas()
	if err != nil {
		return make([]entity.Area, 0)
	}
	return areas
}

func (service *areaService) FindActiveAreas() []entity.Area {
	activeAreas, err := driver.FindActiveAreas()
	if err != nil {
		return make([]entity.Area, 0)
	}
	return activeAreas
}

func (service *areaService) FindNotActiveAreas() []entity.Area {
	notActiveAreas, err := driver.FindNotActiveAreas()
	if err != nil {
		return make([]entity.Area, 0)
	}
	return notActiveAreas
}
