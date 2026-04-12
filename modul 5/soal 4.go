package main
import "fmt"	
func baris(n int) {
	if n == 0 {
		return
	}
	 fmt.Printf("%d ", n)
	baris(n - 1)
	if n != 1 {
		fmt.Printf("%d ", n)
	}
}
func main() {
	var n int
	fmt.Scan(&n)
	baris(n)
}
