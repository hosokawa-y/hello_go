package models

import (
	"crypto/sha1"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"hello/config"
	"log"
	"time"
)

var Db *gorm.DB
var err error

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type User struct {
	Model    `json:"model"`
	UUID     string `json:"UUID"`
	Name     string `json:"Name"`
	Email    string `json:"Email"`
	Password string `json:"PassWord"`
}

type Todo struct {
	gorm.Model
	Content string `json:"Content"`
	UserID  int    `json:"UserID"`
}

func FetchAllUsers(users *[]User) {
	Db.Find(&users)
}

func InsertUser(user *User) {
	Db.Create(&user)
}

func FetchUser(user *User, id string) {
	Db.First(&user, id)

}

func UpdateUser(user *User, id string) {
	Db.Model(&user).Where("id = ?", id).Updates(map[string]interface{}{
		"Name":     user.Name,
		"Email":    user.Email,
		"Password": Encrypt(user.Password),
	})
}

func DeleteUser(id string) {
	Db.Where("id = ?", id).Delete(&User{})
}

func FetchAllTodos(todos *[]Todo) {
	Db.Find(&todos)
}

func FetchTodosByUser(todos *[]Todo, id string) {
	Db.Where("user_id = ?", id).Find(&todos)
}

func FetchTodo(todo *Todo, id string) {
	Db.Where("id = ?", id).Find(&todo)
}

func InsertTodo(todo *Todo) {
	Db.Create(&todo)
}

func UpdateTodo(todo *Todo, id string) {
	Db.Model(&todo).Where("id = ?", id).Updates(map[string]interface{}{
		"Content": todo.Content,
		"UserID":  todo.UserID,
	})
}

func DeleteTodo(id string) {
	Db.Where("id = ?", id).Delete(&Todo{})
}

func Connect() {
	Db, err = gorm.Open(sqlite.Open(config.Config.SQLDriver), &gorm.Config{})
	fmt.Println(Db)
	if err != nil {
		log.Fatalln(err)
		panic("failed to connect database")
	} else {
		fmt.Println("Successfully connect database")
	}
	Db.AutoMigrate(&User{}, &Todo{})

	if !(Db.Migrator().HasTable(&User{})) {
		Db.Migrator().CreateTable(&User{}, &Todo{})
	}
}

func CreateUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
