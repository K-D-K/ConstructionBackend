package models

import "github.com/jinzhu/gorm"

type Image struct {
	gorm.Model
	Description string
	ProjectId   uint `gorm:"column:project_id"`
}
