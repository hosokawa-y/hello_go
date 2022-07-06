package controllers

import (
	"github.com/gorilla/mux"
	"hello/app/models"
	"hello/config"
	"log"
	"net/http"
)

func HandleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", top)
	myRouter.HandleFunc("/users", models.GetUsers)
	myRouter.HandleFunc("/user/{id}", models.GetUser)
	myRouter.HandleFunc("/user", models.CreateUser).Methods("POST")
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatalln(http.ListenAndServe(":"+config.Config.Port, myRouter))
}
