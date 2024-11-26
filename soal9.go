package main
import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, p float64
	
	fmt.Scan(&a, &b, &c)
	p = math.Pow(a, b)
	
	if c == p {
		fmt.Println("benar")
	} else {
		fmt.Println("salah")
	}
}