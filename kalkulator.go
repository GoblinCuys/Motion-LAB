package main

import "fmt"

func main() {
	var a1, a2, tot float64
	var op string
	fmt.Print("SELAMAT DATANG DI KALKULATOR SEDERHANA SALMAN\n")
	fmt.Print("CUMA BISA UNTUK +, -, /, *. SAJA\n")
	fmt.Print("SILAHKAN MASUKAN ANGKA DAN OPERATOR YANG INGIN ANDA HITUNG!!!\n")
	fmt.Print("*pisahkan dengan spasi atau enter\n")
	fmt.Scan(&a1, &op, &a2)

	if op == "+" {
		tot = a1 + a2
	} else if op == "-" {
		tot = a1 - a2
	} else if op == "*" {
		tot = a1 * a2
	} else if op == "/" {
		tot = a1 / a2
	} else {
		fmt.Println("Operasi yang anda masukan salah!!\n")
	}

	fmt.Printf("Hasilnya adalah: %.3f", tot)
	fmt.Println("\nSEKIAN TERIMA KASIH <3")
}
