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

// CreateUser godoc
// @Summary create user
// @Description create user
// @Tags users
// @Accept  json
// @Produce  json
// @Param params body models.User true "user details"
// @Success 200 {array} models.User
// @Router /user [POST]
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

// GetUsers godoc
// @Summary Get details of all users
// @Description Get details of all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User
// @Router /users [GET]
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

// GetUser godoc
// @Summary Get details of user
// @Description Get details of user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "user search by id"
// @Success 200 {object} models.User
// @Router /user/{id} [GET]
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

// UpdateUser godoc
// @Summary update user
// @Description update user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "user search by id"
// @Param params body models.User true "user details"
// @Success 200 {object} models.User
// @Router /user/{id} [PUT]
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

// DeleteUser godoc
// @Summary delete user
// @Description delete user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "user search by id"
// @Success 200 {object} DeleteResponse
// @Router /user/{id} [DELETE]
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
