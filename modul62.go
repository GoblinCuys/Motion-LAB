package main

import "fmt"

func main() {
    var fahrenheit, koncel float64
	var cek  bool
    
    fmt.Print("Masukkan suhu dalam Fahrenheit: ")
    fmt.Scan(&fahrenheit)
    
    
    koncel = 5.0/9.0 * (fahrenheit - 32)
    
    
    cek = koncel > 28
    
   
    fmt.Printf("Suhu dalam Celcius: %.2f\n", koncel)
    fmt.Printf("Apakah termasuk suhu panas? %t\n", cek)
}