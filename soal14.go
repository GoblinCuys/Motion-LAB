package main
import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	
	if a % b == 0 {
		fmt.Printf("Ya, %d kelipatan %d", a, b)
	} else {
		fmt.Printf("Tidak, %d bukan kelipatan %d", a, b)
	}
}