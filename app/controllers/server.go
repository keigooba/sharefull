package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/keigooba/sharefull/app/models"
	"github.com/keigooba/sharefull/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func Session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

var validPathWork = regexp.MustCompile("^/work/(edit|delete|apply|chat)/([0-9]+)$")

var validPathUser = regexp.MustCompile("^/user/(edit|delete|status)/([0-9]+)$")

var validPathApplyUser = regexp.MustCompile("^/apply/(status|delete)/([0-9]+)$")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "work") {
			//work/edit/1
			q := validPathWork.FindStringSubmatch(r.URL.Path)
			if q == nil {
				http.NotFound(w, r)
				return
			}
			qi, err := strconv.Atoi(q[2])
			if err != nil {
				http.NotFound(w, r)
				return
			}
			fn(w, r, qi)
		} else if strings.Contains(r.URL.Path, "user") {
			//user/edit/1
			q := validPathUser.FindStringSubmatch(r.URL.Path)
			if q == nil {
				http.NotFound(w, r)
				return
			}
			qi, err := strconv.Atoi(q[2])
			if err != nil {
				http.NotFound(w, r)
				return
			}
			fn(w, r, qi)
		} else if strings.Contains(r.URL.Path, "apply") {
			//apply/user/1
			q := validPathApplyUser.FindStringSubmatch(r.URL.Path)
			if q == nil {
				http.NotFound(w, r)
				return
			}
			qi, err := strconv.Atoi(q[2])
			if err != nil {
				http.NotFound(w, r)
				return
			}
			fn(w, r, qi)
		}
	}
}

func StartMainServer() error {

	// app/views以下ファイル読み込み
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// チャットルームの作成
	r := models.NewRoom()
	http.Handle("/room", r)
	go r.Run() //チャネルが入ると処理

	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/auth/", auth)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/work/new", workNew)
	http.HandleFunc("/work/edit/", parseURL(workEdit))
	http.HandleFunc("/work/delete/", parseURL(workDelete))
	http.HandleFunc("/work/apply/", parseURL(workApply))
	http.HandleFunc("/user/edit/", parseURL(userEdit))
	http.HandleFunc("/work/chat/", parseURL(workChat))
	http.HandleFunc("/chat/message/", chatMessage)
	http.HandleFunc("/user/delete/", parseURL(userDelete))
	http.HandleFunc("/user/status/", parseURL(userStatus))
	http.HandleFunc("/apply/status/", parseURL(applyUser))
	http.HandleFunc("/apply/delete/", parseURL(applyUserDelete))
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
