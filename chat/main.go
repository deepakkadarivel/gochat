package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
	"github.com/gorilla/mux"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	router := mux.NewRouter()
	r := newRoom()
	router.Handle("/", &templateHandler{filename: "chat.html"})
	router.Handle("/room", r)
	go r.run()
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
