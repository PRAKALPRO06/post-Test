package main

import "fmt"


func main()  {
	var a, b int 
	fmt.Print("Masukan nilai a : ")
	fmt.Scan(&a)
	fmt.Print("Masukan nilai b : ")
	fmt.Scan(&b)

	if a%2 == 0{
		a++
	}else if b%2 == 0{
		b--
	}

	totalGanjil_2311102010 := (b-a) / 2+1

	fmt.Printf("Banyaknya ganjil : %d\n", totalGanjil_2311102010)
}

