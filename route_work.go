package main

import (
	"net/http"

	// "github.com/keigooba/sharefull/data"
)
// GET /work/new 求人作成
func newWork(w http.ResponseWriter, r *http.Request){
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private.navbar", "new.work")
	}
}

