package main

import (
	"log"

	"github.com/keigooba/sharefull/app/controllers"
	"github.com/keigooba/sharefull/app/models"
)

func main() {
	test := models.Db
	if test != nil {
		log.Println("success")
	}

	// models.Migration()
	controllers.StartMainServer()
}
