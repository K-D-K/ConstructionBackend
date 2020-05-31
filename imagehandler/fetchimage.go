package imagehandler

import (
	"os"
	smferror "smfbackend/utils/error"
)

func GetImage(fileName string) *os.File {
	file, err := os.Open("./Images/" + fileName)
	if err != nil {
		panic(smferror.ThrowAPIError("Error while fetching Image"))
	}
	defer file.Close()
	return file
}
