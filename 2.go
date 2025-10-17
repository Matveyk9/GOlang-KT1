package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    n := len(s)
    for i := 0; i < n/2; i++ {
        if s[i] != s[n-1-i] {
            return false
        }
    }
    return true
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Введите строку: ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

    if isPalindrome(input) {
        fmt.Println("Строка является палиндромом.")
    } else {
        fmt.Println("Строка не является палиндромом.")
    }
}
