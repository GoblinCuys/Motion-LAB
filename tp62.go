package main 
import "fmt"

func main() {
	var n, i int
	var v, y float64
	fmt.Scan(&n, &v)
	
	for i=0; i<=n; i++{
		y = 0.0 + v * float64(i) + 0.5 * 9.8 * float64(i) * float64(i)
		fmt.Println(y)
	}
}

