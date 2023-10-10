package service

import (
	"errors"

	"github.com/NomanSalhab/go_gin_my_first_project/driver"
	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

type AreaService interface {
	AddArea(entity.Area) error
	FindAllAreas() []entity.Area
	FindActiveAreas() []entity.Area
	FindNotActiveAreas() []entity.Area
	EditArea(areaEditInfo entity.AreaEditRequest) error
	ActivateArea(areaEditInfo entity.AreaActivateRequest) error
	DeactivateArea(areaEditInfo entity.AreaDeactivateRequest) error
	DeleteArea(areaDeleteInfo entity.AreaEditRequest) error

	AddMockAreas(areas []entity.Area)
}

type areaService struct {
	areas []entity.Area
}

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
	// successArea := area
	// if len(service.areas) > 0 {
	// 	successArea.ID = service.areas[len(service.areas)-1].ID + 1
	// } else {
	// 	successArea.ID = 1
	// }
	// service.areas = append(service.areas, successArea)
	// return nil
}

func (service *areaService) FindAllAreas() []entity.Area {
	areas, err := driver.FindAllAreas()
	if err != nil {
		return make([]entity.Area, 0)
	}
	return areas
	// return service.areas
}

func (service *areaService) FindActiveAreas() []entity.Area {
	activeAreas, err := driver.FindActiveAreas()
	if err != nil {
		return make([]entity.Area, 0)
	}
	return activeAreas
	// var activeAreas []entity.Area
	// for i := 0; i < len(service.areas); i++ {
	// 	if service.areas[i].Active {
	// 		activeAreas = append(activeAreas, service.areas[i])
	// 	}
	// }
	// return activeAreas
}

func (service *areaService) FindNotActiveAreas() []entity.Area {
	notActiveAreas, err := driver.FindNotActiveAreas()
	if err != nil {
		return make([]entity.Area, 0)
	}
	return notActiveAreas
	// var notActiveAreas []entity.Area
	// for i := 0; i < len(service.areas); i++ {
	// 	if !service.areas[i].Active {
	// 		notActiveAreas = append(notActiveAreas, service.areas[i])
	// 	}
	// }
	// return notActiveAreas
}

// func (service *areaService) FindArea(id entity.AreaEditRequest) (entity.Area, error) {
// 	area, _ := driver.FindArea(id.ID)
// 	if area.Name == "" {
// 		return area, errors.New("the area couldn't be found")
// 	}
// 	return area, nil
// 	// areas := service.FindAllAreas()
// 	// var area entity.Area
// 	// for i := 0; i < len(areas) && len(areas) != 0; i++ {
// 	// 	if id.ID != 0 {
// 	// 		if areas[i].ID == id.ID {
// 	// 			area = areas[i]
// 	// 			return area, nil
// 	// 		}
// 	// 	} else {
// 	// 		return area, errors.New("store category id cannot be zero")
// 	// 	}
// 	// }
// 	// return area, errors.New("the store category couldn't be found")
// }

func (service *areaService) EditArea(areaEditInfo entity.AreaEditRequest) error {
	_, err := driver.EditArea(areaEditInfo)
	if err != nil {
		return err
	}
	return nil
	// areas := service.FindAllAreas()
	// var area entity.Area
	// for i := 0; i < len(areas) && len(areas) != 0; i++ {
	// 	if areaEditInfo.ID != 0 {
	// 		if areas[i].ID == areaEditInfo.ID {
	// 			area.ID = areaEditInfo.ID
	// 			return nil
	// 		}
	// 	} else {
	// 		return errors.New("store category id cannot be zero")
	// 	}
	// 	if len(areaEditInfo.Name) != 0 {
	// 		if areas[i].Name == areaEditInfo.Name {
	// 			area.Name = areaEditInfo.Name
	// 			return nil
	// 		}
	// 	}
	// 	if areaEditInfo.Lat != 0 {
	// 		if areas[i].Lat == areaEditInfo.Lat {
	// 			area.Lat = areaEditInfo.Lat
	// 			return nil
	// 		}
	// 	}
	// 	if areaEditInfo.Long != 0 {
	// 		if areas[i].Long == areaEditInfo.Long {
	// 			area.Long = areaEditInfo.Long
	// 			return nil
	// 		}
	// 	}
	// }
	// return errors.New("the store category couldn't be found")
}

func (service *areaService) ActivateArea(areaEditInfo entity.AreaActivateRequest) error {
	err := driver.ActivateArea(areaEditInfo)
	if err != nil {
		return err
	}
	return nil
	// areas := service.FindAllAreas()
	// var area entity.Area
	// for i := 0; i < len(areas) && len(areas) != 0; i++ {
	// 	if areaEditInfo.ID != 0 {
	// 		if areas[i].ID == areaEditInfo.ID {
	// 			area.ID = areaEditInfo.ID
	// 			areas[i].Active = true
	// 			return nil
	// 		}
	// 	} else {
	// 		return errors.New("store category id cannot be zero")
	// 	}
	// }
	// return errors.New("the store category couldn't be found")
}

func (service *areaService) DeactivateArea(areaEditInfo entity.AreaDeactivateRequest) error {
	err := driver.DeactivateArea(areaEditInfo)
	if err != nil {
		return err
	}
	return nil
	// areas := service.FindAllAreas()
	// var area entity.Area
	// for i := 0; i < len(areas) && len(areas) != 0; i++ {
	// 	if areaEditInfo.ID != 0 {
	// 		if areas[i].ID == areaEditInfo.ID {
	// 			area.ID = areaEditInfo.ID
	// 			areas[i].Active = false
	// 			return nil
	// 		}
	// 	} else {
	// 		return errors.New("store category id cannot be zero")
	// 	}
	// }
	// return errors.New("the store category couldn't be found")
}

func (service *areaService) DeleteArea(areaDeleteInfo entity.AreaEditRequest) error {
	err := driver.DeleteArea(areaDeleteInfo.ID)
	if err != nil {
		return err
	}
	return nil
	// areas := service.FindAllAreas()
	// var tempArea []entity.Area
	// for i := 0; i < len(areas) && len(areas) != 0; i++ {
	// 	if areaDeleteInfo.ID != 0 {
	// 		if areas[i].ID != areaDeleteInfo.ID {
	// 			tempArea = append(tempArea, areas[i])
	// 		}
	// 	} else {
	// 		return errors.New("id cannot be zero")
	// 	}
	// }
	// if len(areas) != (len(tempArea) + 1) {
	// 	return errors.New("area could not be found")
	// }
	// service.areas = tempArea
	// return nil
}

func (service *areaService) AddMockAreas(areas []entity.Area) {
	service.areas = append(service.areas, areas...)
}
