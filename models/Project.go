package models

// Project struct
type Project struct {
	Model
	Comment   string  `json:"comment"`
	Name      string  `json:"name"`
	Address   Address `json:"address"`
	AddressID uint    `json:"-"`
	Images    []Image `gorm:"association_save_reference:false" json:"images"`
}
