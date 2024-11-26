package main

import "fmt"

func main() {
	var huruf string
	fmt.Scan(&huruf)

	if huruf == "A" {
		fmt.Println(5)
	} else if huruf == "B" {
		fmt.Println(4)
	} else if huruf == "C" {
		fmt.Println(3)
	} else if huruf == "D" {
		fmt.Println(2)
	} else if huruf == "E" {
		fmt.Println(1)
	} else if huruf == "T" {
		fmt.Println(0)
	}
}


/