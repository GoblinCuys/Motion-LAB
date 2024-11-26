package main
import "fmt"

func main() {
	var a, b, c int 
	
	fmt.Scan(&a, &b, &c)
	
	if (a + b + c) == 180 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}