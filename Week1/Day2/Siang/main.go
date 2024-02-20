package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(pembagian(5, 2))
	fmt.Println(pangkat(25, 2))
	fmt.Println(factorial(10))
}

func pembagian(angka1, angka2 float64) float64 {
	return angka1 / angka2
}

func pangkat(angka1, angka2 float64) float64 {
	return math.Pow(angka1, angka2)
}

func factorial(angka int) int {
	if angka == 0 {
		return 1
	}
	return angka * factorial(angka-1)
}
