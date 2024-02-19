package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Nomor 1
	var myNum int32 = 50
	fmt.Println(myNum)

	var myNum2 float32 = 51
	fmt.Println(myNum2)

	var myNumStr string = "50"
	fmt.Println(myNumStr)

	// Nomor 2
	var x int32 = 5
	var y int32 = 10

	z := x + y

	fmt.Println(z)

	// Nomor 3
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukan nama: ")
	text, _ := reader.ReadString('\n')
	fmt.Println("Hello", text)

	// Nomor 4

	peoples := []string{"Walt", "Jesse", "Skyler", "Saul"}
	fmt.Println(peoples)
	peoples = append(peoples, "Hank", "Marie")
	fmt.Println(peoples)

	// Nomor 5
	var array [3]map[string]string
	array[0] = map[string]string{"Name": "Hank", "Gender": "M"}
	array[1] = map[string]string{"Name": "Heisenber", "Gender": "M"}
	array[2] = map[string]string{"Name": "Skyler", "Gender": "F"}
	fmt.Println(array)

	for _, value := range array {
		if value["Gender"] == "M" {
			value["Name"] = "Mr " + value["Name"]
		} else if value["Gender"] == "F" {
			value["Name"] = "Mrs " + value["Name"]
		}
	}

	fmt.Println(array)
}
