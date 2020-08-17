package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"smfbackend/datastore"
	smferror "smfbackend/utils/error"

	"github.com/jinzhu/gorm"
)

func ExecutorWithDB(handler func(http.ResponseWriter, *http.Request, *gorm.DB)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		db := datastore.GetDBConnection().Begin()
		defer func() {
			if r := recover(); r != nil {
				db.Rollback()
				RespondWithError(w, r.(error))
			} else {
				db.Commit()
			}
		}()
		handler(w, r, db)
		defer db.Close()
	}
}

func Executor(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				RespondWithError(w, r.(error))
			}
		}()
		handler(w, r)
	}
}

// RespondwithJSON : generic handling to send response.
func RespondwithJSON(w http.ResponseWriter, code int, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// RespondWithError : handle errors in project
func RespondWithError(w http.ResponseWriter, err error) {
	switch err.(type) {
	case *smferror.APIError:
		byteArr, _ := json.Marshal(map[string]string{"message": err.Error()})
		RespondwithJSON(w, http.StatusBadRequest, byteArr)
	default:
		byteArr, _ := json.Marshal(map[string]string{"message": "Internal Server Error"})
		RespondwithJSON(w, http.StatusBadRequest, byteArr)
	}
}

// GetContentType : get content type for the file
func GetContentType(file *os.File) string {
	buffer := make([]byte, 512)
	file.Read(buffer)

	// Reset the read pointer if necessary.
	file.Seek(0, 0)

	// Always returns a valid content-type and "application/octet-stream" if no others seemed to match.
	return http.DetectContentType(buffer)
}
