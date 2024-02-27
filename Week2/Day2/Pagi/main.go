package main

import (
	"fmt"
	"runtime"
	"time"
)

func firstProcess(num int) {
	fmt.Println("First Starting")
	for i := 1; i < num; i++ {
		fmt.Println("i = ", i)
	}
}

func secondProcess(num int) {
	fmt.Println("Second Starting")
	for i := 1; i < num; i++ {
		fmt.Println("i = ", i)
	}
}

func main() {
	go firstProcess(8)
	secondProcess(8)
	time.Sleep(5 * time.Second)
	fmt.Println("Num of goroutine", runtime.NumGoroutine())
}
