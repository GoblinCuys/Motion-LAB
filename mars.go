package main

import (
    "fmt"
)

func main() {
    var detikBumi int

    fmt.Print("Input dalam satuan detik: ")
    fmt.Scan(&detikBumi)

    // Konversi ke waktu Mars
    jamMars := detikBumi / 4500
    sisaDetik := detikBumi % 4500
    menitMars := sisaDetik / 75
    detikMars := sisaDetik % 75

    fmt.Printf("Maka hasil konversinya adalah %d jam, %d menit dan %d detik di Mars\n", jamMars, menitMars, detikMars)
}