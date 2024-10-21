package main

import "fmt"

func perfectNumber(a, b int)  {
    for angka := a; angka <= b; angka++ {
        jumlahFaktor := 0
        
        for faktor := 1; faktor < angka; faktor++ {
            if angka % faktor == 0 {
                jumlahFaktor += faktor
            }
        }
        
        if jumlahFaktor == angka {
            fmt.Print(angka)
        }
    }
}

func main() {
	//2311102251_Marsep Trianto Pakondo
    var a, b int
    
    fmt.Print("Masukkan nilai a: ")
    fmt.Scan(&a)
    fmt.Print("Masukkan nilai b: ")
    fmt.Scan(&b)
    
    fmt.Printf("Perfect numbers antara %v dan %v: ", a, b)
    
    
    perfectNumber(a, b)
}