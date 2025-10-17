package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "strings"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    target := rand.Intn(100) + 1

    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("Угадайте число от 1 до 100: ")
        line, _ := reader.ReadString('\n')
        line = strings.TrimSpace(line)
        guess, err := strconv.Atoi(line)
        if err != nil {
            fmt.Println("Ошибка: введите число.")
            continue
        }

        if guess < target {
            fmt.Println("Больше.")
        } else if guess > target {
            fmt.Println("Меньше.")
        } else {
            fmt.Println("Поздравляю! Вы угадали число!")
            break
        }
    }
}
