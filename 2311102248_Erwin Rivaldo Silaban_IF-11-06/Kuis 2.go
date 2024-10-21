package main

import (
	"fmt"
)

func allSameDigits(card string) bool {
	for i := 1; i < len(card); i++ {
		if card[i] != card[0] {
			return false
		}
	}
	return true
}

func allDifferentDigits(card string) bool {
	digitCount := make(map[rune]bool)
	for _, digit := range card {
		if digitCount[digit] {
			return false
		}
		digitCount[digit] = true
	}
	return true
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&n)

	hadiahA, hadiahB, hadiahC := 0, 0, 0

	for i := 1; i <= n; i++ {
		var card string
		fmt.Printf("Masukkan nomor kartu peserta ke-%d: ", i)
		fmt.Scan(&card)

		if allSameDigits(card) {
			fmt.Println("Hadiah A")
			hadiahA++
		} else if allDifferentDigits(card) {
			fmt.Println("Hadiah B")
			hadiahB++
		} else {
			fmt.Println("Hadiah C")
			hadiahC++
		}
	}

	fmt.Printf("Jumlah yang memperoleh Hadiah A: %d\n", hadiahA)
	fmt.Printf("Jumlah yang memperoleh Hadiah B: %d\n", hadiahB)
	fmt.Printf("Jumlah yang memperoleh Hadiah C: %d\n", hadiahC)
}
