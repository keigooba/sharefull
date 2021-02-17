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

	controllers.StartMainServer()

	// u := models.User{}
	// u.Name = "test"
	// u.Email = "test@email.com"
	// u.PassWord = "testtest"
	// u.CreateUser()

	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// u.CreateSession()

	// login_user := models.User{}
	// login_user.Email = "test@email.com"
	// login_user.PassWord = "testtest"

	// user, err := login_user.GetUserLogin()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(user)
}
