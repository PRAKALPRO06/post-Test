package main

import "fmt"

func jumlahKelipatan4(nim int, total int) int {
    if nim < 0 {
        return total
    }

    if nim % 4 == 0 && nim > 0 {
        total += nim
    }

    var next int
    fmt.Scan(&next)

    return jumlahKelipatan4(next, total)
}

func main() {
    fmt.Println("Masukkan bilangan (negatif untuk berhenti):")

    var nim int
    fmt.Scan(&nim)

    hasil := jumlahKelipatan4(nim, 0)

    fmt.Println("Jumlah bilangan kelipatan 4:", hasil)
}