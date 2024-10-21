// LAILATUR RAHMAH
// 2311102177
// IF 11 06

package main

import (
	"fmt"
)

func hitungBiayaSewa(durasiJam_2311102177, durasiMenit int, isMember bool, nomorVoucher string) float64 {
	tarif := 5000.0
	if isMember {
		tarif = 3500.0
	}

	if durasiMenit >= 10 || durasiJam_2311102177 == 0 {
		durasiJam_2311102177 += 1
	}

	biaya := float64(durasiJam_2311102177) * tarif

	if durasiJam_2311102177 > 3 && (len(nomorVoucher) == 5 || len(nomorVoucher) == 6) {
		biaya *= 0.9
	}

	return biaya
}

func main() {
	var durasiJam_2311102177 int
	var durasiMenit int
	var isMember bool
	var nomorVoucher string

	fmt.Print("Masukkan durasi (jam): ")
	fmt.Scanln(&durasiJam_2311102177)
	fmt.Print("Masukkan durasi (menit): ")
	fmt.Scanln(&durasiMenit)
	fmt.Print("Apakah member? (true/false): ")
	fmt.Scanln(&isMember)
	fmt.Print("Masukkan nomor voucher (jika ada): ")
	fmt.Scanln(&nomorVoucher)

	biaya := hitungBiayaSewa(durasiJam_2311102177, durasiMenit, isMember, nomorVoucher)

	fmt.Printf("Biaya sewa setelah diskon (jika memenuhi syarat): Rp %.2f\n", biaya)
}
