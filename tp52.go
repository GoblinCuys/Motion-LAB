package main
import "fmt"

func main() {
	var i,j,tot int
	fmt.Scan(&j,&i)

	for i=i;i<=j;i++{
		fmt.Print(i)
		tot += i
	}
	
	fmt.Printf("\nTotal %d",tot)
}