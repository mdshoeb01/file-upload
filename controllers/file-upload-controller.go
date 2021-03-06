package controllers

import (
	"file-upload/models"
	"io/ioutil"
	"log"
	"net/http"

	// "os"

	"google.golang.org/protobuf/proto"
)

func FileUploadController(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

    w.Header().Set("Access-Control-Allow-Headers", "*")
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("Unable to read message from request : ", err)
		return
	}
	reqMsg := &models.FileUploadRequest{}
	err = proto.Unmarshal(data, reqMsg)
	if err != nil {
		log.Println(err)
	}

	ioutil.WriteFile("/Users/Mohammed.Shoeb/go/src/file-upload/data/" + reqMsg.Name + reqMsg.Type, reqMsg.GetData(), 777)
	

	w.WriteHeader(http.StatusAccepted)
}