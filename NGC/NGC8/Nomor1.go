package main

import (
	"fmt"
	"time"
)

func fizzBuzz(channel <-chan []int, outputChannel chan<- int) {
	total, ganjil, genap := 0, 0, 0
	for _, value := range <-channel {
		if value%3 == 0 && value%5 == 0 {
			fmt.Println(value, "FizzBuzz")
		} else if value%3 == 0 {
			fmt.Println(value, "Fizz")
		} else if value%5 == 0 {
			fmt.Println(value, "Buzz")
		}
		if value%2 == 0 {
			genap += value
		} else {
			ganjil += value
		}
		total += value
	}
	outputChannel <- genap
	outputChannel <- ganjil
	fmt.Println(total)
	time.Sleep(3 * time.Second)
}

func main() {
	channel := make(chan []int)
	outputChannel := make(chan int, 2)
	var intSlice []int
	for i := 0; i < 100; i++ {
		intSlice = append(intSlice, i)
	}
	go func() {
		channel <- intSlice
	}()
	go fizzBuzz(channel, outputChannel)
	fmt.Println("Genap =", <-outputChannel)
	fmt.Println("Ganjil =", <-outputChannel)
}
