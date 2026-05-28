package main

import "fmt"

const N = 1000

func main() {
	var berat [N]float64
	var x, y int
	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	w := (x + y - 1) / y
	var tot [N]float64

	for i := 0; i < x; i++ {
		idx := i / y
		tot[idx] += berat[i]
	}

	for i := 0; i < w; i++ {
		fmt.Printf("%.2f", tot[i])
		if i < w-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	for i := 0; i < w; i++ {
		nIkan := y
		if i == w-1 && x%y != 0 {
			nIkan = x % y
		}
		rata := tot[i] / float64(nIkan)
		fmt.Printf("%.2f", rata)
		if i < w-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}