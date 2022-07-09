package models

import (
	"crypto/sha1"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"hello/config"
	"log"
)

var Db *gorm.DB
var err error

//type Model struct {
//	ID        uint       `gorm:"primary_key" json:"id"`
//	CreatedAt *time.Time `json:"created_at"`
//	UpdatedAt *time.Time `json:"updated_at"`
//	DeletedAt *time.Time `json:"deleted_at"`
//}

type User struct {
	gorm.Model `json:"model"`
	UUID       string `json:"UUID"`
	Name       string `json:"Name"`
	Email      string `json:"Email"`
	PassWord   string `json:"PassWord"`
}

func GetAllUsers(users *[]User) {
	Db.Find(&users)
}

func InsertUser(user *User) {
	Db.Create(&user)
}

func

func Connect() {
	Db, err = gorm.Open(sqlite.Open(config.Config.DbName), &gorm.Config{})
	fmt.Println(Db)
	if err != nil {
		log.Fatalln(err)
		panic("failed to connect database")
	} else {
		fmt.Println("Successfully connect database")
	}
	Db.AutoMigrate(&User{}, &Todo{})

	if !(Db.Migrator().HasTable(&User{})) {
		Db.Migrator().CreateTable(&User{})
	}
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
