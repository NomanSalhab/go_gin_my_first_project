package controller

import (
	"github.com/NomanSalhab/go_gin_my_first_project/entity"
	"github.com/NomanSalhab/go_gin_my_first_project/service"
	"github.com/gin-gonic/gin"
)

type AreaController interface {
	AddArea(ctx *gin.Context) error
	FindAllAreas() []entity.Area
	FindActiveAreas() []entity.Area
	FindNotActiveAreas() []entity.Area
	EditArea(ctx *gin.Context) error
	ActivateArea(ctx *gin.Context) error
	DeactivateArea(ctx *gin.Context) error
	DeleteArea(ctx *gin.Context) error
}

type areaController struct {
	service service.AreaService
}

func NewAreaController(service service.AreaService) AreaController {
	return &areaController{
		service: service,
	}
}

func (c *areaController) FindAllAreas() []entity.Area {
	return c.service.FindAllAreas()
}

func (c *areaController) AddArea(ctx *gin.Context) error /*entity.Video*/ {
	var area entity.Area
	err := ctx.ShouldBindJSON(&area)
	if err != nil {
		return err
	}

	err = validate.Struct(area)
	if err != nil {
		return err
	}

	err = c.service.AddArea(area)
	return err /*video*/
}

func (c *areaController) FindActiveAreas() []entity.Area {
	return c.service.FindActiveAreas()
}

func (c *areaController) FindNotActiveAreas() []entity.Area {
	return c.service.FindNotActiveAreas()
}

func (c *areaController) EditArea(ctx *gin.Context) error {
	var areaEditInfo entity.AreaEditRequest
	err := ctx.ShouldBindJSON(&areaEditInfo)
	if err != nil {
		return err
	}
	err = validate.Struct(areaEditInfo)
	if err != nil {
		return err
	}
	err = c.service.EditArea(areaEditInfo)
	if err != nil {
		return err
	}
	return nil
}

func (c *areaController) ActivateArea(ctx *gin.Context) error {
	var areaActivateInfo entity.AreaActivateRequest
	err := ctx.ShouldBindJSON(&areaActivateInfo)
	if err != nil {
		return err
	}
	err = validate.Struct(areaActivateInfo)
	if err != nil {
		return err
	}
	err = c.service.ActivateArea(areaActivateInfo)
	if err != nil {
		return err
	}
	return nil
}

func (c *areaController) DeactivateArea(ctx *gin.Context) error {
	var areaDeactivateInfo entity.AreaDeactivateRequest
	err := ctx.ShouldBindJSON(&areaDeactivateInfo)
	if err != nil {
		return err
	}
	err = validate.Struct(areaDeactivateInfo)
	if err != nil {
		return err
	}
	err = c.service.DeactivateArea(areaDeactivateInfo)
	if err != nil {
		return err
	}
	return nil
}

func (c *areaController) DeleteArea(ctx *gin.Context) error {
	var areaId entity.AreaEditRequest
	err := ctx.ShouldBindJSON(&areaId)
	if err != nil {
		return err
	}
	err = c.service.DeleteArea(areaId)
	if err != nil {
		return err
	}
	return nil
}
