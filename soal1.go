package main

import "fmt"

func main() {
    var nim = 2311102026
    fmt.Println("NIM:", nim)

    var a, b int
    fmt.Print("Masukkan nilai a: ")
    fmt.Scan(&a)
    fmt.Print("Masukkan nilai b: ")
    fmt.Scan(&b)

    countGanjil := 0

    for i := a; i <= b; i++ {
        if i%2 != 0 { 
            countGanjil++
        }
    }

