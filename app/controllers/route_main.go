package controllers

import (
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "hello", "layout", "top")
}

func index(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "index", "layout", "index")
}
