package models

import (
	"github.com/jinzhu/gorm"
)

// Project struct
type Project struct {
	gorm.Model
	Comment string
	Name    string
	Address Address `gorm:"foreignkey:Id"`
	Images  []Image `gorm:"association_save_reference:false"`
}
