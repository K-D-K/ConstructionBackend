package models

type Address struct {
	Model
	Name     string `json:"name"`
	Street   string `json:"street"`
	City     string `json:"city"`
	District string `json:"district"`
	Pin      string `json:"pin"`
	Mobile   string `json:"mobile"`
	Mail     string `json:"mail"`
	Location string `json:"location"`
}
