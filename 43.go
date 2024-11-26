package main

import "fmt"

func main() {
	var kode int
	fmt.Scan(&kode)
	
	var digitPertama, digitTerakhir int
	digitPertama = kode / 1000
	digitTerakhir = kode % 10
	
	var apakahDPTcashback bool
	apakahDPTcashback = digitPertama == digitTerakhir

	fmt.Printf("%t", apakahDPTcashback)
}