package main

import (
	"fmt"
	"math"
)

func main() {
	// Soal 1a
	soal1a(5)
	// Soal 1b
	soal1b(5, "Ivan Setiawan")
	// Soal 2
	fmt.Println(hitungBonus("A", 10000))
	fmt.Println(hitungBonus("B", 10000))
	fmt.Println(hitungBonus("C", 10000))
	fmt.Println(hitungBonus("D", 10000))
	// Soal 3
	nilaiSiswa := []int{85, 60, 78, 92, 45, 73}
	evaluasiKinerjaSiswa(nilaiSiswa)
}

// Soal 1a
func soal1a(n int) {
	n1 := int(math.Round(float64(n) / 2))
	n2 := n - n1

	for i := 0; i < n1*2; i += 2 {
		for j := n1 * 2; j > i+2; j -= 2 {
			fmt.Printf(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

	for i := 0; i < n2*2; i += 2 {
		if n1 != n2 {
			for j := 0; j <= i; j += 2 {
				fmt.Printf(" ")
			}
		} else {
			for j := 0; j < i; j += 2 {
				fmt.Printf(" ")
			}
		}
		for j := n2 * 2; j > i+1; j-- {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

// Soal 1b
func soal1b(n int, name string) {
	nameIndex := 0
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if nameIndex >= len(name) {
				fmt.Printf("* ")
			} else {
				fmt.Printf("%s ", name[nameIndex:nameIndex+1])
			}
			nameIndex++
		}
		fmt.Println()
	}
}

// Nomor 2
func hitungBonus(performa string, gaji float64) float64 {
	switch performa {
	case "A":
		return gaji * 1.2
	case "B":
		return gaji * 1.1
	case "C":
		return gaji * 1.05
	default:
		return gaji
	}
}

// Nomor 3
func evaluasiKinerjaSiswa(nilaiSiswa []int) {
	for index, value := range nilaiSiswa {
		index++
		if value <= 100 && value >= 85 {
			fmt.Printf("Siswa %d mendapatkan predikat A\n", index)
		} else if value >= 70 {
			fmt.Printf("Siswa %d mendapatkan predikat B\n", index)
		} else if value >= 50 {
			fmt.Printf("Siswa %d mendapatkan predikat C\n", index)
		} else if value >= 0 {
			fmt.Printf("Siswa %d mendapatkan predikat D\n", index)
		} else {
			fmt.Println("Nilai tidak valid")
		}
	}
}
