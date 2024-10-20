package main

import "fmt"

func main() {
	var modal float64
	var thn int
    fmt.Scan(&modal)
    
   
    for thn = 1; thn <= 10; thn += 1{
        modal = modal * (1 + 0.05)
        fmt.Printf("Tahun ke-%d: Rp%.2f\n", thn, modal)
    }
}