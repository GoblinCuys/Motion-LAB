package main

import "fmt"
	
func main() {
	var jmlSiswa, A, B int
	fmt.Scan(&jmlSiswa, &A, &B)

	if ((float64(A) + float64(B)) / float64(jmlSiswa)) >=  0.6 {
		if A > B {
		fmt.Println("Kadidat A menang")
		} else if B > A {
		fmt.Println("Kadidat B menang")
		}
	} else {
		fmt.Println("Tidak ada pemenang")
	}	
}
