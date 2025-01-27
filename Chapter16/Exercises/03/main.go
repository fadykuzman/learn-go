package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	vl := r.URL.Query()
	name, ok := vl["name"]

	if !ok {
		w.WriteHeader(400)
		w.Write([]byte("Missing name"))
		return
	}

	msg := fmt.Sprintf("<h1>Hello, %s</h1>", strings.Join(name, ","))
	w.Write([]byte(msg))
}

func main() {
	http.HandleFunc("/", Hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
