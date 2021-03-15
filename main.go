package main

import (
	"log"

	"github.com/keigooba/sharefull/app/controllers"
	"github.com/keigooba/sharefull/app/models"
	"github.com/keigooba/sharefull/config"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
)

func main() {
	test := models.Db
	if test != nil {
		log.Println("success")
	}

	// Gomniauthのセットアップ
	gomniauth.SetSecurityKey(models.RandString(20))
	gomniauth.WithProviders(
		google.New("587125882573-qc7jp6pmuvps29kd541qi0qds97480mt.apps.googleusercontent.com", "j409W6R1jm6kB8NNrEOWfFjk", config.Config.Url+"auth/callback/google"),
	)

	// models.Migration()
	controllers.StartMainServer()
}
