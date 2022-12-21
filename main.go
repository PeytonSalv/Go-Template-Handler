package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

type PageData struct {
	Title string
	Body  string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		indexPath, err := filepath.Glob("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		stylePath, err := filepath.Glob("style.css")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(indexPath) == 0 {
			http.Error(w, "index.html file not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "text/html")

		data := PageData{
			Title: "My Page",
			Body:  "Welcome to My Page",
		}

		tmpl, err := template.ParseFiles(indexPath[0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(stylePath) == 0 {
			return
		}

		w.Header().Set("Content-Type", "text/css")

		style, err := ioutil.ReadFile(stylePath[0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = w.Write(style)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":8080", nil)
}
