package main

import "fmt"

func main() {
	var detikDBumi int
	
	fmt.Print("Input dalam satuan detik : ")
	fmt.Scan(&detikDBumi)
	
	jamMars := detikDBumi / 4500
	sisaDetik := detikDBumi % 4500
	menitMars := sisaDetik / 75
	detikMars := sisaDetik % 75
	
	
	fmt.Printf("Maka hasil konversinya adalah %d jam, %d menit dan %d detik di Mars", jamMars, menitMars, detikMars)
	
	
	//1 jam M = 75 menit B 
	//1 jam M = 4500 detik B
	//1 Menit M = 75 detik B	
}