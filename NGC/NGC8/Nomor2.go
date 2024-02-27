package main

import (
	"fmt"
	"time"
)

func sumSquare(channel <-chan []uint64, result chan<- uint64) {
	total := uint64(0)
	for _, value := range <-channel {
		total += value
	}
	total *= total
	result <- total
	time.Sleep(3 * time.Second)
}

func squareSum(channel <-chan []uint64, result chan<- uint64) {
	total := uint64(0)
	for _, value := range <-channel {
		total *= total
		total += value
	}
	result <- total
	time.Sleep(3 * time.Second)
}

func main() {
	channel := make(chan []uint64)
	result := make(chan uint64)
	var intSlice []uint64
	for i := uint64(1); i <= 100; i++ {
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
