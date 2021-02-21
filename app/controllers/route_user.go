package controllers

import (
	"log"
	"net/http"

	"github.com/keigooba/sharefull/app/models"
)

func userEdit(w http.ResponseWriter, r *http.Request, id int) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		if r.Method == "GET" {
			user, err := models.GetUser(id)
			if err != nil {
				log.Println(err)
			}

			data := models.Data{User: user}
			generateHTML(w, data, "layout", "private_navbar", "user_edit")
		} else if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				log.Println(err)
			}

			user := &models.User{
				ID:       id,
				Name:     r.PostFormValue("name"),
				Email:    r.PostFormValue("email"),
				PassWord: r.PostFormValue("password"),
			}
			if err := user.UpdateUser(); err != nil {
				log.Println(err)
			}
			http.Redirect(w, r, "/", 302)
		}
	}
}

func userDelete(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		user, err := models.GetUser(id)
		if err != nil {
			log.Println(err)
		}
		if err := user.DeleteUser(); err != nil {
			log.Println(err)
		}
		if err := sess.DeleteSessionByUUID(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/", 302)
	}
}

func userStatus(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		u, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		works_id, err := models.GetApplyUsersByUserID(id)
		if err != nil {
			log.Println(err)
		}

		// 応募した求人
		var works []models.Work
		for _, v := range works_id {
			work, err := models.GetWork(v)
			if err != nil {
				log.Println(err)
			}
			works = append(works, work)
		}

		// 応募中の求人
		a_works, err := u.GetWorksByUser()
		if err != nil {
			log.Println(err)
		}

		data := models.Data{Works: a_works, User: u}
		generateHTML(w, data, "layout", "private_navbar", "user_status")
	}
}
