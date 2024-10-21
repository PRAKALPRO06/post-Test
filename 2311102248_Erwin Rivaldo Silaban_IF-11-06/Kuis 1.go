package main

import (
	"fmt"
	"strconv"
)

func main() {
	var erwin231102248 int
	fmt.Print("Masukkan bilangan bulat positif (>10): ")
	fmt.Scan(&erwin231102248)
	strBilangan := strconv.Itoa(erwin231102248)
	panjang := len(strBilangan)

	var kiri, kanan string
	if panjang%2 == 0 {
		kiri = strBilangan[:panjang/2]
		kanan = strBilangan[panjang/2:]
	} else {
		kiri = strBilangan[:(panjang/2)+1]
		kanan = strBilangan[(panjang/2)+1:]
	}
	bilanganKiri, _ := strconv.Atoi(kiri)
	bilanganKanan, _ := strconv.Atoi(kanan)

	fmt.Println("Bilangan 1:", bilanganKiri)
	fmt.Println("Bilangan 2:", bilanganKanan)
	fmt.Println("Hasil penjumlahan:", bilanganKiri+bilanganKanan)
}
