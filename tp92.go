package main
import "fmt"

func main() {
	var ktm, sim bool
	
	fmt.Scan(&ktm, &sim)
	
	if ktm || sim {
		fmt.Println("Diizinkan masuk")
	}else {
		fmt.Println("Tidak diizinkan masuk")
	}
}