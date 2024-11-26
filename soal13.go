package main

import "fmt"

func main() {
	var a, b, c, d, genap, ganjil int
	fmt.Scan(&a, &b, &c, &d)

	if a%2 == 0 {
		genap = genap + 1
	} else {
		ganjil = ganjil + 1
	}

	if b%2 == 0 {
		genap = genap + 1
	} else {
		ganjil = ganjil + 1
	}

	if c%2 == 0 {
		genap = genap + 1
	} else {
		ganjil = ganjil + 1
	}

	if d%2 == 0 {
		genap = genap + 1
	} else {
		ganjil = ganjil + 1
	}

	fmt.Println(genap, ganjil)
}
