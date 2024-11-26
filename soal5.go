package main

import "fmt"

func main() {
	var h, denda int
	fmt.Scan(&h)

	if h <= 10 {
		denda = h * 2500
	} else if h > 10 {
		denda = h * 5000
	}

	fmt.Println(denda)
}
