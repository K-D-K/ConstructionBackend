package image

import (
	"fmt"
	"mime/multipart"
	"smfbackend/imagehandler"
	"smfbackend/models"
	"smfbackend/utils"
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

// NOTE: Need to handle delete images case
func (controller Controller) AssociateImages(projects []models.Project) {
	projectIds := utils.ExtractProjectIds(projects)
	projectVsImageMap := groupImagesByProjects(controller.GetForProjects(projectIds))
	fmt.Println("Project ids ", projectIds)
	for _, project := range projects {
		images := project.Images
		var imageIds []uint
		associatedImages, ok := projectVsImageMap[project.ID]
		imageSet := map[uint]struct{}{}

		if ok {
			imageSet = createImageIdsSet(associatedImages)
		}

		for _, image := range images {
			if _, ok := imageSet[image.ID]; !ok {
				imageIds = append(imageIds, image.ID)
			}
		}

		fmt.Println("Images ", images)
		fmt.Println("Images Set ", imageSet)
		fmt.Println("Images ids ", imageIds)

		if len(imageIds) > 0 {
			controller.db.Model(&models.Image{}).Where("id in (?)", imageIds).Update("project_id", project.ID)
		}
	}
}

func groupImagesByProjects(images []models.Image) map[uint][]models.Image {
	imageMap := make(map[uint][]models.Image)
	for _, image := range images {
		projectImages, ok := imageMap[image.ProjectId]
		if !ok {
			projectImages = []models.Image{}
		}
		projectImages = append(projectImages, image)
		imageMap[image.ProjectId] = projectImages
	}
	return imageMap
}

func createImageIdsSet(images []models.Image) map[uint]struct{} {
	set := make(map[uint]struct{})
	for _, image := range images {
		set[image.ID] = struct{}{}
	}
	return set
}

func (controller Controller) GetForProjects(projectIDs []uint) []models.Image {
	images := []models.Image{}
	err := controller.db.Where("project_id in (?)", projectIDs).Find(&images).Error
	if err != nil {
		panic(err)
	}
	return images
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
