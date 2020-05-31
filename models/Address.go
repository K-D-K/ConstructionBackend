package models

import "github.com/jinzhu/gorm"

type Address struct {
	gorm.Model
	Name     string
	Street   string
	City     string
	District string
	Pin      string
	Mobile   string
	Mail     string
}
