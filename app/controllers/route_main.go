package controllers

import (
	"net/http"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	time := time.Now()
	const layout = "2006/01/02"

	_, err := session(w, r)
	if err != nil {
		generateHTML(w, time.Format(layout), "layout", "public_navbar", "index")
	} else {
		generateHTML(w, time.Format(layout), "layout", "private_navbar", "index")
	}
}
