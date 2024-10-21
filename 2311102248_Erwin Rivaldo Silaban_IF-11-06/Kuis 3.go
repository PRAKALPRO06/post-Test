package main

import "fmt"

func erwin248(n, m int) int {
	if m == 0 {
		return 0
	}
	return n + erwin248(n, m-1)
}

func main() {
	var n, m int

	fmt.Print("Masukkan bilangan n: ")
	fmt.Scan(&n)

	fmt.Print("Masukkan bilangan m: ")
	fmt.Scan(&m)

	hasil := erwin248(n, m)
	fmt.Printf("Hasil dari %d x %d = %d\n", n, m, hasil)
}
