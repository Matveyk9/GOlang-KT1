package main

import (
	"encoding/json"
	"fmt"
)

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Book struct {
	Title     string  `json:"title"`
	Author    Author  `json:"author"`
	Price     float64 `json:"price"`
	Available bool    `json:"available"`
}

func (b Book) Info() string {
	return fmt.Sprintf("'%s' by %s (%.2f USD) - Available: %t", b.Title, b.Author.Name, b.Price, b.Available)
}

func main() {
	book := Book{
		Title: "Go Programming",
		Author: Author{
			Name:  "Alice Johnson",
			Email: "alice@example.com",
		},
		Price:     29.99,
		Available: true,
	}

	// Сериализация в JSON
	jsonData, err := json.MarshalIndent(book, "", "  ")
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		return
	}
	fmt.Println("Сериализованный JSON:")
	fmt.Println(string(jsonData))

	// Десериализация из JSON
	var newBook Book
	err = json.Unmarshal(jsonData, &newBook)
	if err != nil {
		fmt.Println("Ошибка десериализации:", err)
		return
	}

	fmt.Println("\nИнформация о книге после десериализации:")
	fmt.Println(newBook.Info())
}
