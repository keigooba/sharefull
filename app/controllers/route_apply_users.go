package controllers

import (
	"log"
	"net/http"

	"github.com/keigooba/sharefull/app/models"
)

func applyUser(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		u, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		applys_id, users_id, err := models.GetApplyUsersByWorkID(id)
		if err != nil {
			log.Println(err)
		}
		var users []models.User
		for _, v := range users_id {
			user, err := models.GetUser(v)
			if err != nil {
				log.Println(err)
			}
			users = append(users, user)
		}
		data := models.Data{User: u, ApplyUsers: users, ApplysID: applys_id}
		generateHTML(w, data, "layout", "private_navbar", "apply_user")
	}
}

func applyUserDelete(w http.ResponseWriter, r *http.Request, id int) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err := models.ApplyUserDelete(id)
		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/", 302)
	}
}
