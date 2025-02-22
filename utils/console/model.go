package utils_console

import (
	colors "bookmarkapp/utils/colors"
	"fmt"
	"log"
	"time"
)

// Println дублирование fmt.Println для консистенции модуля
func Println(value string) {
	fmt.Println(value)
}

// PrintlnEmptyLine добавляет отступ
func PrintlnEmptyLine() {
	fmt.Println("")
}

// PrintlnWithColor выводит текст в консоль с указанным цветом.
func PrintlnWithColor(value string, color colors.Color) {
	colorCode := colors.GetColorCode(color)
	fmt.Println(colorCode, value, "\033[0m")
}

// Sleep приостанавливает выполнение программы на указанное количество секунд.
func Sleep(seconds float32) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

// InputValue определяет типы, которые могут быть использованы для ввода.
type InputValue interface {
	int | string
}

// UserInputValue запрашивает ввод у пользователя и возвращает значение указанного типа.
func UserInputValue[V InputValue](text string, color colors.Color) V {
	var userInput V
	
	PrintlnWithColor(text, color)
	fmt.Scan(&userInput)

	return userInput
}

func Log(message string, err error){
	log.Println("Error getting bookmarks:", err)
}