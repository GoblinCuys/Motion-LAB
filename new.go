package main

import "fmt"

func main() {
	var detikDBumi int
	
	fmt.Print("Input dalam bentuk detik: ")
	fmt.Scan(&detikDBumi)
	
	jDM := detikDBumi / 4500
	sD := detikDBumi % 4500
	mDM := sD / 75
	dDM := sD % 75
	
	fmt.Printf("Maka hasil konvesinya adalah, %d jam, %d menit, %d detik di MARS", jDM, mDM,dDM)
}