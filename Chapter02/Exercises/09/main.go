package main

import "fmt"

func main() {
	names := []string{"Jim", "Zita", "JoJo", "Jane"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}
