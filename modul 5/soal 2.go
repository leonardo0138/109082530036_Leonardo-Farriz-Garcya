package main

import "fmt"

func bintang(n int) {
	if n == 0 {
		return
	}
	bintang(n - 1)
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Scan(&n)
	bintang(n)
}