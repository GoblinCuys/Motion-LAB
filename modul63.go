package main

import "fmt"

func main() {
	var tugas, kuis, ujian, praktikum float64

	// Input nilai
	fmt.Print("Masukkan nilai tugas: ")
	fmt.Scan(&tugas)

	fmt.Print("Masukkan nilai kuis: ")
	fmt.Scan(&kuis)

	fmt.Print("Masukkan nilai ujian: ")
	fmt.Scan(&ujian)

	fmt.Print("Masukkan nilai praktikum: ")
	fmt.Scan(&praktikum)

	// Menghitung nilai akhir
	nilaiAkhir := (0.1 * tugas) + (0.2 * kuis) + (0.5 * ujian) + (0.2 * praktikum)

	// Menampilkan hasil
	fmt.Printf("\nNilai Akhir: %.2f\n", nilaiAkhir)
}
