package controller

import (
	"github.com/NomanSalhab/go_gin_my_first_project/entity"
	"github.com/NomanSalhab/go_gin_my_first_project/service"
	"github.com/gin-gonic/gin"
)

type ComplaintController interface {
	FindAllComplaints() ([]entity.Complaint, error)
	FindAboutDeliveryComplaints() ([]entity.Complaint, error)
	FindAboutTheAppComplaints() ([]entity.Complaint, error)
	FindImprovementSuggestionComplaints() ([]entity.Complaint, error)
	FindOtherReasonComplaints() ([]entity.Complaint, error)
	FindUserComplaints(wantedId int) ([]entity.Complaint, error)

	AddComplaint(ctx *gin.Context) error
	DeleteComplaint(ctx *gin.Context) error
}

type complaintController struct {
	service service.ComplaintService
}

func NewComplaintController(service service.ComplaintService) ComplaintController {
	return &complaintController{
		service: service,
	}
}

func (c *complaintController) FindAllComplaints() ([]entity.Complaint, error) {
	return c.service.FindAllComplaints()
}

func (c *complaintController) FindAboutDeliveryComplaints() ([]entity.Complaint, error) {
	return c.service.FindAboutDeliveryComplaints()
}

func (c *complaintController) FindAboutTheAppComplaints() ([]entity.Complaint, error) {
	return c.service.FindAboutTheAppComplaints()
}

func (c *complaintController) FindImprovementSuggestionComplaints() ([]entity.Complaint, error) {
	return c.service.FindImprovementSuggestionComplaints()
}

func (c *complaintController) FindOtherReasonComplaints() ([]entity.Complaint, error) {
	return c.service.FindOtherReasonComplaints()
}

func (c *complaintController) FindUserComplaints(wantedId int) ([]entity.Complaint, error) {
	return c.service.FindUserComplaints(wantedId)
}

func (c *complaintController) AddComplaint(ctx *gin.Context) error {
	var complaint entity.Complaint
	err := ctx.ShouldBindJSON(&complaint)
	if err != nil {
		return err
	}

	err = validate.Struct(complaint)
	if err != nil {
		return err
	}

	err = c.service.AddComplaint(complaint)
	return err /*video*/
}

func (c *complaintController) DeleteComplaint(ctx *gin.Context) error {
	var complaint entity.Complaint
	err := ctx.ShouldBindJSON(&complaint)
	if err != nil {
		return err
	}
	err = c.service.DeleteComplaint(complaint)
	if err != nil {
		return err
	}
	return nil
}
