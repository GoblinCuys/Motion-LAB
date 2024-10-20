package main

import "fmt"

func main() {
    var jam int
    var totJam, totHari, rata float64
    
    fmt.Scan(&jam)
    
    for jam >= 0 {
        totJam += float64(jam)
        totHari++
        fmt.Scan(&jam)
    }
    
    rata = totJam / totHari
    fmt.Printf("%.2f\n", rata)
}