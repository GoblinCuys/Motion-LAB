package main
import "fmt"

func main () {
	var makanan, minuman, total int
	var tip bool
	fmt.Scan(&makanan, &minuman, &tip)
	
	if tip == true {
		total = makanan + minuman + 5000
	} else if  tip == false {
		total = makanan + minuman
	}
	
	fmt.Println(total)
}