package main

import "fmt"

type Result struct {
	Operation string
	Value     float64
}

func penjumlahan(chInput chan [2]float64, chOutput chan float64) {
	data := <-chInput
	chOutput <- data[0] + data[1]
}

func pengurangan(chInput chan [2]float64, chOutput chan float64) {
	data := <-chInput
	chOutput <- data[0] + data[1]
}

func perkalian(chInput chan [2]float64, chOutput chan float64) {
	data := <-chInput
	chOutput <- data[0] + data[1]
}

func pembagian(chInput chan [2]float64, chOutput chan float64) {
	data := <-chInput
	chOutput <- data[0] + data[1]
}

func main() {
	chInput := make(chan [2]float64)
	chOutput := make(chan float64)

	go penjumlahan(chInput, chOutput) // Mulai goroutine penjumlahan
	chInput <- [2]float64{5.5, 6.5}   // Kirim data ke channel chInput

	fmt.Println(<-chOutput) // Tunggu dan cetak hasil dari channel chOutput

	channel := make(chan int)
	go func() {
		channel <- 10
	}()
	data := <-channel

	fmt.Println(data)
}
