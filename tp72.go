package main
import "fmt"

func main() {
    var n, i, hasil int
    fmt.Scan(&n)
    
	i = 100
    for i > 0 {
        hasil = (hasil * 100) + ((n/i%10)*11)
		i = i/10
    }
    
    fmt.Println(hasil)
}