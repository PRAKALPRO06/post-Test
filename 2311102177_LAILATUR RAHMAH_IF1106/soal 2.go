// LAILATUR RAHMAH
// 2311102177
// IF 11 06

package main

import (
	"fmt"
)

func isPerfectNumber(n int) bool {
	sum_2311102177 := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum_2311102177 += i
		}
	}
	return sum_2311102177 == n
}

func perfectNumbersInRange(a, b int) {
	fmt.Printf("Perfect numbers antara %d dan %d: ", a, b)
	for i := a; i <= b; i++ {
		if isPerfectNumber(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

func main() {
	var a, b int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scanln(&b)

	perfectNumbersInRange(a, b)
}
