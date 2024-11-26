package main

import "fmt"

func main() {
	var angka int
	fmt.Scan(&angka)

	if (angka >= 1) && (angka <= 9) {
		fmt.Println("Satuan")
	} else if (angka >= 10) && (angka <= 99) {
		fmt.Println("Puluhan")
	} else if (angka >= 100)&& (angka <= 999) {
		fmt.Println("Ratusan")
	} else if (angka >= 10000) && (angka <= 99999) {
		fmt.Println("Puluhan ribu")
	} else if (angka >= 100000) && (angka <= 999999) {
		fmt.Println("Ratusan ribu")
	}
}
