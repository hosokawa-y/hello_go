package controllers

import (
	"hello/app/models"
	"hello/config"
	"log"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/", top)
	http.HandleFunc("/users", models.GetUsers)
	log.Fatalln(http.ListenAndServe(":"+config.Config.Port, nil))
}
