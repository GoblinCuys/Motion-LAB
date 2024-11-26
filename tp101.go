package main

import "fmt"

func main() {
	var uang, total float64
	fmt.Scan(&uang)

	if uang >= 1000 {
		total = uang * 0.2
		total = uang - total
	}

	fmt.Println(total)
}
