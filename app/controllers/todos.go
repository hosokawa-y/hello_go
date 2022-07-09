package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"hello/app/models"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	todo := models.Todo{}
	json.Unmarshal(reqBody, &todo)
	models.InsertTodo(&todo)
	responseBody, err := json.Marshal(todo)

	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)

}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	models.FetchAllTodos(&todos)
	responseBody, err := json.Marshal(todos)

	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	todo := models.Todo{}
	models.FetchTodo(&todo, id)
	responseBody, err := json.Marshal(todo)

	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func GetTodosByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	var todos []models.Todo
	models.FetchTodosByUser(&todos, userID)
	responseBody, err := json.Marshal(todos)

	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	todo := models.Todo{}

	json.Unmarshal(reqBody, &todo)
	models.UpdateTodo(&todo, id)
	responseBody, err := json.Marshal(todo)

	if err != nil {
		log.Fatalln(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	models.DeleteTodo(id)

	responseBody, err := json.Marshal(DeleteResponse{ID: id})

	if err != nil {
		log.Fatalln(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}
