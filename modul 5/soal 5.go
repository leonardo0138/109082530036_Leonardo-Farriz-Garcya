package main

import "fmt"
func ganjil (n int) {
	if n == 0 {
		return
	}
	ganjil (n - 1)
	if n % 2 != 0 {
		fmt.Printf("%d ", n)
	}
}
func main() {	
	var n int	
	fmt.Scan(&n)
	ganjil (n)
}	

	