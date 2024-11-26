package main

import "fmt"

const Phi float64 = 3.14

func main() {
	var r int = 10
	var luas float64
	
	//Phi = 1000
	
	luas = Phi*float64(r*r)
	fmt.Println(luas, Phi)
} 