package main 

import "fmt"

func main() {
	var x, n int
	fmt.Print("Bilangan x dan N = ")
	fmt.Scan(&x, &n)
	
	var cek_kelipatan bool
	
	cek_kelipatan = x%n == 0
	
	fmt.Printf("%d kelipatan %d? %t", x, n, cek_kelipatan)
}