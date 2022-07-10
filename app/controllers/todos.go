package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"hello/app/models"
	"io/ioutil"
	"log"
	"net/http"
)

// CreateTodo godoc
// @Summary create Todo
// @Description create Todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param params body models.Todo true "todo details"
// @Success 200 {array} models.Todo
// @Router /todo [POST]
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

// GetTodos godoc
// @Summary Get details of all todos
// @Description Get details of all todos
// @Tags todos
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Todo
// @Router /todos [GET]
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

// GetTodo godoc
// @Summary Get details of todo
// @Description Get details of todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param id path int true "todo search by id"
// @Success 200 {object} models.Todo
// @Router /todo/{id} [GET]
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

// GetTodosByUser godoc
// @Summary Get details of todo by user
// @Description Get details of todo by user
// @Tags todos
// @Accept  json
// @Produce  json
// @Param id path int true "user search by id"
// @Success 200 {array} models.Todo
// @Router /todo/user/{id} [GET]
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

// UpdateTodo godoc
// @Summary update todo
// @Description update todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param id path int true "todo search by id"
// @Param params body models.Todo true "user details"
// @Success 200 {object} models.Todo
// @Router /todo/{id} [PUT]
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

// DeleteTodo godoc
// @Summary delete todo
// @Description delete todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param id path int true "todo search by id"
// @Success 200 {object} DeleteResponse
// @Router /todo/{id} [DELETE]
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
