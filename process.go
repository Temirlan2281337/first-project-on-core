package main

import (
	"strconv"
	"strings"
)

// Главная функция: режет текст на строки и собирает обратно
func processText(input string) string {
	// Унифицируем переносы строк на случай, если файл создан в Windows (\r\n -> \n)
	input = strings.ReplaceAll(input, "\r\n", "\n")
	
	// Разделяем текст на массив отдельных строк
	lines := strings.Split(input, "\n")
	var resultLines []string

	for _, line := range lines {
		// Если строка пустая (просто Enter), оставляем её пустой
		if line == "" {
			resultLines = append(resultLines, "")
			continue
		}
		// Отправляем строку на обработку и добавляем в готовый список
		resultLines = append(resultLines, processLine(line))
	}

	// Склеиваем строки обратно, вставляя между ними перенос строки (\n)
	return strings.Join(resultLines, "\n")
}

// processLine делает всю ту же работу, что раньше, но только для одной строки
func processLine(input string) string {
	words := strings.Fields(input)
	var result []string

	// ШАГ 1: Обработка всех команд и систем счисления
	for i := 0; i < len(words); i++ {
		word := words[i]

		if word == "(hex)" {
			if len(result) > 0 {
				lastIdx := len(result) - 1
				if num, err := strconv.ParseInt(result[lastIdx], 16, 64); err == nil {
					result[lastIdx] = strconv.FormatInt(num, 10)
				}
			}
		} else if word == "(bin)" {
			if len(result) > 0 {
				lastIdx := len(result) - 1
				if num, err := strconv.ParseInt(result[lastIdx], 2, 64); err == nil {
					result[lastIdx] = strconv.FormatInt(num, 10)
				}
			}
		} else if word == "(up)" || word == "(low)" || word == "(cap)" {
			if len(result) > 0 {
				action := strings.Trim(word, "()")
				result[len(result)-1] = formatWord(result[len(result)-1], action)
			}
		} else if word == "(up," || word == "(low," || word == "(cap," {
			if i+1 < len(words) {
				nextWord := words[i+1]
				numStr := strings.TrimRight(nextWord, ")")
				
				if num, err := strconv.Atoi(numStr); err == nil {
					action := strings.Trim(word, "(,")
					for j := 1; j <= num; j++ {
						idx := len(result) - j
						if idx >= 0 {
							result[idx] = formatWord(result[idx], action)
						}
					}
				}
				i++ // Пропускаем число
			}
		} else {
			result = append(result, word)
		}
	}

	// ШАГ 2: Железобетонная обработка знаков препинания
	var withPunct []string
	for i := 0; i < len(result); i++ {
		w := result[i]
		
		isPunct := true
		for _, char := range w {
			if !strings.ContainsRune(".,!?:;", char) {
				isPunct = false
				break
			}
		}
		
		if isPunct && len(w) > 0 {
			if len(withPunct) > 0 {
				withPunct[len(withPunct)-1] += w
			} else {
				withPunct = append(withPunct, w)
			}
		} else {
			withPunct = append(withPunct, w)
		}
	}
	result = withPunct

	// ШАГ 3: Обработка одинарных кавычек
	var withQuotes []string
	openQuote := false
	for i := 0; i < len(result); i++ {
		if result[i] == "'" {
			if !openQuote {
				if i+1 < len(result) {
					result[i+1] = "'" + result[i+1]
					openQuote = true
				} else {
					withQuotes = append(withQuotes, "'")
				}
			} else {
				if len(withQuotes) > 0 {
					withQuotes[len(withQuotes)-1] += "'"
				} else {
					withQuotes = append(withQuotes, "'")
				}
				openQuote = false
			}
		} else {
			withQuotes = append(withQuotes, result[i])
		}
	}
	result = withQuotes

	// ШАГ 4: Умная замена a -> an
	for i := 0; i < len(result)-1; i++ {
		if result[i] == "a" || result[i] == "A" {
			nextWord := result[i+1]
			if len(nextWord) > 0 {
				firstChar := strings.ToLower(string(nextWord[0]))
				
				if firstChar == "'" && len(nextWord) > 1 {
					firstChar = strings.ToLower(string(nextWord[1]))
				}
				
				if firstChar == "a" || firstChar == "e" || firstChar == "i" || firstChar == "o" || firstChar == "u" || firstChar == "h" {
					result[i] += "n"
				}
			}
		}
	}

	return strings.Join(result, " ")
}