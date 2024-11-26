package main

import "fmt"

func main() {
	var a, a1 int
	fmt.Print("Masukan angka untuk menghitung faktorial: ")
	fmt.Scan(&a)

	a1 = 1

	if a >= 0 {
		for i := 1; i <= a; i++ {
			a1 = a1 * i
		}
		fmt.Printf("Hasil faktorial dari %d adalah %d", a, a1)
	} else {
		fmt.Println("Faktorial tidak terdefinisi untuk angka negatif!!!")
	}
}
