package main

import (
	"fmt"
	"time"
)

func sumSquare(channel <-chan []int, result chan<- int) {
	total := 0
	for _, value := range <-channel {
		total += value
	}
	total *= total
	result <- total
	time.Sleep(3 * time.Second)
}

func squareSum(channel <-chan []int, result chan<- int) {
	total := 0
	for _, value := range <-channel {
		total += value * value
	}
	result <- total
	time.Sleep(3 * time.Second)
}

func main() {
	channel := make(chan []int)
	result := make(chan int)
	var intSlice []int
	for i := 1; i <= 100; i++ {
		intSlice = append(intSlice, i)
	}
	go func() {
		channel <- intSlice
	}()
	go sumSquare(channel, result)
	fmt.Println("SumSquare =", <-result)
	go func() {
		channel <- intSlice
	}()
	go squareSum(channel, result)
	fmt.Println("SquareSum =", <-result)
}
