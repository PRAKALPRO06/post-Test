// LAILATUR RAHMAH
// 2311102177
// IF 11 06

package main

import (
	"fmt"
)

func hitungPertemuanAgen_231102177(x, y int) int {
	jumlahPertemuan := 0

	for hari := 1; hari <= 365; hari++ {

		if hari%x == 0 && hari%y != 0 {
			jumlahPertemuan++
		}
	}

	return jumlahPertemuan
}

func main() {
	var x, y int

	fmt.Print("Masukkan nilai x: ")
	fmt.Scanln(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scanln(&y)

	jumlah := hitungPertemuanAgen_231102177(x, y)

	fmt.Printf("Jumlah pertemuan dalam setahun: %d\n", jumlah)
}
