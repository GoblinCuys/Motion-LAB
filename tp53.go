package main
import "fmt"

func main (){
    var i,j int
    var alas,tinggi, tot float64
    fmt.Scan(&j)
    
	for i=1;i<=j;i++{
    fmt.Scan(&alas,&tinggi)
    tot = (alas*tinggi)/2
    fmt.Printf("%.1f\n",tot)
	}
}