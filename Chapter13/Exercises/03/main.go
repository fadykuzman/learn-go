package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func rot13(s string) string {
	result := make([]byte, len(s))

	for i := 0; i < len(s); i++ {
		char := s[i]
		switch {
		case char >= 'a' && char <= 'z':
			result[i] = 'a' + (char-'a'+13)%26
		case char >= 'A' && char <= 'Z':
			result[i] = 'A' + (char-'A'+13)%26
		default:
			result[i] = char
		}
	}
	return string(result)
}

func processStdin() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, error := reader.ReadString('\n')
		if error == io.EOF {
			break
		} else if error != nil {
			fmt.Println("Error reading stdin:", error)
		}
		encoded := rot13(input)
		fmt.Println(encoded)
	}
}

func processInput() {
	var inputReader io.Reader

	args := os.Args
	if len(args) > 1 {
		file, err := os.Open(args[1])
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()
		inputReader = file
	} else {
		fmt.Print("Enter text:")
		inputReader = os.Stdin
	}

	scanner := bufio.NewScanner(inputReader)
	for scanner.Scan() {
		encoded := rot13(scanner.Text())
		fmt.Println(encoded)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}

func main() {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		processStdin()
	} else {
		processInput()
	}
}
