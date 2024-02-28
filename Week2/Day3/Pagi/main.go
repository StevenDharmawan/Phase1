package main

import (
	"fmt"
	"os"
)

func main() {

}

func deferGolang() {
	defer fmt.Println("Hola hola 1")
	fmt.Println("Hello World 1")

	defer fmt.Println("Hola hola 2")
	fmt.Println("Hello World 2")

	defer fmt.Println("Hola hola 3")
	fmt.Println("Hello World 3")
}

func implementationDeferExit() {
	defer cleanUp()
	fmt.Println("Doing some critical work")

	if true {
		fmt.Println("Ada error nih!")
		os.Exit(1)
	}

	fmt.Println("Critical work done successfully")
}

func cleanUp() {
	fmt.Println("CleanUp: Program already clean")
}
