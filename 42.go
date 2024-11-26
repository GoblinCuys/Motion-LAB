package main

import "fmt"

func main() {
	var thn int
	fmt.Scan(&thn)
	
	var cek_kabisat bool
	cek_kabisat = (thn%4 == 0 && thn%5 == 0 && thn%100 != 0)
	fmt.Println(cek_kabisat)
}