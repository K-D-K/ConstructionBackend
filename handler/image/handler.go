package image

import (
	"encoding/json"
	"log"
	"net/http"
	"smfbackend/controller/image"
	"smfbackend/handler"
	smferror "smfbackend/utils/error"

	"github.com/jinzhu/gorm"
)

// POST : To create Image
func POST(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	r.ParseMultipartForm(32 << 20)
	uploadedFile, _, err := r.FormFile("uploadfile")
	defer uploadedFile.Close()
	if err != nil {
		log.Panicln("Error while reading file : " + err.Error())
		panic(smferror.ThrowAPIError("Error while reading file"))
	}
	description := r.FormValue("description")
	imageInstance := image.INSTANCE(db)
	imageData := imageInstance.Create(uploadedFile, description)

	byteArr, _ := json.Marshal(imageData)

	handler.RespondwithJSON(w, http.StatusOK, byteArr)
}
