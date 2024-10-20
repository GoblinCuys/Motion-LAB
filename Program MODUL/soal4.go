package main
import "fmt"

func main() {
    var n, m, tot, i int
    n = 2
    m = 50
    
    for i = n; i <= m;  i++{
        tot += i
    }
    fmt.Println(tot)
}