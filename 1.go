package main

import (
    "fmt"
)

func main() {
    var num1, num2 float64
    var op string

    fmt.Print("Введите первое число: ")
    fmt.Scan(&num1)

    fmt.Print("Введите второе число: ")
    fmt.Scan(&num2)

    fmt.Print("Выберите операцию (+, -, *, /): ")
    fmt.Scan(&op)

    switch op {
    case "+":
        fmt.Printf("Результат: %.2f\n", num1+num2)
    case "-":
        fmt.Printf("Результат: %.2f\n", num1-num2)
    case "*":
        fmt.Printf("Результат: %.2f\n", num1*num2)
    case "/":
        if num2 == 0 {
            fmt.Println("Ошибка: деление на ноль невозможно!")
        } else {
            fmt.Printf("Результат: %.2f\n", num1/num2)
        }
    default:
        fmt.Println("Ошибка: неизвестная операция.")
    }
}
