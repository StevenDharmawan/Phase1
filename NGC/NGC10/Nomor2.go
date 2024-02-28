package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	defer func() {
		if message := recover(); message != nil {
			fmt.Println(message)
		}
	}()
	scanner := bufio.NewScanner(os.Stdin)
	operator := inputUser("Pilih operasi aritmatika:\n>Penjumlahan (+) a\n>Pengurangan (-) b\n> Perkalian (*) c\n> Pembagian (/) d", scanner)
	angka1 := isInteger(inputUser("Masukkan angka 1: ", scanner))
	angka2 := isInteger(inputUser("Masukkan angka 2: ", scanner))
	chooseOperator(operator, angka1, angka2)

}

func inputUser(text string, scanner *bufio.Scanner) string {
	fmt.Println(text)
	scanner.Scan()
	return scanner.Text()
}

func isInteger(text string) int {
	i, err := strconv.Atoi(text)
	if err != nil {
		panic("Input bukan Integer")
	}
	return i
}

func chooseOperator(text string, angka1 int, angka2 int) {
	switch text {
	case "a":
		fmt.Printf("hasil penjumlahan %d dan %d adalah %d\n", angka1, angka2, angka1+angka2)
	case "b":
		fmt.Printf("hasil pengurangan %d dan %d adalah %d\n", angka1, angka2, angka1-angka2)
	case "c":
		fmt.Printf("hasil perkalian %d dan %d adalah %d\n", angka1, angka2, angka1*angka2)
	case "d":
		if angka2 == 0 {
			panic("Tidak dapat melakukan pembagian 0")
		}
		fmt.Printf("hasil pembagian %d dan %d adalah %d\n", angka1, angka2, angka1/angka2)
	default:
		panic("Input operator tidak valid")
	}
}
