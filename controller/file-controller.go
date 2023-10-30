package controller

import (
	"errors"
	"path/filepath"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
	"github.com/NomanSalhab/go_gin_my_first_project/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FileController interface {
	GetFileInfo(uuidValue string) (entity.File, error)
	AddFile(file entity.File) error
	DeleteFile(fileUUID string) error
}

type fileController struct {
	service service.FileService
}

func NewFileController(service service.FileService) FileController {
	return &fileController{
		service: service,
	}
}

func (c *fileController) AddFile(file entity.File) error {
	err := c.service.AddFile(file)
	return err
}

func (c *fileController) GetFileInfo(uuidValue string) (entity.File, error) {
	var fileInfo entity.File

	if len(uuidValue) == 0 {
		return entity.File{}, errors.New("file id can not be zero")
	}

	fileInfo = c.service.GetFileInfo(uuidValue)
	return fileInfo, nil
}

func (c *fileController) DeleteFile(fileUUID string) error {
	err := c.service.DeleteFile(fileUUID)
	if err != nil {
		return err
	}
	return nil
}

func GetImage(ctx *gin.Context) (entity.File, error) {

	file, err := ctx.FormFile("image")
	if err != nil {
		return entity.File{}, err
	}
	filePath := filepath.Join("images", file.Filename)
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		return entity.File{}, errors.New("failed to save file")
	}
	uuid := uuid.New().String()
	fileMetadata := entity.File{
		Filename: file.Filename,
		UUID:     uuid,
	}
	return fileMetadata, nil
}
