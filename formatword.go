package main

import (
	"strings"
)

// formatWord меняет регистр слова
func formatWord(word string, action string) string {
	switch action {
	case "up":
		return strings.ToUpper(word)
	case "low":
		return strings.ToLower(word)
	case "cap":
		if len(word) > 0 {
			return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	return word
}
