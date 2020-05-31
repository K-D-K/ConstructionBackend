package imagehandler

import (
	"io"
	"mime/multipart"
	"os"
)

func UploadImage(file multipart.File, imageId string) {
	createdFile, fileCreateErr := os.OpenFile("./Images/"+imageId, os.O_WRONLY|os.O_CREATE, 0666)
	defer createdFile.Close()
	if fileCreateErr != nil {
		panic(fileCreateErr)
	}
	_, fileCopyErr := io.Copy(createdFile, file)
	if fileCopyErr != nil {
		panic(fileCopyErr)
	}
}
