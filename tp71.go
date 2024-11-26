package main
import "fmt"

func main() {
	var dek, bang, i int
	var win bool
	fmt.Scan(&dek, &bang)
	
	win = false
	
	for i = 0; i < 1; i++ {
		win = (dek == bang) || (dek - bang == 1) || (dek - bang == -1) || (dek - bang == 5) || (dek - bang == -5)
	}
	fmt.Println(win)
}
