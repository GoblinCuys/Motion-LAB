package main

import "fmt"

func main() {
	var produk, jumlah, harga int
	fmt.Scan(&produk, &jumlah)

	if produk == 1 {
		harga = jumlah * 2980
	} else if produk == 2 {
		harga = jumlah * 4500
	} else if produk == 3 {
		harga = jumlah * 9980
	} else if produk == 4 {
		harga = jumlah * 4490
	} else if produk == 5 {
		harga = jumlah * 6870
	}

	fmt.Println(harga)
}
