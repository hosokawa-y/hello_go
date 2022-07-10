package controllers

import (
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	"hello/config"
	_ "hello/docs"
	"log"
	"net/http"
)

// @title My Swagger Example API
// @version 2.0
// @description This is a sample server ToDoeeApp server.
// @termsOfService http://swagger.io/terms/

// @host localhost:8080
// @BasePath /

func HandleRequests() {
	// creates a new instance of a mux router
	//myRouter := mux.NewRouter().StrictSlash(true)
	myRouter := mux.NewRouter()
	myRouter.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", top)
	myRouter.HandleFunc("/users", GetUsers)
	myRouter.HandleFunc("/user/{id}", GetUser).Methods("GET")
	myRouter.HandleFunc("/user", CreateUser).Methods("POST")
	myRouter.HandleFunc("/user/{id}", UpdateUser).Methods("PUT")
	myRouter.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")

	myRouter.HandleFunc("/todos", GetTodos).Methods("GET")
	myRouter.HandleFunc("/todo", CreateTodo).Methods("POST")
	myRouter.HandleFunc("/todo/{id}", GetTodo).Methods("GET")
	myRouter.HandleFunc("/todo/{id}", UpdateTodo).Methods("PUT")
	myRouter.HandleFunc("/todo/{id}", DeleteTodo).Methods("DELETE")
	myRouter.HandleFunc("/todo/user/{id}", GetTodosByUser).Methods("GET")
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second argument
	log.Fatalln(http.ListenAndServe(":"+config.Config.Port, myRouter))
}
