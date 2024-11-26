package main 
import "fmt"

func main() {
	var num, digit1, digit2, digit3 int
	var tot bool
	
	fmt.Scan(&num)
	
	digit1 = num / 100
	digit2 = (num / 10)%10
	digit3 = num % 10
	
	tot = digit1 < digit2 && digit2 < digit3
	
	fmt.Println(tot)
}