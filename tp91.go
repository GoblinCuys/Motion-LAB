package main
import "fmt"

func main() {
	var  suhu, hitung int
	var nama string
	
	fmt.Scan(&suhu, &nama)
	
	if nama == "Celcius" {
		hitung = 32 + ((9*suhu)/5)
		fmt.Printf("Suhu dalam Fahrenheit adalah %d F", hitung)
	}
	if nama == "Fahrenheit" {
		hitung = (suhu-32) * 5/9 
		fmt.Printf("Suhu dalam Celcius adalah %d C", hitung)
	}
}