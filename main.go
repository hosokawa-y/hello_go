package main

import (
	"fmt"
	"hello/app/controllers"
	"hello/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()

}
