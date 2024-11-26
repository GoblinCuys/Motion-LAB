package main
import "fmt"


func main() {
	var m, n, i int
	var cek bool
	fmt.Scan(&m, &n)
	
	cek = true 

	var bil int
	for i = 1; i <= n; i++ {
		fmt.Scan(&bil)

		cek = cek && (bil == m*i)
	}

	fmt.Println(cek)
}
