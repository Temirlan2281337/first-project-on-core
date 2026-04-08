package main

import (
	"fmt"
	"os"
	"strings"
)

func processText(input string) string {
	words := strings.Fields(input)
	var result []string

	for _, word := range words {
		if word == "(up)" {
		} else {
			result = append(result, word)
		}
	}
	return strings.Join(result, " ")
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("error")
		return

	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	text := string(data)

	fmt.Println(text)
	err = os.WriteFile(outputFile, []byte(text), 0o644)
	if err != nil {
		fmt.Println("error", err)
		return
	}
}
