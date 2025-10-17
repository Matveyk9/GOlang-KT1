package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    const n = 10
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        arr[i] = rand.Intn(100) + 1
    }

    fmt.Println("Массив:", arr)

    sum := 0
    min := arr[0]
    max := arr[0]
    for _, v := range arr {
        sum += v
        if v < min {
            min = v
        }
        if v > max {
            max = v
        }
    }
    avg := float64(sum) / float64(n)

    fmt.Println("Сумма:", sum)
    fmt.Println("Минимум:", min)
    fmt.Println("Максимум:", max)
    fmt.Printf("Среднее: %.2f\n", avg)
}
