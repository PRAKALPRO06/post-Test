package main

import "fmt"

func jumlahPertemuan(x, y int) int {
    pertemuan := 0
    for hari := 1; hari <= 365; hari++ {
        if hari % x == 0 && hari % y != 0 {
            pertemuan++
        }
    }
    return pertemuan
}

func main() {
	//2311102251_Marsep Trianto Pakondo
    var x, y int

    fmt.Print("Masukkan nilai X: ")
    fmt.Scan(&x)
    fmt.Print("Masukkan nilai Y: ")
    fmt.Scan(&y)

    hasil := jumlahPertemuan(x, y)
    fmt.Printf("Jumlah pertemuan dalam setahun: %d\n", hasil)
}
