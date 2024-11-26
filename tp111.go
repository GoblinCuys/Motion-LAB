package main

import "fmt"

func main() {
	var a, b, hitung float64
	var op string
	fmt.Scan(&a, &op, &b)

	switch {
	case op == "+":
		hitung = a + b
	case op == "-":
		hitung = a - b
	case op == "*":
		hitung = a * b
	case op == "/":
		hitung = a / b
	}

	fmt.Println(hitung)
}
