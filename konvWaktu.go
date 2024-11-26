package main

import "fmt"

func main() {
	var a, hasil int
	var w1, w2 string
	fmt.Print("KONVERSI WAKTU!!\n")
	fmt.Print("*konversi dari apa ke apa\n")
	fmt.Print("*format ex: 1 jam menit -> 1 jam berapa menit\n")
	fmt.Scan(&a, &w1, &w2)

	// jam ke menit & detik
	if ((w1 == "jam") || (w1 == "Jam") || (w1 == "JAM")) && ((w2 == "menit") || (w2 == "Menit") || (w2 == "MENIT")) {
		hasil = a * 60
		fmt.Printf("%d %s = %d %s", a, w1, hasil, w2)
	} else if ((w1 == "jam") || (w1 == "Jam") || (w1 == "JAM")) && ((w2 == "detik") || (w2 == "Detik") || (w2 == "DETIK")) {
		hasil = a * 60 * 60
		fmt.Printf("%d %s = %d %s", a, w1, hasil, w2)
		// menit ke jam & detik
	} else if ((w1 == "menit") || (w1 == "Menit") || (w1 == "MENIT")) && ((w2 == "detik") || (w2 == "Detik") || (w2 == "DETIK")) {
		hasil = a * 60
		fmt.Printf("%d %s = %d %s", a, w1, hasil, w2)
	} else if ((w1 == "menit") || (w1 == "Menit") || (w1 == "MENIT")) && ((w2 == "jam") || (w2 == "Jam") || (w2 == "JAM")) {
		mj := float64(a) / 60.0
		fmt.Printf("%d %s = %.2f %s", a, w1, mj, w2)
		// detik ke menit & jam
	} else if ((w1 == "detik") || (w1 == "Detik") || (w1 == "DETIK")) && ((w2 == "jam") || (w2 == "Jam") || (w2 == "JAM")) {
		dj := float64(a) / 3600.0
		fmt.Printf("%d %s = %.2f %s", a, w1, dj, w2)
	} else if ((w1 == "detik") || (w1 == "Detik") || (w1 == "DETIK")) && ((w2 == "menit") || (w2 == "Menit") || (w2 == "MENIT")) {
		dm := float64(a) / 60.0
		fmt.Printf("%d %s = %.2f %s", a, w1, dm, w2)
	} else {
		fmt.Println("TIDAK DAPAT DI TEMUKAN !!!!!!\n")
		fmt.Println("*silahkan masukan konversi dengan benar")
	}

	fmt.Println("\nSEKIAN TERIMA KASIHH <3")
}
