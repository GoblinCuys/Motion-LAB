package main
import "fmt"

func main() {
    var gol, totGol int
    fmt.Scan(&gol)
    
    for gol >= 0 {
        totGol += gol
        fmt.Scan(&gol)
    }
    
    fmt.Println(totGol >= 10)
}