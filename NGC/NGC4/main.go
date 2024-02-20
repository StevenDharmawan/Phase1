package main

import (
	"fmt"
)

func main() {
	//Nomor 1
	fmt.Println(alayGen("hello", "world", "zzz"))

	//Nomor 2
	fmt.Println(fibonacci(9))
}

// Nomor 1
func alayGen(kata ...string) string {
	newKata := ""
	for i, value := range kata {
		if i != len(kata)-1 {
			newKata += value
			newKata += " "
		} else {
			newKata += value
		}
	}
	return ganti(newKata)
}

// Nomor 1
func ganti(kata string) string {
	newKata := ""
	for _, value := range kata {
		if value == 'a' {
			newKata += "4"
		} else if value == 'e' {
			newKata += "3"
		} else if value == 'i' {
			newKata += "!"
		} else if value == 'l' {
			newKata += "1"
		} else if value == 'n' {
			newKata += "N"
		} else if value == 's' {
			newKata += "5"
		} else if value == 'x' {
			newKata += "*"
		} else {
			newKata += string(value)
		}
	}
	return newKata
}

// Nomor 2
func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
