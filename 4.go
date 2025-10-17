package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sumSlice(arr []int, ch chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	ch <- sum 
}

func main() {
	rand.Seed(time.Now().UnixNano())

	size := 10
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(100) + 1
	}

	fmt.Println("Массив:", arr)

	ch := make(chan int)

	mid := len(arr) / 2
	go sumSlice(arr[:mid], ch)
	go sumSlice(arr[mid:], ch)

	sum1 := <-ch
	sum2 := <-ch

	total := sum1 + sum2
	fmt.Println("Сумма элементов массива:", total)
}
