package main

import "fmt"

func main() {
	var tb, bb, bmi float64
	fmt.Scan(&tb, &bb)

	bmi = bb / ((tb / 100) * (tb / 100))
	if bmi < 18.5 {
		fmt.Println("Berat badan kurang")
	} else if bmi <= 22.9 {
		fmt.Println("Berat badan normal")
	} else {
		fmt.Println("Kelebihan berat badan")
	}
}
