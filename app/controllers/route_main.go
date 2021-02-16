package controllers

import (
	"net/http"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	time := time.Now()
	const layout = "2006/01/02"
	generateHTML(w, time.Format(layout), "layout", "index")
}
