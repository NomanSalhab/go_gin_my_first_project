package service

import (
	"errors"
	"time"

	"github.com/NomanSalhab/go_gin_my_first_project/driver"
	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

type ComplaintService interface {
	FindAllComplaints() ([]entity.Complaint, error)
	FindAboutDeliveryComplaints() ([]entity.Complaint, error)
	FindAboutTheAppComplaints() ([]entity.Complaint, error)
	FindImprovementSuggestionComplaints() ([]entity.Complaint, error)
	FindOtherReasonComplaints() ([]entity.Complaint, error)
	FindUserComplaints(wantedId int) ([]entity.Complaint, error)

	AddComplaint(complaint entity.Complaint) error
	DeleteComplaint(complaint entity.Complaint) error
}

type complaintService struct {
	driver driver.ComplaintDriver
}

func NewComplaintService(driver driver.ComplaintDriver) ComplaintService {
	return &complaintService{
		driver: driver,
	}
}

func (service *complaintService) AddComplaint(complaint entity.Complaint) error {
	if complaint.UserID != 0 {
		if len(complaint.Text) != 0 {
			if complaint.AboutDelivery || complaint.AboutTheApp || complaint.ImprovementSuggestion || complaint.OtherReason {
				complaint.Date = time.Now()
				err := service.driver.AddComplaint(complaint)
				if err != nil {
					return err
				}
				return nil
			} else {
				return errors.New("complaint type should be specified")
			}
		} else {
			return errors.New("complaint text should be added")
		}
	} else {
		return errors.New("complaint does not have a user id")
	}
}

func (service *complaintService) FindAllComplaints() ([]entity.Complaint, error) {
	allComplaints, err := service.driver.FindAllComplaints()
	if err != nil {
		return make([]entity.Complaint, 0), err
	}
	return allComplaints, nil
}

func (service *complaintService) FindAboutDeliveryComplaints() ([]entity.Complaint, error) {
	allAboutDeliveryComplaints, err := service.driver.FindAboutDeliveryComplaints()
	if err != nil {
		return make([]entity.Complaint, 0), err
	}
	return allAboutDeliveryComplaints, nil
}

func (service *complaintService) FindAboutTheAppComplaints() ([]entity.Complaint, error) {
	allAboutTheAppComplaints, err := service.driver.FindAboutTheAppComplaints()
	if err != nil {
		return make([]entity.Complaint, 0), err
	}
	return allAboutTheAppComplaints, nil
}

func (service *complaintService) FindImprovementSuggestionComplaints() ([]entity.Complaint, error) {
	allImprovementSuggetionsComplaints, err := service.driver.FindImprovementSuggestionComplaints()
	if err != nil {
		return make([]entity.Complaint, 0), err
	}
	return allImprovementSuggetionsComplaints, nil
}

func (service *complaintService) FindOtherReasonComplaints() ([]entity.Complaint, error) {
	allOtherReasonComplaints, err := service.driver.FindOtherReasonComplaints()
	if err != nil {
		return make([]entity.Complaint, 0), err
	}
	return allOtherReasonComplaints, nil
}

func (service *complaintService) FindUserComplaints(wantedId int) ([]entity.Complaint, error) {
	allUserComplaints, err := service.driver.FindUserComplaints(wantedId)
	if err != nil {
		return make([]entity.Complaint, 0), err
	}
	return allUserComplaints, nil
}

func (service *complaintService) DeleteComplaint(complaint entity.Complaint) error {
	err := service.driver.DeleteComplaint(complaint.ID)
	if err != nil {
		return err
	}
	return nil
}
