package main

import (
	"fmt"
)

func main() {
	//Nomor 1
	slice1 := []float64{1, 5, 7, 8, 10, 24, 33}
	slice2 := []float64{1.1, 5.4, 6.7, 9.2, 11.3, 25.2, 33.1}
	fmt.Println(average(slice1))
	fmt.Println(average(slice2))
	fmt.Println(sum(slice1))
	fmt.Println(sum(slice2))
	fmt.Println(median(slice1))
	fmt.Println(median(slice2))

	//Nomor 2
	kata1 := "katak"
	kata2 := "Beras"
	fmt.Println(palindrome(kata1))
	fmt.Println(palindrome(kata2))

	//Nomor 3
	checker1 := "xoxoxo"
	checker2 := "xxxoxx"
	fmt.Println(checker(checker1))
	fmt.Println(checker(checker2))

	//Nomor 4
	sliceInt1 := []int{500, 3, 2, 6, 1, 5, 7, 8, 10, 24, 33}
	fmt.Println(BubbleSort(sliceInt1))

	//Nomor 5
	Asterisk1(5)

	//Nomor 6
	Asterisk2(5)

	//Nomor 7
	Asterisk3(5)

	//Nomor 8
	Asterisk4(5)
}

// Nomor 1
func average(slice []float64) float64 {
	var avg float64
	for _, value := range slice {
		avg += value
	}
	avg /= float64(len(slice))
	return avg
}

// Nomor 1
func sum(slice []float64) float64 {
	var sum float64
	for _, value := range slice {
		sum += value
	}
	return sum
}

// Nomor 1
func median(slice []float64) float64 {
	length := len(slice)
	mid := length / 2
	if length%2 == 0 {
		return (slice[mid] + slice[mid+1]) / float64(2)
	} else {
		return slice[mid]
	}
}

// Nomor 2
func palindrome(kata string) bool {
	for i := 0; i < len(kata)/2; i++ {
		if kata[i] != kata[len(kata)-1-i] {
			return false
		}
	}
	return true
}

// Nomor 3
func checker(input string) bool {
	var sumX, sumO int
	for _, value := range input {
		if value == 'x' {
			sumX++
		} else if value == 'o' {
			sumO++
		}
	}
	if sumX == sumO {
		return true
	}
	return false
}

// Nomor 4
func BubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

// Nomor 5
func Asterisk1(row int) {
	for i := 0; i < row; i++ {
		fmt.Println("*")
	}
}

// Nomor 6
func Asterisk2(row int) {
	for i := 0; i < row; i++ {
		for j := 0; j < row; j++ {
			fmt.Printf("*")
		}
		fmt.Println("")
	}
}

// Nomor 7
func Asterisk3(row int) {
	for i := 0; i < row; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Println("")
	}
}

func Asterisk4(row int) {
	for i := row; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Printf("*")
		}
		fmt.Println("")
	}
}
