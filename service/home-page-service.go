package service

import (
	"github.com/NomanSalhab/go_gin_my_first_project/driver"
	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

type HomePageService interface {
	GetHomePage(limit int, appVersion float32, wantedAreaId int) (entity.HomePage, error)
}

type homePageService struct {
	driver driver.HomePageDriver
}

func NewHomePageService(driver driver.HomePageDriver) HomePageService {
	return &homePageService{
		driver: driver,
	}
}

func (service *homePageService) GetHomePage(limit int, appVersion float32, wantedAreaId int) (entity.HomePage, error) {
	homePage, err := service.driver.GetHomePage(limit, appVersion, wantedAreaId)
	if err != nil {
		return homePage, err
	}
	return homePage, nil
}
