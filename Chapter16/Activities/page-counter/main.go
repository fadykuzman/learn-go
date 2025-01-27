package main

import (
	"fmt"
	"log"
	"net/http"
)

type PageWithCounter struct {
	counter int
	content string
	heading string
}

func (p PageWithCounter) ServeHttp(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf(`
		<h1>%s</h1>
		<p>%s</p>
		<div>views: %d</div>
	`, p.heading, p.content, p.counter)
	w.Write([]byte(msg))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", PageWithCounter{}))
}
