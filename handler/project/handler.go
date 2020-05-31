package project

import (
	"encoding/json"
	"net/http"
	"smfbackend/controller/image"
	"smfbackend/controller/project"
	"smfbackend/handler"
	"smfbackend/models"

	"github.com/jinzhu/gorm"
)

func GET(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	projectInstance := project.INSTANCE(db)
	projects := projectInstance.GetProjects()

	byteArr, _ := json.Marshal(projects)
	handler.RespondwithJSON(w, http.StatusOK, byteArr)
}

func POST(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	decoder := json.NewDecoder(r.Body)
	decoder.Token()
	projects := []models.Project{}
	projectModel := models.Project{}
	for decoder.More() {
		decoder.Decode(&projectModel)
		projects = append(projects, projectModel)
	}
	projectInstance := project.INSTANCE(db)
	projectInstance.Create(projects)

	imageInstance := image.INSTANCE(db)
	imageInstance.AssociateImages(projects)

	byteArr, _ := json.Marshal(projects)
	w.Write(byteArr)
}
