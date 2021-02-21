package controllers

import (
	"log"
	"net/http"

	"github.com/keigooba/sharefull/app/models"
)

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, nil, "layout", "public_navbar", "signup")
		} else {
			http.Redirect(w, r, "/login", 302)
		}
	} else if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			log.Fatalln(err)
		}
		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil {
			log.Fatalln(err)
		}

		// ログイン処理
		login_user := models.User{
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		u, err := login_user.GetUserLogin()
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/login", 302)
		}

		session, err := u.CreateSession()
		if err != nil {
			log.Fatalln(err)
		}

		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, "/", 302)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, nil, "layout", "public_navbar", "login")
		} else {
			http.Redirect(w, r, "/", 302)
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		login_user := models.User{
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		u, err := login_user.GetUserLogin()
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/login", 302)
		} else {
			session, err := u.CreateSession()
			if err != nil {
				log.Fatalln(err)
			}

			cookie := http.Cookie{
				Name:     "_cookie",
				Value:    session.UUID,
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)

			http.Redirect(w, r, "/", 302)
		}
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}

	if err != http.ErrNoCookie {
		sess := models.Session{UUID: cookie.Value}
		sess.DeleteSessionByUUID()
	}
	http.Redirect(w, r, "/", 302)
}
