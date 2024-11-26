package main

import (
    "fmt"
)

func main() {
    var a, b, c float64

    fmt.Print("Masukkan tiga panjang sisi: ")
    fmt.Scan(&a, &b, &c)

    // Memeriksa apakah semua sisi positif
    if a <= 0 || b <= 0 || c <= 0 {
        fmt.Printf("%.1f, %.1f, dan %.1f segitiga? false\n", a, b, c)
        return
    }

    // Memeriksa apakah jumlah dua sisi selalu lebih besar dari sisi ketiga
    if a+b > c && b+c > a && c+a > b {
        fmt.Printf("%.1f, %.1f, dan %.1f segitiga? true\n", a, b, c)
    } else {
        fmt.Printf("%.1f, %.1f, dan %.1f segitiga? false\n", a, b, c)
    }
}