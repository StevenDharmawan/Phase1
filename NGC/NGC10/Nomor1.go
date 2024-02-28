package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func stringChecker(input string) {
	for _, character := range input {
		if !unicode.IsLetter(character) {
			panic("Input tidak valid")
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	defer func() {
		if message := recover(); message != nil {
			fmt.Println(message)
		}
	}()
	kata1 := userInput("Masukkan kata pertama: ", scanner)
	stringChecker(kata1)

	kata2 := userInput("Masukkan kata kedua: ", scanner)
	stringChecker(kata2)

	if len(kata1) != len(kata2) {
		fmt.Printf("%s & %s bukan merupakan anagram", kata1, kata2)
		return
	}
	isAnagram(kata1, kata2)
}

func userInput(text string, scanner *bufio.Scanner) string {
	fmt.Println(text)
	scanner.Scan()
	return scanner.Text()
}

func isAnagram(text1 string, text2 string) {
	slice := make([]int, 26)
	for _, value := range text1 {
		slice[value-'a']++
	}
	for _, value := range text2 {
		slice[value-'a']--

		if slice[value-'a'] < 0 {
			fmt.Printf("%s & %s bukan merupakan anagram", text1, text2)
			return
		}
	}
	fmt.Printf("%s & %s merupakan anagram", text1, text2)
}
