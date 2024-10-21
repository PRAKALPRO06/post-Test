package main

import "fmt"


func hitungKelipatan4_2311102010(bilangan int) int {
	if bilangan < 0 {
	return 0
}

	var next int
	fmt.Scan(&next)

	if bilangan > 0 && bilangan%4 == 0 {
		return bilangan + hitungKelipatan4_2311102010(next)
	}

	return hitungKelipatan4_2311102010(next)
	}

	func main() {
	fmt.Println("Masukkan bilangan (negatif untuk berhenti):")

	var bilangan int
	fmt.Scan(&bilangan)

	hasil := hitungKelipatan4_2311102010(bilangan)

	fmt.Println("Jumlah bilangan kelipatan4:",hasil)
}