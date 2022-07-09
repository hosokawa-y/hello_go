package main

import (
	"hello/app/controllers"
	"hello/app/models"
)

func main() {
	models.Connect()
	controllers.HandleRequests()
}
