package main

import (
	"flag"
	"fmt"
)

var (
	nameFlag  = flag.String("name", "Sam", "Name of the person to greet")
	quietFlag = flag.Bool("quiet", false, "Toggle to be quiet when saying hello")
)

func main() {
	flag.Parse()
	if !*quietFlag {
		greeting := fmt.Sprintf("Hello, %s", *nameFlag)
		fmt.Println(greeting)
	}
}
