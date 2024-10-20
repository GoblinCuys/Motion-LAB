package main
import "fmt"
    
func main() {
    var tabung, minggu, totTabung, totSimpan, i int
    fmt.Scan(&tabung)
    fmt.Scan(&minggu)
  
    totTabung = tabung
    totSimpan = tabung
  
    for i = 1; i < minggu; i++ {
        totTabung += 2500
        totSimpan = totTabung + totSimpan
    }
    
    fmt.Println(totSimpan)
} 