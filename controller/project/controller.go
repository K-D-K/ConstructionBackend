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

func (controller Controller) Create(project interface{}) {
	controller.db.Create(project)
}

func (controller Controller) GetProjects(fields []string) []models.Project {
	projects := []models.Project{}
	db := controller.db
	if len(fields) > 0 {
		db = db.Select(fields)
	} else {
		db = db.Preload("Address").Preload("Images")
	}
	err := db.Find(&projects).Error
	if err != nil {
		panic(err)
	}
	return projects
}

func (controller Controller) GetProject(projectId uint64) models.Project {
	project := models.Project{}
	err := controller.db.Preload("Address").Preload("Images").First(&project, projectId).Error
	if err != nil {
		panic(err)
	}
	return project
}
