package main

import (
	"log"
	"fmt"
	"file-upload/routes"
	"github.com/gorilla/mux"
	"net/http"
	"file-upload/config"
)


func main() {
	// controllers.SaveStarUser()
	config.ReadConfig()
	handleRequests()
}


func handleRequests() {
	muxRouter := mux.NewRouter()
	routes.Routes(muxRouter)
	port := fmt.Sprintf("%v%v", ":", config.Cfg.Server.Port)
	log.Println("port connected", port)
	log.Println(http.ListenAndServe(port, muxRouter))

}
