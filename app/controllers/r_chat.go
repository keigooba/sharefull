package controllers

import (
	"log"
	"net/http"

	"github.com/keigooba/sharefull/app/models"
)

func workChat(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := Session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		u, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		data := models.Data{
			WorkID:   id,
			User:     u,
			ChatUUID: models.RandString(30),
			Host:     r.Host,
		}
		generateHTML(w, data, "layout", "private_navbar", "chat/index", "js/chat")
	}
}
