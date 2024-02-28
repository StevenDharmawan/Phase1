package main

import (
	"fmt"
	"math/rand"
)

func main() {
	random := rand.Intn(100) + 1
	_ = random
	channel := make(chan int)
	go func() {
		channel <- 10
	}()
	result := Factorial(channel)
	fmt.Print(<-result)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func Factorial(channel <-chan int) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		result <- factorial(<-channel)
	}()
	return result
}
