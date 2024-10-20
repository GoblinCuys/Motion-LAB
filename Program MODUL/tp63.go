package main
import "fmt"

func main() {
	var n, i int
	var star string
	fmt.Scan(&n)
	
	star = "*"
	for i = 1; i <= n; i++ {
		fmt.Println(star)
		star += "*"
	}
}