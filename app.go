package main

import (
	"net/http"

	"smfbackend/handler"
	"smfbackend/handler/image"
	"smfbackend/handler/project"
	"smfbackend/migration"

	"github.com/gorilla/mux"
)

/*
	Run server command
	go run .

	https://github.com/J7mbo/go-subdirectories-with-modules
	https://github.com/golang/go/wiki/Modules#when-should-i-use-the-replace-directive
	https://github.com/golang/go/wiki/Modules#faqs--multi-module-repositories

	Image Resize
	https://karthikkaranth.me/blog/image-resizing-server-go-libvips/

	Write File
	https://zupzup.org/go-http-file-upload-download/
	https://github.com/luisguve/golang-http-file-upload-download/blob/master/main.go

	Request Time out
	https://play.golang.org/p/v9IAu2Xu3_

	Bulk Insert
	https://github.com/t-tiger/gorm-bulk-insert/blob/master/bulk_insert.go

	Wait Group (sync/async)
	https://medium.com/@gauravsingharoy/asynchronous-programming-with-go-546b96cd50c1

	Beautify Golang project
	https://itnext.io/beautify-your-golang-project-f795b4b453aa

	GO Code Structure
	https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091
*/

/*
	1) Handle Bulk Update
*/

func main() {
	// Declare a new router
	router := mux.NewRouter()

	migration.RunMigration()

	router.HandleFunc("/projects/{project_id}", handler.ExecutorWithDB(project.GET)).Methods("GET")
	router.HandleFunc("/projects", handler.ExecutorWithDB(project.GET_ALL)).Methods("GET")
	router.HandleFunc("/projects", handler.ExecutorWithDB(project.POST)).Methods("POST")

	router.Handle("/images/{image_id}", http.StripPrefix("/images", http.FileServer(http.Dir("./Images"))))
	router.HandleFunc("/images", handler.ExecutorWithDB(image.POST)).Methods("POST")

	http.ListenAndServe(":8001", router)
}
