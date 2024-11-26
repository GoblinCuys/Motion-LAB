package main 

import "fmt"

func main() {
	var a, b, c float64
	
	fmt.Scan(&a, &b, &c)
	
	//mengecek apakah bilangan bulat positif
	if a <= 0 || b <= 0 || c <= 0 {
		fmt.Printf("%.1f, %.1f, dan %.1f segitiga? ", a, b, c)
	}
	
	//mengecek sifat sg3
	if a+b > c && b+c > a && c+a > b {
	fmt.Printf("%.1f, %.1f, dan %.1f segitiga? true ", a, b, c)
	} else {
	fmt.Printf("%.1f, %.1f, dan %.1f segitiga? false", a, b, c)}
}