package main

import (
	"log"
	"net/http"
)

type server struct{}

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := `
	{"names":["Electric","Electric","Electric","Boogaloo","Boogaloo","Boogaloo","Boogaloo"]}
    `
	w.Write([]byte(msg))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}
