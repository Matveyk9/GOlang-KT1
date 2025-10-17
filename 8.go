package main

import (
	"fmt"
)

func sumSlice(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func runTests() {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Пустой массив", []int{}, 0},
		{"Положительные числа", []int{1, 2, 3, 4, 5}, 15},
		{"С отрицательными числами", []int{-1, -2, 3}, 0},
		{"Один элемент", []int{42}, 42},
	}

	for _, tt := range tests {
		result := sumSlice(tt.input)
		if result != tt.expected {
			fmt.Printf("❌ %s: ожидалось %d, получили %d\n", tt.name, tt.expected, result)
		} else {
			fmt.Printf("✅ %s: прошло успешно\n", tt.name)
		}
	}
}

func main() {

	runTests()
}
