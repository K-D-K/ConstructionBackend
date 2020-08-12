package models

type Image struct {
	Model
	Description string `json:"description"`
	ProjectId   uint   `gorm:"column:project_id" json:"project_id"`
}
