package image

import (
	"fmt"
	"mime/multipart"
	"smfbackend/imagehandler"
	"smfbackend/models"
	"strings"

	smferror "smfbackend/utils/error"

	"github.com/jinzhu/gorm"
)

type Controller struct {
	db *gorm.DB
}

// INSTANCE : Get Image Controller instance
func INSTANCE(db *gorm.DB) *Controller {
	return &Controller{db}
}

func (controller Controller) AssociateImages(projects []models.Project) {
	for _, project := range projects {
		images := project.Images
		projectID := project.ID
		imageIds := make([]uint, len(images))
		for index, image := range images {
			imageIds[index] = image.ID
		}

		associatedImages, err := controller.GetForProject(projectID)
		if err != nil {
			panic(err)
		}
		associatedImages = associatedImages
	}
}

func (controller Controller) GetForProject(projectID uint) ([]models.Image, error) {
	images := []models.Image{}
	err := controller.db.Where("project_id = ?", projectID).Find(&images).Error
	return images, err
}

// Create an image
func (controller Controller) Create(file multipart.File, description string) models.Image {
	image := models.Image{}
	if len(strings.TrimSpace(description)) != 0 {
		image.Description = description
	}

	createImageError := controller.db.Create(&image).Error
	if createImageError != nil {
		smferror.ThrowAPIError("Error while creating Image")
	}

	imagehandler.UploadImage(file, fmt.Sprint(image.ID))
	return image
}
