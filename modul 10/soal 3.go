package main 
import "fmt"

type balita [100] float64

func hitung(bb balita, n int, min *float64, max *float64){
	*min= bb[0]
	*max= bb[0]

	for  i:= 1;  i<n; i++ {
		if bb[i]<*min {
			*min= bb[i]
		}

		if bb[i]>*max{
			*max= bb[i]
		}
	}
}
func rerata (bb balita, n int)float64{
	total := 0.0
	for i:=0; i<n;i++{
		total += bb[i]
	}
	return total/float64(n)
}
func main(){
	var bb balita
	var n int
	var min, max float64
	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i:= 0; i<n;i++{
		fmt.Print("masukan berat balita ke - ",i+1," : ")
		fmt.Scan(&bb[i])
	}
	hitung(bb,n,&min,&max)
	fmt.Printf("Berat balita minimum: %.2f kg \n",min)
	fmt.Printf("Berat balita maximum: %.2f kg \n",max)
	fmt.Printf("Rerata berat balita: %.2f kg", rerata(bb,n))
}