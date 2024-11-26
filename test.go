package main

import "fmt"

func main() {
    var namaSendiri, namaTeman string

    fmt.Print("Masukkan nama sendiri: ")
    fmt.Scanln(&namaSendiri)
    fmt.Print("Masukkan nama teman: ")
    fmt.Scanln(&namaTeman)

    hasil := len(namaSendiri) > len(namaTeman)

    fmt.Println(hasil)
}
