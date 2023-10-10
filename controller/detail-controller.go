package controller

import (
	"github.com/NomanSalhab/go_gin_my_first_project/entity"
	"github.com/NomanSalhab/go_gin_my_first_project/service"
	"github.com/gin-gonic/gin"
)

type DetailController interface {
	AddDetail(ctx *gin.Context) error
	FindAllDetails() []entity.DetailEditRequest
	FindAllAddons() []entity.Detail
	FindAllFlavors() []entity.Detail
	FindAllVolumes() []entity.Detail
	EditDetail(ctx *gin.Context) error
	DeleteDetail(ctx *gin.Context) error
}

type detailController struct {
	service service.DetailService
}

func NewDetailController(service service.DetailService) DetailController {
	return &detailController{
		service: service,
	}
}

func (c *detailController) FindAllDetails() []entity.DetailEditRequest {
	return c.service.FindAllDetails()
}

func (c *detailController) FindAllAddons() []entity.Detail {
	return c.service.FindAllAddons()
}

func (c *detailController) FindAllFlavors() []entity.Detail {
	return c.service.FindAllFlavors()
}

func (c *detailController) FindAllVolumes() []entity.Detail {
	return c.service.FindAllVolumes()
}

func (c *detailController) AddDetail(ctx *gin.Context) error /*entity.Video*/ {
	var detail entity.Detail
	err := ctx.ShouldBindJSON(&detail)
	if err != nil {
		return err
	}

	err = validate.Struct(detail)
	if err != nil {
		return err
	}

	err = c.service.AddDetail(detail)
	return err /*video*/
}

func (c *detailController) EditDetail(ctx *gin.Context) error {
	var detailEditInfo entity.DetailEditRequest
	err := ctx.ShouldBindJSON(&detailEditInfo)
	if err != nil {
		return err
	}
	err = validate.Struct(detailEditInfo)
	if err != nil {
		return err
	}
	err = c.service.EditDetail(detailEditInfo)
	if err != nil {
		return err
	}
	return nil
}

func (c *detailController) DeleteDetail(ctx *gin.Context) error {
	var detailId entity.DetailEditRequest
	err := ctx.ShouldBindJSON(&detailId)
	if err != nil {
		return err
	}
	err = c.service.DeleteDetail(detailId)
	if err != nil {
		return err
	}
	return nil
}
