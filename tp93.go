package main
import "fmt"

func main() {
	var gram, x, perKarung, butuh int
	
	fmt.Scan(&gram, &x)
	
	perKarung = 1000/x
	
	butuh = gram / perKarung
	
	if gram % perKarung != 0 {
		butuh = butuh + 1
	}
	fmt.Printf("%d karung", butuh)
	
}