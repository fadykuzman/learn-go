package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type messageData struct {
	Names []string `json:names`
}

func getDataAndReadResponse() (int, int) {
	r, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	message := messageData{}
	err = json.Unmarshal(data, &message)
	electricCount := 0
	boogalooCount := 0
	for _, name := range message.Names {
		if name == "Electric" {
			electricCount++
		} else if name == "Boogaloo" {
			boogalooCount++
		}
	}
	return electricCount, boogalooCount
}

func main() {
	electric, boogaloo := getDataAndReadResponse()
	fmt.Println("Electric Count: ", electric)
	fmt.Println("Boogaloo Count: ", boogaloo)
}
