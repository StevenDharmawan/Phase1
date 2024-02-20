package main

import (
	"fmt"
)

func main() {
	grade(120)
	grade(90)
	grade(80)
	grade(70)
	grade(60)
	grade(0)
	grade(-20)
}

func grade(score int) {
	if score <= 100 && score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else if score >= 60 {
		fmt.Println("D")
	} else if score < 60 && score >= 0 {
		fmt.Println("E")
	} else {
		fmt.Println("Invalid Score")
	}

	for i := 1; i <= 15; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "Hello World")
		} else if i%3 == 0 {
			fmt.Println(i, "Hello")
		} else if i%5 == 0 {
			fmt.Println(i, "World")
		} else {
			fmt.Println(i)
		}
	}
}
