package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Data struct {
	Host       string  `json:"host"`
	UserAgent  string  `json:"user_agent"`
	RequestURI string  `json:"request_uri"`
	Headers    headers `json:"headers"`
}

type headers struct {
	Accept    []string `json:"Accept"`
	UserAgent []string `json:"User-Agent"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := Data{
		r.Host,
		r.UserAgent(),
		r.RequestURI,
		headers{r.Header["Accept"], r.Header["User-Agent"]},
	}

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
