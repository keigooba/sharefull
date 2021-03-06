package controllers

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/keigooba/sharefull/app/models"
)

func userEdit(w http.ResponseWriter, r *http.Request, id int) {
	_, err := Session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		user, err := models.GetUser(id)
		if err != nil {
			log.Println(err)
		}
		if r.Method == "GET" {
			data := models.Data{User: user}
			generateHTML(w, data, "layout", "private_navbar", "user_edit", "js/index")
		} else if r.Method == "POST" {
			err := r.ParseMultipartForm(32 << 20)
			if err != nil {
				log.Println(err)
			}

			// サーバー上に画像ファイルを保存
			file, header, err := r.FormFile("avatar_url")
			if err != nil {
				io.WriteString(w, err.Error())
				return
			}
			defer file.Close()
			data, err := ioutil.ReadAll(file) //バイト列のデータをすべて持つ
			if err != nil {
				io.WriteString(w, err.Error())
				return
			}
			filename := filepath.Join("app/views/avatars", user.AvatarID+filepath.Ext(header.Filename))
			fmt.Println(filename)
			err = ioutil.WriteFile(filename, data, 0777)
			if err != nil {
				io.WriteString(w, err.Error())
				return
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
	sess, err := Session(w, r)
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
	sess, err := Session(w, r)
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
		a_works, _ := u.GetWorksByUser()

		//応募中の求人からのメッセージ
		var messages []models.Message
		for _, v := range a_works {
			message, err := models.GetChatUUIDByWorkID(v.ID)
			if err == nil {
				messages = append(messages, message)
			}
		}
		// 送ったメッセージ
		s_messages, _ := models.GetChatUUIDByUserID(u.ID)

		data := models.Data{Works: a_works, User: u, Messages: messages, SendMessages: s_messages}
		generateHTML(w, data, "layout", "private_navbar", "user_status", "js/index")
	}
}
