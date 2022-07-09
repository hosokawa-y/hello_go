package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"hello/app/models"
	"io/ioutil"
	"log"
	"net/http"
)

type DeleteResponse struct {
	ID string `json:"id"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	user := models.User{}

	json.Unmarshal(reqBody, &user)
	user.Password = models.Encrypt(user.Password)
	user.UUID = models.CreateUUID().String()
	models.InsertUser(&user)
	responseBody, err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	models.FetchAllUsers(&users)
	responseBody, err := json.Marshal(users)

	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	user := models.User{}
	models.FetchUser(&user, id)
	responseBody, err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	user := models.User{}

	json.Unmarshal(reqBody, &user)
	models.UpdateUser(&user, id)
	responseBody, err := json.Marshal(user)

	if err != nil {
		log.Fatalln(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	models.DeleteUser(id)
	responseBody, err := json.Marshal(DeleteResponse{ID: id})

	if err != nil {
		log.Fatalln(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}
