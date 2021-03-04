package routes



import (
	"github.com/gorilla/mux"
	"file-upload/controllers"
)

func Routes(router *mux.Router) {

	router.HandleFunc("/upload", controllers.FileUploadController).Methods("POST", "OPTIONS")

}