package controllers

import (
	"log"
	"net/http"
	"strings"

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

func chatMessage(w http.ResponseWriter, r *http.Request) {
	sess, err := Session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		u, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		segs := strings.Split(r.URL.Path, "/")
		uuid := segs[3]
		if len(uuid) != 30 {
			http.NotFound(w, r)
			return
		}
		messages, err := models.GetMessagesByUUID(uuid)
		id := messages[0].WorkID
		data := models.Data{
			WorkID:   id,
			User:     u,
			ChatUUID: uuid,
			Messages: messages,
			Host:     r.Host,
		}
		generateHTML(w, data, "layout", "private_navbar", "chat/index", "js/chat")
	}
}
