package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/keigooba/sharefull/app/models"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
)

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_, err := Session(w, r)
		if err != nil {
			generateHTML(w, nil, "layout", "public_navbar", "signup", "js/index")
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
		_, err := Session(w, r)
		if err != nil {
			generateHTML(w, nil, "layout", "public_navbar", "login", "js/index")
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

func auth(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	action := segs[2]
	provider := segs[3]
	switch action {
	case "login":
		switch provider {
		case "google":
			// 認証プロバイダー取得
			provider, err := gomniauth.Provider(provider)
			if err != nil {
				log.Fatalln("認証プロバイダーの取得に失敗しました:", provider, "-", err)
			}
			// 認証プロセスを開始するURLを取得
			loginUrl, err := provider.GetBeginAuthURL(nil, nil)
			if err != nil {
				log.Fatalln("認証プロセスを開始するURLの取得に失敗しました:", loginUrl, "-", err)
			}
			w.Header().Set("Location", loginUrl)
			w.WriteHeader(http.StatusTemporaryRedirect)
		default:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "プロバイダー%sには非対応です", provider)
		}
	case "callback":
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			log.Fatalln("認証プロバイダーの取得に失敗しました:", provider, "-", err)
		}

		creds, err := provider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
		if err != nil {
			log.Fatalln("認証を完了できませんでした", provider, "-", err)
		}

		creds_user, err := provider.GetUser(creds)
		if err != nil {
			log.Fatalln("ユーザーの取得に失敗しました", provider, "-", err)
		}

		user := models.User{
			Name:      creds_user.Name(),
			Email:     creds_user.Email(),
			PassWord:  models.RandString(5), //ランダムの5桁で生成
			AvatarURL: creds_user.AvatarURL(),
		}
		auth_user, err := user.AuthGetUser() //名前とメールアドレスで検索
		if err != nil { //ない場合生成
			if err := user.CreateUser(); err != nil {
				log.Fatalln(err)
			}
			auth_user, err = user.AuthGetUser()
			if err != nil {
				log.Fatalln(err)
			}
		}
		session, err := auth_user.CreateSession()
		if err != nil {
			log.Fatalln(err)
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			Path:     "/", //パスの指定が必要
			HttpOnly: true,
		})

		w.Header()["Location"] = []string{"/"}
		w.WriteHeader(http.StatusTemporaryRedirect)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "アクション%sには非対応です", action)

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
	http.SetCookie(w, &http.Cookie{
		Name:   "_cookie",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	http.Redirect(w, r, "/", 302)
}
