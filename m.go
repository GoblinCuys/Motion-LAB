package main

import (
    "fmt"
)

func marsToEarth(hours float64, minutes float64) (float64, float64) {
    // Konversi jam Mars ke menit Bumi
    totalMinutesInMars := hours * 60 + minutes
    totalMinutesInEarth := totalMinutesInMars * (75.0 / 60.0)
    
    // Menghitung jam dan menit Bumi
    earthHours := int(totalMinutesInEarth) / 60
    earthMinutes := int(totalMinutesInEarth) % 60

    return float64(earthHours), float64(earthMinutes)
}

func main() {
    // Input waktu dalam jam dan menit di Mars
    var marsHours, marsMinutes float64
    fmt.Print("Masukkan waktu di Mars (jam dan menit): ")
    fmt.Scan(&marsHours, &marsMinutes)

    // Mengonversi waktu Mars ke waktu Bumi
    earthHours, earthMinutes := marsToEarth(marsHours, marsMinutes)

    // Menampilkan hasil konversi
    fmt.Printf("%.0f jam %.0f menit di Mars = %.0f jam %.0f menit di Bumi\n", marsHours, marsMinutes, earthHours, earthMinutes)
}
