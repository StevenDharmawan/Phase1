package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Area          float64
	Circumference float64
}

func circleArea(channel <-chan float64, result chan<- float64) {
	result <- math.Pi * math.Pow(<-channel/2, 2)
}

func circleCircumference(channel <-chan float64, result chan<- float64) {
	result <- math.Pi * <-channel
}

func main() {
	inputChannel := make(chan float64)
	result := make(chan float64)
	defer close(inputChannel)
	go func() {
		inputChannel <- 10
	}()
	go circleArea(inputChannel, result)
	fmt.Printf("Luas Lingkaran: %.2f\n", <-result)
	go func() {
		inputChannel <- 10
	}()
	go circleCircumference(inputChannel, result)
	fmt.Printf("Keliling Lingkaran: %.2f\n", <-result)
}
