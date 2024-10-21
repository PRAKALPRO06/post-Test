package main

import "fmt"

func diskon(jam, menit, voucher int, member bool) {
	var tarif, diskon float64
	for jam > 0 {
		if member {
			tarif += 3500
		} else {
			tarif += 5000
		}
		jam--
	}

	if menit > 0 {
		if menit <= 10 {
		} else if member {
			tarif += 3500 
		} else {
			tarif += 5000
		}
	}

	if voucher >= 10000 {
		diskon = tarif * 0.10 
		tarif = tarif - diskon
	}

	fmt.Printf("Biaya sewa setelah diskon: %.2f\n", tarif)
}

func main() {
	//2311102251_Marsep Trianto Pakondo
	var jam, menit, voucher int
	var member bool

	fmt.Print("Masukkan durasi (jam): ")
	fmt.Scan(&jam)
	fmt.Print("Masukkan durasi (menit): ")
	fmt.Scan(&menit)
	fmt.Print("Apakah member? (true/false): ")
	fmt.Scan(&member)
	fmt.Print("Masukkan nomor voucher (jika ada): ")
	fmt.Scan(&voucher)

	diskon(jam, menit, voucher, member)
}
