package controllers

import (
	"log"
	"net/http"

	"github.com/keigooba/sharefull/app/models"
)

func workChat(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		u, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		data := models.Data{
			User: u,
		}
		generateHTML(w, data, "layout", "private_navbar", "chat/index")
	}
}
