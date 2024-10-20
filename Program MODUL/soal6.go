package main
import "fmt"

func main() {
    var duit, saldo int
    fmt.Scan(&duit)
    
    for duit != 0 {
        saldo = saldo + duit
        fmt.Scan(&duit)
    }
    
    fmt.Println(saldo)
}