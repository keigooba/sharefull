package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/keigooba/sharefull/app/models"
)

func index(w http.ResponseWriter, r *http.Request) {
	data := models.Data{}
	// 現在の日付
	now_date := time.Now().Format("2006/01/02")
	data.NowDate = now_date

	works, err := models.GetWorks()
	if err != nil {
		log.Fatalln(err)
	}
	data.Works = works

	sess, err := session(w, r)
	if err != nil {
		generateHTML(w, data, "layout", "public_navbar", "index")
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Fatalln(err)
		}
		data.User = user
		generateHTML(w, data, "layout", "private_navbar", "index")
	}
}

func workNew(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		if r.Method == "GET" {
			work := models.JobList()
			generateHTML(w, work, "layout", "private_navbar", "work_new")
		} else if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				log.Println(err)
			}

			user, err := sess.GetUserBySession()
			if err != nil {
				log.Println(err)
			}

			work := &models.Work{
				Date:      r.PostFormValue("date"),
				Title:     r.PostFormValue("title"),
				Money:     r.PostFormValue("money"),
				JobID:     r.PostFormValue("job_id"),
				Evalution: r.PostFormValue("evalution"),
			}
			if err = user.CreateWork(work); err != nil {
				log.Println(err)
			}
			http.Redirect(w, r, "/", 302)
		}
	}
}

func workEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		if r.Method == "GET" {
			work, err := models.GetWork(id)
			if err != nil {
				log.Println(err)
			}
			work.User = user
			generateHTML(w, work, "layout", "private_navbar", "work_edit")
		} else if r.Method == "POST" {
			work := &models.Work{
				ID:        id,
				Date:      r.PostFormValue("date"),
				Title:     r.PostFormValue("title"),
				Money:     r.PostFormValue("money"),
				JobID:     r.PostFormValue("job_id"),
				Evalution: r.PostFormValue("evalution"),
				UserID:    user.ID,
			}
			if err := work.UpdateWork(); err != nil {
				log.Println(err)
			}
			http.Redirect(w, r, "/", 302)
		}
	}
}

func workDelete(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		work, err := models.GetWork(id)
		if err != nil {
			log.Println(err)
		}
		if err := work.DeleteWork(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/", 302)
	}
}
