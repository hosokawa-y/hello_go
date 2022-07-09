package controllers

import (
	"encoding/json"
	"hello/app/models"
	"io/ioutil"
	"net/http"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	todo := models.Todo{}
	json.Unmarshal(reqBody, &todo)

	//cmd := `insert into todos (content, user_id, created_at) values (?,?,?)`

	//_, err = Db.Exec(cmd, todo.Content, todo.UserID, time.Now())
	//if models.err != nil {
	//	log.Fatalln(models.err)
	//}
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//todos := Todos{}
	//
	//cmd := `select id, content, user_id, created_at from todos`
	//rows, err := Db.Query(cmd)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//for rows.Next() {
	//	var todo Todo
	//	err = rows.Scan(&todo.ID, &todo.Content, &todo.UserID, &todo.CreatedAt)
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//	todos = append(todos, todo)
	//}
	//rows.Close()
	//
	//json.NewEncoder(w).Encode(todos)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	//id := vars["id"]
	//todo := Todo{}
	//
	//cmd := `select id, content, user_id, created_at from todos where id = ?`
	//
	//err = Db.QueryRow(cmd, id).Scan(
	//	&todo.ID,
	//	&todo.Content,
	//	&todo.UserID,
	//	&todo.CreatedAt)
	//
	//json.NewEncoder(w).Encode(todo)
}

func GetTodosByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	//userID := vars["id"]
	//todos := Todos{}
	//cmd := `select id, user_id, content, created_at from todos where user_id = ?`
	//
	//rows, err := Db.Query(cmd, userID)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//for rows.Next() {
	//	var todo Todo
	//	err = rows.Scan(&todo.ID, &todo.UserID, &todo.Content, &todo.CreatedAt)
	//
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//	todos = append(todos, todo)
	//}
	//rows.Close()
	//
	//json.NewEncoder(w).Encode(todos)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//id := vars["id"]
	//reqBody, _ := ioutil.ReadAll(r.Body)
	//todo := Todo{}
	//json.Unmarshal(reqBody, &todo)
	//
	//cmd := `update todos set content = ?, user_id = ? where id = ?`
	//_, err = Db.Exec(cmd, todo.Content, todo.UserID, id)
	//if models.err != nil {
	//	log.Fatalln(models.err)
	//}
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//id := vars["id"]
	//
	//cmd := `delete from todos where id = ?`
	//_, err = Db.Exec(cmd, id)
	//if models.err != nil {
	//	log.Fatalln(models.err)
	//}
}
