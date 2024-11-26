package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	switch {
	case num < 10:
		fmt.Println("Satuan")
	case num < 100:
		fmt.Println("Puluhan")
	case num < 1000:
		fmt.Println("Ratusan")
	case num < 10000:
		fmt.Println("Ribuan")
	case num < 100000:
		fmt.Println("Puluhan Ribu")
	case num < 1000000:
		fmt.Println("Ratusan Ribu")
	}
}
