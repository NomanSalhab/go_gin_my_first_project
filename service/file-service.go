package service

import (
	"github.com/NomanSalhab/go_gin_my_first_project/driver"
	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

type FileService interface {
	GetFileInfo(uuidValue string) entity.File
	AddFile(file entity.File) error
	DeleteFile(fileUUID string) error
}

type fileService struct {
	driver driver.FileDriver
}

func NewFileService(driver driver.FileDriver) FileService {
	return &fileService{
		driver: driver,
	}
}

func (service *fileService) AddFile(file entity.File) error {
	err := service.driver.AddFile(file)
	return err
}

func (service *fileService) DeleteFile(fileUUID string) error {
	err := service.driver.DeleteFile(fileUUID)
	if err != nil {
		return err
	}
	return nil
}

func (service *fileService) GetFileInfo(uuidValue string) entity.File {
	file, err := service.driver.GetFileInfo(uuidValue)
	if err != nil {
		return entity.File{}
	}
	return file
	// return service.details
}
