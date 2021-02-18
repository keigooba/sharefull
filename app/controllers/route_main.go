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
	timedate := time.Now()
	const layout = "2006/01/02"
	now_date := timedate.Format(layout)
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
