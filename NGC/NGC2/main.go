package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	//Nomor 1
	fmt.Print("Masukan nama: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	random := rand.Intn(100-0) + 0

	if random <= 100 && random > 80 {
		fmt.Printf("Selamat %s anda sangat beruntung\n", scanner.Text())
	} else if random > 60 {
		fmt.Printf("Selamat %s anda beruntung\n", scanner.Text())
	} else if random > 40 {
		fmt.Printf("Mohon maaf %s anda kurang beruntung\n", scanner.Text())
	} else {
		fmt.Printf("Mohon maaf %s anda sial\n", scanner.Text())
	}

	//Nomor 2
	fmt.Print("Masukan nama: ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Print("Masukan umur: ")
	scanner.Scan()
	age, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println("Input tidak valid")
	} else {
		if age <= 100 && age > 18 {
			fmt.Printf("Silahkan masuk %s\n", name)
		} else if age <= 18 && age > 0 {
			fmt.Printf("Dilarang masuk, minimal umur 19 %s\n", name)
		} else {
			fmt.Println("Input tidak valid")
		}
	}
}
