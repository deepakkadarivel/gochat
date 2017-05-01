package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
	"github.com/gorilla/mux"
	"flag"
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
	t.templ.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse()
	router := mux.NewRouter()
	r := newRoom()
	router.Handle("/", &templateHandler{filename: "chat.html"})
	router.Handle("/room", r)
	go r.run()
	log.Println("Started server and listining at port : ", *addr)
	if err := http.ListenAndServe(*addr, router); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
