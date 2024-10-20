package main
import "fmt"

func main() {
	var N, i int
	var faktor bool
	fmt.Scan(&N)

	
	for i = 1; i <= N; i++ {
		faktor = N%i == 0
		fmt.Println(i, faktor)
	}
}
