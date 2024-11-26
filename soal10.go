package main
import "fmt"

func main() {
	var kendaraan string
	var lamaParkir, jamAwalMotor, jamAwalMobil, total int
	
	fmt.Scan(&kendaraan, &lamaParkir)
	jamAwalMotor = 2000
	jamAwalMobil = 3000
	
	if kendaraan == "motor" {
		total = jamAwalMotor + 500 * (lamaParkir - 1)
	}
	if kendaraan == "mobil" {
		total = jamAwalMobil + 1000 * (lamaParkir - 1)
	}
	
	fmt.Println(total)
}