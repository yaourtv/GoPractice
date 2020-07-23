package main

import (
	"html/template"
	"path/filepath"
	"log"
	"net/http"
)

func onErrFail(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func errHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		tmpl, err := template.New("").ParseFiles(
			filepath.Join("templates", "layout.html"),
			filepath.Join("templates", "404.html"),
		)
		onErrFail(err)
		tmpl.ExecuteTemplate(w, "layout", nil)
	}
	if status == http.StatusMethodNotAllowed {
		tmpl, err := template.New("").ParseFiles(
			filepath.Join("templates", "layout.html"),
			filepath.Join("templates", "405.html"),
		)
		onErrFail(err)
		tmpl.ExecuteTemplate(w, "layout", nil)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errHandler(w, r, http.StatusNotFound)
		return
	}

	tmpl, err := template.New("").ParseFiles(
		filepath.Join("templates", "layout.html"),
		filepath.Join("templates", "home.html"),
	)
	onErrFail(err)

	switch r.Method {
	case http.MethodGet:
		tmpl.ExecuteTemplate(w, "layout", nil)
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			onErrFail(err)
			return
		}
		name := r.FormValue("name")
		address := r.FormValue("address")
		cookie := http.Cookie{
			Name: "token",
			Value: name + ":" + address,
		}
		http.SetCookie(w, &cookie);
		tmpl.ExecuteTemplate(w, "layout", nil)
	default:
		errHandler(w, r, http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
