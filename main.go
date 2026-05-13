package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. Проверка на правильный запуск
	if len(os.Args) != 3 {
		fmt.Println("❌ Ошибка: Неверное количество аргументов.")
		fmt.Println("💡 Как запустить: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// 2. Проверка на ошибку чтения (например, файла не существует)
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("❌ Ошибка: Не удалось прочитать файл '%s'. Проверь, существует ли он!\n", inputFile)
		return
	}

	// Отправляем текст на обработку (в process.go)
	modifiedText := processText(string(data))
	
	// 3. Проверка на ошибку записи (например, нет прав или места на диске)
	err = os.WriteFile(outputFile, []byte(modifiedText), 0o644)
	if err != nil {
		fmt.Printf("❌ Ошибка: Не удалось сохранить результат в файл '%s'.\n", outputFile)
		return
	}

	// 4. Если код дошел сюда, значит всё прошло идеально!
	fmt.Printf("✅ Готово! Можешь проверять файл '%s'\n", outputFile)
} 
