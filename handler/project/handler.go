package project

import (
	"encoding/json"
	"net/http"
	"smfbackend/controller/image"
	"smfbackend/controller/project"
	"smfbackend/handler"
	"smfbackend/models"
	"smfbackend/utils"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GET(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	params := mux.Vars(r)
	projectID, _ := strconv.ParseUint(params["project_id"], 10, 64)
	projectInstance := project.INSTANCE(db)
	project := projectInstance.GetProject(projectID)

	byteArr, _ := json.Marshal(project)

	handler.RespondwithJSON(w, http.StatusOK, byteArr)
}

func GET_ALL(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	query := r.URL.Query()
	fields := query["field"]
	projectInstance := project.INSTANCE(db)
	projects := projectInstance.GetProjects(fields)

	byteArr, _ := json.Marshal(projects)

	if len(fields) > 0 {
		byteArr, _ = utils.ExtractKeys(byteArr, fields)
	}

	handler.RespondwithJSON(w, http.StatusOK, byteArr)
}

func POST(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	decoder := json.NewDecoder(r.Body)
	decoder.Token()
	projects := []models.Project{}
	projectModel := models.Project{}
	projectInstance := project.INSTANCE(db)
	for decoder.More() {
		decoder.Decode(&projectModel)
		projectInstance.Create(&projectModel)
		projects = append(projects, projectModel)
	}

	imageInstance := image.INSTANCE(db)
	imageInstance.AssociateImages(projects)

	byteArr, _ := json.Marshal(projects)
	handler.RespondwithJSON(w, http.StatusOK, byteArr)
}
