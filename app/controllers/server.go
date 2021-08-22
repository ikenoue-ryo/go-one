package controllers

import (
	"fmt"
	"go-one/config"
	"html/template"
	"net/http"
)

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

func StartMainServer() error {
	files := http.FileServer(http.Dir("app/views/img"))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", top)
	http.HandleFunc("/index", index)
	http.HandleFunc("/signup", signup)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
