package main

import "fmt"

func main() {
	var curahHujan string
	fmt.Scan(&curahHujan)

	if curahHujan == "tinggi" {
		fmt.Println("macet")
	} else if curahHujan == "rendah" {
		fmt.Println("tidak macet")
	}
}
