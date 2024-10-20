package main
import "fmt"

func main() {
    var n, m, tot, i int
    fmt.Scan(&n,&m)
    
    for i = n; i <= m;  i++{
        tot += i
    }
    fmt.Println(tot)
}