package main

import (
	"fmt"
	"os"
	"time"
)

// Уровни логирования
const (
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
)

func logMessage(level, message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logLine := fmt.Sprintf("[%s] [%s] %s\n", timestamp, level, message)

	fmt.Print(logLine)

	// файл для ошибок
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Ошибка открытия файла для логов: %v\n", err)
		return
	}
	defer file.Close()

	// Запись логов в файл
	if _, err := file.WriteString(logLine); err != nil {
		fmt.Printf("Ошибка записи в файл: %v\n", err)
	}
}

func main() {
	logMessage(INFO, "Сервис запущен")
	logMessage(WARN, "Тестовое предупреждение")
	logMessage(ERROR, "Произошла ошибка!")
}
