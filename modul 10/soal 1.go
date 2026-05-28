package main
import "fmt"

const N = 1000
func hitung(berat[N] float64,n int, min *float64, max *float64){
	*min = berat[0]
	*max= berat[0]

	for  i:= 1;  i<n; i++ {
		if berat[i]<*min {
			*min= berat[i]
		}
		if berat[i]>*max{
			*max= berat[i]
		}
	}
}
func main (){
	var berat[N] float64
	var max, min float64
	var n int
	fmt.Scan(&n)

	for i:= 0; i<n; i++{
		fmt.Scan(&berat[i])
	}
	hitung(berat, n, &min,&max)
	fmt.Printf("%.2f %.2f\n",min, max )
}