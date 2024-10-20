package main
import "fmt"

func main() {
		var jml, i int
		jml = 0
		i = 1

	for i <= 99 {
		jml += i
		i += 2 
	}

	fmt.Println(jml)
}
