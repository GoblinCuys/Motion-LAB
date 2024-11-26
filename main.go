package main

import "fmt"

func hitung(x, y float64) float64 {
	return (5/((2*y) + 7) ) + (x*x) - (3*y)
}

func main() {
	var n1, n2 float64
	
	fmt.Print("Masukan Dua angka (pisahkan menggunakan spasi): ")
	fmt.Scan(&n1, &n2)
	
	hasil1 := hitung(n1,n2)
	hasil2 := hitung(n2,n1)
	
	fmt.Printf("%.6f %.6f\n", hasil1, hasil2)
}