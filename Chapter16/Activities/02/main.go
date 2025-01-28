package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Visitor struct {
	Name string
}

type Hello struct {
	tmpl *template.Template
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vis := Visitor{}
	vl := r.URL.Query()
	name, ok := vl["name"]
	if ok {
		vis.Name = strings.Join(name, ",")
	}
	h.tmpl.Execute(w, vis)
}

func newHello(tplPath string) (*Hello, error) {
	tmpl, error := template.ParseFiles(tplPath)
	if error != nil {
		return nil, error
	}
	return &Hello{tmpl}, nil
}
func main() {
	hello, err := newHello("index.html")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
