package main

import "fmt"

func main() {
	const skala float64 = 300000
	var jarakPeta, jarakNuBener, jarakKM float64

	fmt.Scan(&jarakPeta)
	jarakNuBener = jarakPeta * skala
	jarakKM = jarakNuBener / 100000

	fmt.Println(jarakKM)
}
