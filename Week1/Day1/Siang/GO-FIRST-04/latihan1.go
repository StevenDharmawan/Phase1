package main

import (
	"fmt"
)

func main() {
	i := 21
	j := true
	russia := "Я"
	base10 := 21
	base16 := 15
	float := 123.456
	fmt.Println(i)
	fmt.Printf("%T\n", i)
	fmt.Println("%")
	fmt.Println(j, "\n")
	fmt.Printf("%b", []rune(russia)[0])
	fmt.Println("\n?")
	fmt.Println(base10)
	fmt.Printf("%o\n", base10)
	fmt.Printf("%x\n", base16)
	fmt.Printf("%X\n", base16)
	fmt.Printf("%U", []rune(russia)[0])
	fmt.Printf("\n%.6f\n", float)
	fmt.Printf("%E\n", float)

	// %b	base 2
	// %c	the character represented by the corresponding Unicode code point
	// %d	base 10
	// %o	base 8
	// %O	base 8 with 0o prefix
	// %q	a single-quoted character literal safely escaped with Go syntax.
	// %x	base 16, with lower-case letters for a-f
	// %X	base 16, with upper-case letters for A-F
	// %U	Unicode format: U+1234; same as "U+%04X"

}
