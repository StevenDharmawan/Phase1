package main

import (
	"bufio"
	"fmt"
)

func InputUser(text string, scanner *bufio.Scanner) string {
	fmt.Println(text)
	scanner.Scan()
	return scanner.Text()
}

func isInteger(text string) {

}

func chooseOperator(text string, angka1 int, angka2 int) {
	switch text {
	case "a":
		fmt.Println("one")
	case "b":
		fmt.Println("two")
	case "c":
		fmt.Println("three")
	case "d":
		fmt.Println("four")
	default:
		panic("Input tidak valid")
	}
}
