package main

import (
	"fmt"
	"log"
	"net/http"
)

type PageWithCounter struct {
	Counter int    `json:"counter"`
	Content string `json:"content"`
	Heading string `json:"heading"`
}

func (p *PageWithCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.Counter++
	msg := fmt.Sprintf(`
		<h1>%s</h1>
		<p>%s</p>
		<div>views: %d</div>
	`, p.Heading, p.Content, p.Counter)
	w.Write([]byte(msg))
}

func main() {
	hello := PageWithCounter{
		Heading: "Hello, World!",
		Content: "Hey, There!",
	}
	ch1 := PageWithCounter{
		Heading: "Chapter 1",
		Content: "Chapter 1: How to start",
	}
	ch2 := PageWithCounter{
		Heading: "Chapter 2",
		Content: "Chapter 2: More Advanced Greetings",
	}
	http.Handle("/", &hello)
	http.Handle("/chapter1", &ch1)
	http.Handle("/chapter2", &ch2)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
