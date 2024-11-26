package main

import "fmt"

func main() {
	var lvl int
	fmt.Scan(&lvl)

	switch {
	case lvl >= 1 && lvl <= 10:
		fmt.Println("Pemula")
	case lvl >= 11 && lvl <= 20:
		fmt.Println("Menengah")
	case lvl >= 21 && lvl <= 30:
		fmt.Println("Ahli")
	case lvl > 30:
		fmt.Println("Master")
	default:
		fmt.Println("Level tidak valid")
	}
}
