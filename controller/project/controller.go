package project

import (
	"smfbackend/models"

	"github.com/jinzhu/gorm"
)

type Controller struct {
	db *gorm.DB
}

// INSTANCE : Get Project Controller instance
func INSTANCE(db *gorm.DB) *Controller {
	return &Controller{db}
}

func (controller Controller) Create(projects []models.Project) {
	for _, project := range projects {
		controller.db.Create(&project)
	}
}

func (controller Controller) GetProjects() []models.Project {
	projects := []models.Project{}
	err := controller.db.Preload("Address").Limit(4).Find(&projects).Error
	if err != nil {
		panic(err)
	}
	return projects
}
