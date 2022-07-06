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
	myRouter.HandleFunc("/user/{id}", models.GetUser).Methods("GET")
	myRouter.HandleFunc("/user", models.CreateUser).Methods("POST")
	myRouter.HandleFunc("/user/{id}", models.UpdateUser).Methods("PUT")
	myRouter.HandleFunc("/user/{id}", models.DeleteUser).Methods("DELETE")
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatalln(http.ListenAndServe(":"+config.Config.Port, myRouter))
}
