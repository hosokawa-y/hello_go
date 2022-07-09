package models

import (
	"encoding/json"
	"fmt"
	_ "github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	//"time"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	user := User{}

	json.Unmarshal(reqBody, &user)
	user.PassWord = Encrypt(user.PassWord)
	user.UUID = createUUID().String()
	InsertUser(&user)
	responseBody, err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("start get users")
	var users []User
	GetAllUsers(&users)
	fmt.Println("fetch users")
	responseBody, err := json.Marshal(users)
	//cmd := `select id, uuid, name, email, password, created_at from users`
	//rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	//for rows.Next() {
	//	var user User
	//	err = rows.Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.PassWord, &user.CreatedAt)
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//	users = append(users, user)
	//}
	//rows.Close()
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)

	//json.NewEncoder(w).Encode(users)
}

//func GetUser(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	vars := mux.Vars(r)
//	key := vars["id"]
//	user := User{}
//	cmd := `select id, uuid, name, email, password, created_at from users where id = ?`
//	err = Db.QueryRow(cmd, key).Scan(
//		&user.ID,
//		&user.UUID,
//		&user.Name,
//		&user.Email,
//		&user.PassWord,
//		&user.CreatedAt,
//	)
//	json.NewEncoder(w).Encode(user)
//}
//
//func UpdateUser(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	id := vars["id"]
//	reqBody, _ := ioutil.ReadAll(r.Body)
//	user := User{}
//	json.Unmarshal(reqBody, &user)
//
//	cmd := `update users set name = ?, email = ? where id = ?`
//	_, err = Db.Exec(cmd, user.Name, user.Email, id)
//	if err != nil {
//		log.Fatalln(err)
//	}
//}
//
//func DeleteUser(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	id := vars["id"]
//
//	cmd := `delete from users where id = ?`
//	_, err = Db.Exec(cmd, id)
//	if err != nil {
//		log.Fatalln(err)
//	}
//}
