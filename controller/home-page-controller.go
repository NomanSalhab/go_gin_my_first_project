package controller

import (
	"github.com/NomanSalhab/go_gin_my_first_project/entity"
	"github.com/NomanSalhab/go_gin_my_first_project/service"
)

type HomePageController interface {
	GetHomePage(limit int, appVersion float32, wantedAreaId int) (entity.HomePage, error)
}

type homePageController struct {
	service service.HomePageService
}

func NewHomePageController(service service.HomePageService) HomePageController {
	return &homePageController{
		service: service,
	}
}

func (c *homePageController) GetHomePage(limit int, appVersion float32, wantedAreaId int) (entity.HomePage, error) {
	return c.service.GetHomePage(limit, appVersion, wantedAreaId)
}
