package main

import "fmt"

func main() {
    var plat, jml int
  
    fmt.Scan(&plat)
    
    jml= 0
    for plat > 0 {
        jml = jml + plat % 10
        plat /= 10
    }
    
    for jml > 9 {
        jml = (jml % 10) + (jml / 10)
    }
    
    fmt.Println(jml)
}