package models

import "time"

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id,string"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}