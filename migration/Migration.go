package migration

import (
	"smfbackend/models"

	"github.com/jinzhu/gorm"
)

func RunMigration() {
	db, err := gorm.Open("postgres", "port=5432 user=kdk dbname=smf")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&models.Address{}, &models.Project{}, &models.Image{})
}
