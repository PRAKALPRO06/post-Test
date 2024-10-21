package main

import "fmt"


func main()  {
	var rombongan int
	var jumlahMenu_2311102010,orang int
	var sisa int
	const hargaBiasa = 10000
	const hargaPerMenu = 2500
	const hargaMaks = 50000
	
	fmt.Print("Masukan Jumlah Rombongan : ")
	fmt.Scan(&rombongan)

	for i:= 1; i<= rombongan; i++ {
		fmt.Printf("Masukan jumlah menu, jumlah orang, dan status sisa makanan (0 untuk tidak, 1 untuk iya)", i)
		fmt.Scan(&jumlahMenu_2311102010, &orang, &sisa)
		totalHarga:= hargaBiasa

		if jumlahMenu_2311102010 > 3 {
			totalHarga += (jumlahMenu_2311102010 - 3) * hargaPerMenu
		}

		totalHarga *= orang
		if sisa == 1 {
			totalHarga += orang * totalHarga
		}

		fmt.Print("Total Biaya untuk rombongan %d : Rp %d\n", i, totalHarga)
	}
}