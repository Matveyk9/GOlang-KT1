package main

import (
    "fmt"
    "math/rand"
    "time"
)

func selectionSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        minIdx := i
        for j := i + 1; j < n; j++ {
            if arr[j] < arr[minIdx] {
                minIdx = j
            }
        }
        arr[i], arr[minIdx] = arr[minIdx], arr[i]
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())

    size := 10
    arr := make([]int, size)
    for i := 0; i < size; i++ {
        arr[i] = rand.Intn(100) 
    }

    fmt.Println("Исходный массив:", arr)


    selectionSort(arr)

    fmt.Println("Отсортированный массив:", arr)
}
